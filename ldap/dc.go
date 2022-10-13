package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"tdk/conf"
	"tdk/process"
)

func SearchDc(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var listDc []string
	searchDc := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.Ldap_queries["dc"],
		[]string{"name"},
		nil)
	user, err := (*l).Search(searchDc)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range user.Entries {
		listDc = append(listDc, entry.GetAttributeValue("name"))
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] DC Computer:\n", outputFile)
		for _, j := range listDc {
			process.OutFile("\t"+j+"\n", outputFile)
		}
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] DC save file to: %s\n", outputFile)
	} else {
		fmt.Printf("[*] DC Computer:\n")
		for _, j := range listDc {
			fmt.Println("\t" + j)
		}
		fmt.Printf("\n")
	}
}
