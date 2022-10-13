package ldap

import (
	"darksteel/process"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

func SearchMaq(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var listMaq []string
	//查询MAQ
	searchMaq := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		fmt.Sprintf("(&(distinguishedName=%s))", domain),
		[]string{"ms-DS-MachineAccountQuota"},
		nil)
	msq, err := (*l).Search(searchMaq)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range msq.Entries {
		listMaq = append(listMaq, entry.GetAttributeValue("ms-DS-MachineAccountQuota"))
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] Maq Number:\n", outputFile)
		for _, j := range listMaq {
			process.OutFile("\t"+j+"\n", outputFile)
		}
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] Maq save file to: %s\n", outputFile)
	} else {
		fmt.Printf("[*] Maq Number:\n")
		for _, j := range listMaq {
			fmt.Println("\t" + j)
		}
		fmt.Printf("\n")
	}
}
