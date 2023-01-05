package ldap

import (
	"darksteel/conf"
	"darksteel/process"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

func SearchEsc2(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var listEsc2 []string
	searchEsc2s := ldap.NewSearchRequest("CN=Configuration,"+domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["esc2"],
		[]string{"cn"},
		nil)
	searchEsc2, err := (*l).Search(searchEsc2s)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range searchEsc2.Entries {
		for _, j := range entry.GetAttributeValues("cn") {
			listEsc2 = append(listEsc2, j)
		}
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] Esc2 vulnerability template:\n", outputFile)
		for _, j := range listEsc2 {
			process.OutFile("\t"+j+"\n", outputFile)
		}
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] Esc2 vulnerability template:  %s\n", outputFile)
	} else {
		fmt.Printf("[*] Esc2 vulnerability template:\n")
		for _, j := range listEsc2 {
			fmt.Println("\t" + j)
		}
		fmt.Printf("\n")
	}
}
