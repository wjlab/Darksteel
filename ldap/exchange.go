package ldap

import (
	"darksteel/conf"
	"darksteel/process"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

// 查询exchange计算机
func SearchExchangeComputer(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var listExchangeComputers []string
	searchExchangeComputers := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["exchangeComputer"],
		[]string{"member"},
		nil)
	searchExchangeComputer, err := (*l).Search(searchExchangeComputers)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range searchExchangeComputer.Entries {
		for _, j := range entry.GetAttributeValues("member") {
			listExchangeComputers = append(listExchangeComputers, j)
		}
	}
	if len(listExchangeComputers) > 1 {
		listExchangeComputers = listExchangeComputers[1:]
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] Exchange Servers:\n", outputFile)
		for _, j := range listExchangeComputers {
			process.OutFile("\t"+j+"\n", outputFile)
		}
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] Exchange Servers save file to:  %s\n", outputFile)
	} else {
		fmt.Printf("[*] Exchange Servers:\n")
		for _, j := range listExchangeComputers {
			fmt.Println("\t" + j)
		}
		fmt.Printf("\n")
	}
}

// 查询Exchange Trusted Subsystem 组成员
func SearchExchangeTrustedSubsystem(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var listExchangeTrustedSubsystems []string
	searchExchangeTrustedSubsystems := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["exchangeTrustedSubsystem"],
		[]string{"member"},
		nil)
	searchExchangeTrustedSubsystem, err := (*l).Search(searchExchangeTrustedSubsystems)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range searchExchangeTrustedSubsystem.Entries {
		for _, j := range entry.GetAttributeValues("member") {
			listExchangeTrustedSubsystems = append(listExchangeTrustedSubsystems, j)
		}
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] Exchange Trusted Subsystem:\n", outputFile)
		for _, j := range listExchangeTrustedSubsystems {
			process.OutFile("\t"+j+"\n", outputFile)
		}
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] Exchange Trusted Subsystem save file to:  %s\n", outputFile)
	} else {
		fmt.Printf("[*] Exchange Trusted Subsystem:\n")
		for _, j := range listExchangeTrustedSubsystems {
			fmt.Println("\t" + j)
		}
		fmt.Printf("\n")
	}
}

// 查询Exchange Organization Management 组成员
func SearchExchangeOrganizationManagement(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var listExchangeOrganizationManagements []string
	searchExchangeOrganizationManagements := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["exchangeOrganizationManagement"],
		[]string{"member"},
		nil)
	searchExchangeOrganizationManagement, err := (*l).Search(searchExchangeOrganizationManagements)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range searchExchangeOrganizationManagement.Entries {
		for _, j := range entry.GetAttributeValues("member") {
			listExchangeOrganizationManagements = append(listExchangeOrganizationManagements, j)
		}
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] Exchange Organization Management:\n", outputFile)
		for _, j := range listExchangeOrganizationManagements {
			process.OutFile("\t"+j+"\n", outputFile)
		}
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] Exchange Organization Management save file to:  %s\n", outputFile)
	} else {
		fmt.Printf("[*] Exchange Organization Management:\n")
		for _, j := range listExchangeOrganizationManagements {
			fmt.Println("\t" + j)
		}
		fmt.Printf("\n")
	}
}
