package ldap

import (
	"darksteel/conf"
	"darksteel/process"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

func SearchMsSqlServer(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	var listMsSqlServer []string
	searchMsSqlServer := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["ms_sql"],
		[]string{"name"},
		nil)
	//user, err := (*l).Search(searchMsSqlServer)
	searchMsSqlServers, err := (*l).SearchWithPaging(searchMsSqlServer, 10000)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range searchMsSqlServers.Entries {
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
