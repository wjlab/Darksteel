package conf

import (
	"encoding/hex"
	"fmt"
	"github.com/vincd/savoir/modules/paquet/krb5/crypto"
	"github.com/vincd/savoir/utils"
	"golang.org/x/net/proxy"
	"strings"
)

func Banner() {
	fmt.Println(" ____    ______  ____    __  __   ____    ______  ____    ____    __       ")
	fmt.Println("/\\  _`\\ /\\  _  \\/\\  _`\\ /\\ \\/\\ \\ /\\  _`\\ /\\__  _\\/\\  _`\\ /\\  _`\\ /\\ \\      ")
	fmt.Println("\\ \\ \\/\\ \\ \\ \\L\\ \\ \\ \\L\\ \\ \\ \\/'/'\\ \\,\\L\\_\\/_/\\ \\/\\ \\ \\L\\_\\ \\ \\L\\_\\ \\ \\    ")
	fmt.Println(" \\ \\ \\ \\ \\ \\  __ \\ \\ ,  /\\ \\ , <  \\/_\\__ \\  \\ \\ \\ \\ \\  _\\L\\ \\  _\\L\\ \\ \\  _")
	fmt.Println("  \\ \\ \\_\\ \\ \\ \\/\\ \\ \\ \\\\ \\\\ \\ \\\\`\\  /\\ \\L\\ \\ \\ \\ \\ \\ \\ \\L\\ \\ \\ \\L\\ \\ \\ \\L\\ \\ ")
	fmt.Println("   \\ \\____/\\ \\_\\ \\_\\ \\_\\ \\_\\ \\_\\ \\_\\\\ `\\____\\ \\ \\_\\ \\ \\____/\\ \\____/\\ \\____/  ")
	fmt.Println("    \\/___/  \\/_/\\/_/\\/_/\\/ /\\/_/\\/_/ \\/_____/  \\/_/  \\/___/  \\/___/  \\/___/   \n")
	fmt.Println("   v2.0.0\n")
}

var LdapQueries = map[string]string{
	"users":                          "(objectClass=user)",
	"computers":                      "(objectClass=Computer)",
	"dc":                             "(&(objectCategory=Computer)(userAccountControl:1.2.840.113556.1.4.803:=8192))",
	"spn":                            "(&(servicePrincipalName=*))",
	"unconstrained_users":            "(&(samAccountType=805306368)(userAccountControl:1.2.840.113556.1.4.803:=524288))",
	"unconstrained_computers":        "(&(samAccountType=805306369)(userAccountControl:1.2.840.113556.1.4.803:=524288))",
	"constrained_computers":          "(&(samAccountType=805306369)(msds-allowedtodelegateto=*))",
	"constrained_users":              "(&(samAccountType=805306368)(msds-allowedtodelegateto=*))",
	"ms_sql":                         "(&(objectCategory=computer)(servicePrincipalName=MSSQLSvc*))",
	"ou":                             "(&(objectCategory=organizationalUnit)(ou=*))",
	"asreproast":                     "(&(UserAccountControl:1.2.840.113556.1.4.803:=4194304)(!(UserAccountControl:1.2.840.113556.1.4.803:=2))(!(objectCategory=computer)))",
	"kerberoasting":                  "(&(!(UserAccountControl:1.2.840.113556.1.4.803:=2))(samAccountType=805306368)(servicePrincipalName=*)(!samAccountName=krbtgt))",
	"fuzzy":                          "(description=*)",
	"survivalComputer":               "(&(objectcategory=computer)(!(useraccountcontrol:1.2.840.113556.1.4.803:=2))(pwdlastset>=131932198595370000)(|(!lastlogontimestamp=*)(&(lastlogontimestamp=*)(lastlogontimestamp>=131932198595370000))))",
	"admins":                         "(&(sAMAccountName=Domain Admins))",
	"enterprises":                    "(&(sAMAccountName=Enterprise Admins))",
	"exchangeComputer":               "(&(objectClass=group)(cn=Exchange Servers))",
	"exchangeTrustedSubsystem":       "(&(objectClass=group)(cn=Exchange Trusted Subsystem))",
	"exchangeOrganizationManagement": "(&(objectClass=group)(cn=Organization Management))",
	"trustDomain":                    "(&(objectClass=trustedDomain))",
	"adminSDHolder":                  "(&(objectcategory=person)(samaccountname=*)(admincount=1))",
	"sIDHistory":                     "(&(sIDHistory=*))",
	"caComputer":                     "(&(objectCategory=pKIEnrollmentService))",
	"esc1":                           "(&(objectclass=pkicertificatetemplate)(!(mspki-enrollment-flag:1.2.840.113556.1.4.804:=2))(|(mspki-ra-signature=0)(!(mspki-ra-signature=*)))(|(pkiextendedkeyusage=1.3.6.1.4.1.311.20.2.2)(pkiextendedkeyusage=1.3.6.1.5.5.7.3.2)(pkiextendedkeyusage=1.3.6.1.5.2.3.4)(pkiextendedkeyusage=2.5.29.37.0)(!(pkiextendedkeyusage=*)))(mspki-certificate-name-flag:1.2.840.113556.1.4.804:=1)(!(cn=OfflineRouter))(!(cn=CA))(!(cn=SubCA)))",
	"esc2":                           "(&(objectclass=pkicertificatetemplate)(!(mspki-enrollment-flag:1.2.840.113556.1.4.804:=2))(|(mspki-ra-signature=0)(!(mspki-ra-signature=*)))(|(pkiextendedkeyusage=2.5.29.37.0)(!(pkiextendedkeyusage=*)))(!(cn=CA))(!(cn=SubCA)))",
	"aclUser":                        "(&(objectClass=top)(|(objectClass=user)(objectClass=group)(objectClass=domainDNS)))",
}

