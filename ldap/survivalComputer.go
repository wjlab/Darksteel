package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"tdk/conf"
	"tdk/process"
)

func SearchSurvivalComputer(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	type survivalComputer struct {
		CN          string
		DESCRIPTION string
	}
	var listSurvivalComputer []survivalComputer
	searchSurvivalComputer := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.Ldap_queries["survivalComputer"],
		[]string{"cn", "operatingSystem"},
		nil)
	searchSurvival, err := (*l).Search(searchSurvivalComputer)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range searchSurvival.Entries {
		a := survivalComputer{
			entry.GetAttributeValue("cn"),
			entry.GetAttributeValue("operatingSystem"),
		}
		listSurvivalComputer = append(listSurvivalComputer, a)
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] Survival Computer:\n", outputFile)
		for _, j := range listSurvivalComputer {
			process.OutFile(fmt.Sprintf("\t%s --> %s\n", j.CN, j.DESCRIPTION), outputFile)
		}
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] Survival computer save file to: %s\n", outputFile)
	} else {
		fmt.Println("[*] Survival Computer:")
		for _, j := range listSurvivalComputer {
			fmt.Println(fmt.Sprintf("\t%s --> %s", j.CN, j.DESCRIPTION))
		}
		fmt.Printf("\n")
	}
}
