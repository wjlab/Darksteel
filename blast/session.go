package blast

import (
	"fmt"
	"github.com/ropnop/kerbrute/util"
	"html/template"
	"strings"

	"github.com/ropnop/gokrb5/v8/iana/errorcode"

	kclient "github.com/ropnop/gokrb5/v8/client"
	kconfig "github.com/ropnop/gokrb5/v8/config"
	"github.com/ropnop/gokrb5/v8/messages"
)

const krb5ConfigTemplateDNS = `[libdefaults]
dns_lookup_kdc = true
default_realm = {{.Realm}}
`

const krb5ConfigTemplateKDC = `[libdefaults]
default_realm = {{.Realm}}
[realms]
{{.Realm}} = {
	kdc = {{.DomainController}}
	admin_server = {{.DomainController}}
}
`

type KerbruteSession struct {
	Domain string
	Realm  string
	kDcs   map[int]string
	Config *kconfig.Config
	Logger *util.Logger
}

func NewKerbruteSession(domain string, domainController string) (k KerbruteSession, err error) {
	if domain == "" {
		return k, fmt.Errorf("domain must not be empty")
	}

	//realm大写,配置config格式
	realm := strings.ToUpper(domain)
	configString := buildKrb5Template(realm, domainController)
	Config, err := kconfig.NewFromString(configString)
	if err != nil {
		panic(err)
	}

	_, kDcs, err := Config.GetKDCs(realm, false)
	if err != nil {
		err = fmt.Errorf("[!] Couldn't find any KDCs for realm %s. Please specify a Domain Controller", realm)
	}
	k = KerbruteSession{
		Domain: domain,
		Realm:  realm,
		kDcs:   kDcs,
		Config: Config,
	}
	return k, err

}

// 构建krb5配置文件格式
func buildKrb5Template(realm, domainController string) string {
	data := map[string]interface{}{
		"Realm":            realm,
		"DomainController": domainController,
	}
	var kTemplate string
	if domainController == "" {
		kTemplate = krb5ConfigTemplateDNS
	} else {
		kTemplate = krb5ConfigTemplateKDC
	}
	t := template.Must(template.New("krb5ConfigString").Parse(kTemplate))
	builder := &strings.Builder{}
	if err := t.Execute(builder, data); err != nil {
		panic(err)
	}
	return builder.String()
}

func (k KerbruteSession) TestUsername(username string) (bool, error) {
	// 不判断预身份认证
	cl := kclient.NewWithPassword(username, k.Realm, "", k.Config, kclient.DisablePAFXFAST(true))

	req, err := messages.NewASReqForTGT(cl.Credentials.Domain(), cl.Config, cl.Credentials.CName())
	if err != nil {
		fmt.Printf(err.Error())
	}
	b, err := req.Marshal()
	if err != nil {
		return false, err
	}
	rb, err := cl.SendToKDC(b, k.Realm)
	if err == nil {
		// 如果没有错误则判断设置了不需要预身份认证
		var ASRep messages.ASRep
		err = ASRep.Unmarshal(rb)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	e, ok := err.(messages.KRBError)
	if !ok {
		return false, err
	}
	switch e.ErrorCode {
	case errorcode.KDC_ERR_PREAUTH_REQUIRED:
		return true, nil
	default:
		return false, err
	}
}

func (k KerbruteSession) TestLogin(username, password string) (bool, error) {
	Client := kclient.NewWithPassword(username, k.Realm, password, k.Config, kclient.DisablePAFXFAST(true), kclient.AssumePreAuthentication(true))
	defer Client.Destroy()
	if ok, err := Client.IsConfigured(); !ok {
		return false, err
	}
	err := Client.Login()
	if err == nil {
		return true, err
	}
	success, err := k.TestLoginError(err)
	return success, err
}
