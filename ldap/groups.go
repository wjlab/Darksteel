package ldap

import (
	"darksteel/conf"
	"darksteel/process"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

// 查询域管组
func SearchAdmins(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var listAdmins []string
	searchAdmins := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["admins"],
		[]string{"member"},
		nil)
	searchAdmin, err := (*l).Search(searchAdmins)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range searchAdmin.Entries {
		for _, j := range entry.GetAttributeValues("member") {
			listAdmins = append(listAdmins, j)
		}
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] Domain Admins:\n", outputFile)
		for _, j := range listAdmins {
			process.OutFile("\t"+j+"\n", outputFile)
		}
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] Domain Admins save file to:  %s\n", outputFile)
	} else {
		fmt.Printf("[*] Domain Admins:\n")
		for _, j := range listAdmins {
			fmt.Println("\t" + j)
		}
		fmt.Printf("\n")
	}
}

// 查询企业管理员组
func SearchEnterprise(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var listEnterprises []string
	searchEnterprises := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["enterprises"],
		[]string{"member"},
		nil)
	searchEnterprise, err := (*l).Search(searchEnterprises)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range searchEnterprise.Entries {
		for _, j := range entry.GetAttributeValues("member") {
			listEnterprises = append(listEnterprises, j)
		}
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] Enterprise Admins:\n", outputFile)
		for _, j := range listEnterprises {
			process.OutFile("\t"+j+"\n", outputFile)
		}
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] Enterprise Admins save file to:  %s\n", outputFile)
	} else {
		fmt.Printf("[*] Enterprise Admins:\n")
		for _, j := range listEnterprises {
			fmt.Println("\t" + j)
		}
		fmt.Printf("\n")
	}
}