var supportedETypeMapping = map[string]int32{
	// "3":      crypto.DES_CBC_MD5,
	// "des":    crypto.DES_CBC_MD5,
	"17":     crypto.AES128_CTS_HMAC_SHA1_96,
	"aes128": crypto.AES128_CTS_HMAC_SHA1_96,
	"18":     crypto.AES256_CTS_HMAC_SHA1_96,
	"aes256": crypto.AES256_CTS_HMAC_SHA1_96,
	"23":     crypto.RC4_HMAC,
	"rc4":    crypto.RC4_HMAC,
	"ntlm":   crypto.RC4_HMAC,
}

// aes256作为默认值
func GetETypeFromFlagValue(enctype string) int32 {
	if val, ok := supportedETypeMapping[strings.ToLower(enctype)]; ok {
		return val
	}

	return supportedETypeMapping["aes256"]
}

// 将密钥字符串解析为字节片
func GetKeyFlagValue(key string) ([]byte, error) {
	b, err := hex.DecodeString(key)
	if err != nil {
		return nil, fmt.Errorf("flag --key value cannot be unhexlify")
	}

	return b, nil
}

// 验证域用户凭证
func ValidateDomainUserFlags(username string, password string, key string) error {
	// 即使需要用户名，也要检查用户名是否不为空
	if len(username) == 0 {
		return fmt.Errorf("-username cannot be empty")
	}

	if len(password) > 0 && len(key) > 0 {
		return fmt.Errorf("flags --password and --key cannot be set on the same command")
	}

	if len(key) > 0 {
		if _, err := hex.DecodeString(key); err != nil {
			return fmt.Errorf("flag --key is not a valid hex string")
		}
	}

	return nil
}

// 验证域用户凭证和票据
func ValidateDomainUserFlagsWithTicket(username string, password string, key string, ticket string) error {
	if err := ValidateDomainUserFlags(username, password, key); err != nil {
		return err
	}

	if len(ticket) > 0 && (len(password) > 0 || len(key) > 0) {
		return fmt.Errorf("-ticket and -password or -key cannot be set on the same command")
	}
	return nil
}

// Validate the flag has the EType values from the supported ones. Call this function in `Args`.
func ValidateETypeFlag(enctype string) error {
	if _, ok := supportedETypeMapping[strings.ToLower(enctype)]; !ok {
		return fmt.Errorf("flag --enctype value is not valid")
	}

	return nil
}

// 验证输出格式是否正确
func ValidateFormatFlag(format string) error {
	if format != "" && format != "john" && format != "hashcat" {
		return fmt.Errorf("-format value should be `john` or `hashcat`")
	}

	return nil
}

// 输出域相关信息
func PrintDomainInformation(domain string, dcIp string) {
	fmt.Printf("[*] Target domain: %s (%s)\n", domain, dcIp)
}

// socks代理使用
func GetKdcDialer(socksAddress string) (proxy.Dialer, error) {
	return utils.GetDialerWithSocks(socksAddress)
}
