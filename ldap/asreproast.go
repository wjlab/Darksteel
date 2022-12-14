package ldap

import (
	"darksteel/conf"
	"darksteel/process"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

func SearchRoast(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var listRoast []string
	searchRoast := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["asreproast"],
		[]string{"sAMAccountName"},
		nil)
	user, err := (*l).Search(searchRoast)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range user.Entries {
		listRoast = append(listRoast, entry.GetAttributeValue("sAMAccountName"))
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] Asreproast User:\n", outputFile)
		for _, j := range listRoast {
			process.OutFile("\t"+j+"\n", outputFile)
		}
		fmt.Printf("[*] Asreproast save file to: %s\n", outputFile)
	} else {
		fmt.Printf("[*] Asreproast User:\n")
		for _, j := range listRoast {
			fmt.Println("\t" + j)
		}
		fmt.Printf("\n")
	}
}
