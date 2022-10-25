package ldap

import (
	"darksteel/conf"
	"darksteel/process"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"strings"
)

func Fuzzy(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string, fuzzName string) {
	type fuzz struct {
		CN          string
		DESCRIPTION string
	}
	var listFuzz []fuzz
	searchFuzz := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["fuzzy"],
		[]string{"description"},
		nil)
	user, err := (*l).Search(searchFuzz)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range user.Entries {
		a := fuzz{
			entry.DN,
			entry.GetAttributeValue("description"),
		}
		listFuzz = append(listFuzz, a)
	}
	if len(outputFile) != 0 {
		for _, j := range listFuzz {
			if strings.Contains(j.DESCRIPTION, fuzzName) {
				process.OutFile(fmt.Sprintf("[*] %s   --> %s\n", j.CN, j.DESCRIPTION), outputFile)
			}

		}
		fmt.Printf("[*] fuzz save file to: %s\n", outputFile)
	} else {
		for _, j := range listFuzz {
			if strings.Contains(j.DESCRIPTION, fuzzName) {
				fmt.Printf("[*] %s   --> %s\n", j.CN, j.DESCRIPTION)
			}
		}
	}
}
