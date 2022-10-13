package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"tdk/conf"
	"tdk/process"
)

func SearchMsSqlServer(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var listMsSqlServer []string
	searchMsSqlServer := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.Ldap_queries["ms_sql"],
		[]string{"name"},
		nil)
	user, err := (*l).Search(searchMsSqlServer)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range user.Entries {
		listMsSqlServer = append(listMsSqlServer, entry.GetAttributeValue("name"))
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] MsSql Computer:\n", outputFile)
		for _, j := range listMsSqlServer {
			process.OutFile("\t"+j+"\n", outputFile)
		}
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] Mssql save file to: %s\n", outputFile)
	} else {
		fmt.Printf("[*] MsSql Computer:\n")
		for _, j := range listMsSqlServer {
			fmt.Println("\t" + j)
		}
		fmt.Printf("\n")
	}
}
