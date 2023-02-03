package ldap

import (
	"darksteel/conf"
	"darksteel/process"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

func SearchOU(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var listOU []string
	searchOU := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["ou"],
		[]string{"name"},
		nil)
	//user, err := (*l).Search(searchOU)
	searchOUs, err := (*l).SearchWithPaging(searchOU, 10000)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range searchOUs.Entries {
		listOU = append(listOU, entry.GetAttributeValue("name"))
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] OU :\n", outputFile)
		for _, j := range listOU {
			process.OutFile("\t"+j+"\n", outputFile)
		}
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] OU save file to: %s\n", outputFile)
	} else {
		fmt.Printf("[*] OU :\n")
		for _, j := range listOU {
			fmt.Println("\t" + j)
		}
		fmt.Printf("\n")
	}
}
