package ldap

import (
	"darksteel/conf"
	"darksteel/process"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

func SearchSpn(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var spnStructureList []constrained
	//查询spn
	searchSpn := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.Ldap_queries["spn"],
		[]string{"dn", "cn", "servicePrincipalName"},
		nil)
	spnName, err := (*l).Search(searchSpn)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range spnName.Entries {
		a := constrained{
			entry.DN,
			entry.GetAttributeValues("servicePrincipalName"),
			entry.GetAttributeValue("cn"),
		}
		spnStructureList = append(spnStructureList, a)

	}

	if len(outputFile) != 0 {
		for _, j := range spnStructureList {
			process.OutFile("\n\n[*] SPN Info："+j.CN, outputFile)
			for _, k := range j.SPN {
				process.OutFile("\n\t"+k, outputFile)
			}
		}
		fmt.Printf("[*] Spn save file to: %s\n", outputFile)
	} else {
		for _, j := range spnStructureList {
			fmt.Println("\n[*] SPN：" + j.CN)
			for _, k := range j.SPN {
				fmt.Println("\t" + k)
			}
		}
	}
}
