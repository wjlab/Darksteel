package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"strings"
)

func SearchLdap(l **ldap.Conn, domain string, searchValue string, outputContent string, ldapSizeLimit int) {
	if len(outputContent) != 0 {
		countSplit := strings.Split(outputContent, ",")
		searchLdap := ldap.NewSearchRequest(domain,
			ldap.ScopeWholeSubtree,
			ldap.NeverDerefAliases,
			ldapSizeLimit,
			0,
			false,
			searchValue,
			countSplit,
			nil)
		user, err := (*l).Search(searchLdap)
		if err != nil {
			fmt.Println(err)
		}
		user.Print()
		return
	} else {
		searchLdap := ldap.NewSearchRequest(domain,
			ldap.ScopeWholeSubtree,
			ldap.NeverDerefAliases,
			ldapSizeLimit,
			0,
			false,
			searchValue,
			[]string{},
			nil)
		user, err := (*l).Search(searchLdap)
		if err != nil {
			fmt.Println(err)
		}
		user.Print()
		return
	}
}
