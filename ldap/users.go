package ldap

import (
	"darksteel/conf"
	"darksteel/process"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

func SearchUsers(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var listUser []string
	var listNumber []int
	searchUsers := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["users"],
		[]string{"sAMAccountName"},
		nil)
	//user, err := (*l).Search(searchUsers)
	user, err := (*l).SearchWithPaging(searchUsers, 10000)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range user.Entries {
		listUser = append(listUser, entry.GetAttributeValue("sAMAccountName"))
	}
	if len(outputFile) != 0 {
		process.OutFile(("[*] Domain User:\n"), outputFile)
		for i, j := range listUser {
			listNumber = append(listNumber, i)
			process.OutFile("\t"+j+"\n", outputFile)
		}
		process.OutFile(fmt.Sprintf("Number of users: %d\n", len(listNumber)), outputFile)
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] Users save file to: %s\n", outputFile)
	} else {
		fmt.Printf("[*] Domain User:\n")
		for i, j := range listUser {
			fmt.Println("\t" + j)
			listNumber = append(listNumber, i)
		}
		fmt.Printf("Number of users: %d\n", len(listNumber))
		fmt.Printf("\n")
	}
}
