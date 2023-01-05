package ldap

import (
	"darksteel/conf"
	"darksteel/process"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

func SearchCaComputer(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var listCaComputer []string
	searchCaComputers := ldap.NewSearchRequest("CN=Configuration,"+domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["caComputer"],
		[]string{"cn"},
		nil)
	searchCaComputer, err := (*l).Search(searchCaComputers)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range searchCaComputer.Entries {
		for _, j := range entry.GetAttributeValues("cn") {
			listCaComputer = append(listCaComputer, j)
		}
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] Ca Computer:\n", outputFile)
		for _, j := range listCaComputer {
			process.OutFile("\t"+j+"\n", outputFile)
		}
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] Ca Computer save file to:  %s\n", outputFile)
	} else {
		fmt.Printf("[*] Ca Computer:\n")
		for _, j := range listCaComputer {
			fmt.Println("\t" + j)
		}
		fmt.Printf("\n")
	}
}
