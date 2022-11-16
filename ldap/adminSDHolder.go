package ldap

import (
	"darksteel/conf"
	"darksteel/process"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

func SearchAdminSDHolder(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var listAdminSDHolders []string
	searchAdminSDHolders := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["adminSDHolder"],
		[]string{"cn"},
		nil)
	searchAdminSDHolder, err := (*l).Search(searchAdminSDHolders)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range searchAdminSDHolder.Entries {
		for _, j := range entry.GetAttributeValues("cn") {
			listAdminSDHolders = append(listAdminSDHolders, j)
		}
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] AdminSDHolder:\n", outputFile)
		for _, j := range listAdminSDHolders {
			process.OutFile("\t"+j+"\n", outputFile)
		}
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] AdminSDHolder save file to:  %s\n", outputFile)
	} else {
		fmt.Printf("[*] AdminSDHolder:\n")
		for _, j := range listAdminSDHolders {
			fmt.Println("\t" + j)
		}
		fmt.Printf("\n")
	}
}
