package ldap

import (
	"darksteel/conf"
	"darksteel/process"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

// 查询信任域
func SearchTrustDomain(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var listTrustDomains []string
	searchTrustDomains := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["trustDomain"],
		[]string{"cn"},
		nil)
	searchTrustDomain, err := (*l).Search(searchTrustDomains)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range searchTrustDomain.Entries {
		for _, j := range entry.GetAttributeValues("cn") {
			listTrustDomains = append(listTrustDomains, j)
		}
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] Trust Domain:\n", outputFile)
		for _, j := range listTrustDomains {
			process.OutFile("\t"+j+"\n", outputFile)
		}
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] Trust Domain save file to:  %s\n", outputFile)
	} else {
		fmt.Printf("[*] Trust Domain:\n")
		for _, j := range listTrustDomains {
			fmt.Println("\t" + j)
		}
		fmt.Printf("\n")
	}
}
