package ldap

import (
	"darksteel/conf"
	"darksteel/process"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

func SearchComputers(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var listComputers []string
	var listNumber []int
	searchComputers := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["computers"],
		[]string{"name"},
		nil)
	//searchComputer, err := (*l).Search(searchComputers)
	searchComputer, err := (*l).SearchWithPaging(searchComputers, 10000)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range searchComputer.Entries {
		listComputers = append(listComputers, entry.GetAttributeValue("name"))
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] Domain Computers:\n", outputFile)
		for i, j := range listComputers {
			listNumber = append(listNumber, i)
			process.OutFile("\t"+j+"\n", outputFile)
		}
		process.OutFile(fmt.Sprintf("Number of computers: %d\n", len(listNumber)), outputFile)
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] Computers save file to:  %s\n", outputFile)
	} else {
		fmt.Printf("[*] Domain Computers:\n")
		for i, j := range listComputers {
			fmt.Println("\t" + j)
			listNumber = append(listNumber, i)
		}
		fmt.Printf("Number of computers: %d\n", len(listNumber))
		fmt.Printf("\n")
	}
}
