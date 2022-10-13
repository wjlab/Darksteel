package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"tdk/conf"
	"tdk/process"
)

func SearchComputers(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var listComputers []string
	searchComputers := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.Ldap_queries["computers"],
		[]string{"name"},
		nil)
	searchComputer, err := (*l).Search(searchComputers)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range searchComputer.Entries {
		listComputers = append(listComputers, entry.GetAttributeValue("name"))
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] Domain Computers:\n", outputFile)
		for _, j := range listComputers {
			process.OutFile("\t"+j+"\n", outputFile)
		}
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] Computers save file to:  %s\n", outputFile)
	} else {
		fmt.Printf("[*] Domain Computers:\n")
		for _, j := range listComputers {
			fmt.Println("\t" + j)
		}
		fmt.Printf("\n")
	}
}
