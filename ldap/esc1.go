package ldap

import (
	"darksteel/conf"
	"darksteel/process"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

func SearchEsc1(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var listEsc1 []string
	searchEsc1s := ldap.NewSearchRequest("CN=Configuration,"+domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["esc1"],
		[]string{"cn"},
		nil)
	searchEsc1, err := (*l).Search(searchEsc1s)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range searchEsc1.Entries {
		for _, j := range entry.GetAttributeValues("cn") {
			listEsc1 = append(listEsc1, j)
		}
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] Esc1 vulnerability template:\n", outputFile)
		for _, j := range listEsc1 {
			process.OutFile("\t"+j+"\n", outputFile)
		}
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] Esc1 vulnerability template:  %s\n", outputFile)
	} else {
		fmt.Printf("[*] Esc1 vulnerability template:\n")
		for _, j := range listEsc1 {
			fmt.Println("\t" + j)
		}
		fmt.Printf("\n")
	}
}
