package ldap

import (
	"darksteel/conf"
	"darksteel/process"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

func SearchUsers(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var listUser []string
	searchUsers := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.Ldap_queries["users"],
		[]string{"sAMAccountName"},
		nil)
	user, err := (*l).Search(searchUsers)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range user.Entries {
		listUser = append(listUser, entry.GetAttributeValue("sAMAccountName"))
	}
	if len(outputFile) != 0 {
		process.OutFile(("[*] Domain User:\n"), outputFile)
		for _, j := range listUser {
			process.OutFile("\t"+j+"\n", outputFile)
		}
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] Users save file to: %s\n", outputFile)
	} else {
		fmt.Printf("[*] Domain User:\n")
		for _, j := range listUser {
			fmt.Println("\t" + j)
		}
		fmt.Printf("\n")
	}
}
