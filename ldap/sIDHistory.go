package ldap

import (
	"darksteel/conf"
	"darksteel/process"
	"encoding/base64"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

func SearchSIDHistory(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	type sIDHistory struct {
		SIDHISTORYNAME string
		SIDNAME        string
		SIDHISTORY     string
	}
	var listSIDHistorys []sIDHistory
	searchSIDHistorys := ldap.NewSearchRequest(domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["sIDHistory"],
		[]string{"cn", "sIDHistory"},
		nil)
	searchSIDHistory, err := (*l).Search(searchSIDHistorys)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range searchSIDHistory.Entries {
		// 查询sIDHistory
		for _, i := range entry.GetAttributeValues("sIDHistory") {
			base64CreatorSID := base64.StdEncoding.EncodeToString([]byte(i))
			transferSID, _ := base64.StdEncoding.DecodeString(base64CreatorSID)
			creatorSID := process.Decode(transferSID)

			//通过sid查询cn
			filter1 := fmt.Sprintf("(&(objectSid=" + creatorSID.String() + "))")
			msDSConstraint := ldap.NewSearchRequest(
				domain,
				ldap.ScopeWholeSubtree,
				ldap.NeverDerefAliases,
				ldapSizeLimit,
				0,
				false,
				filter1,
				[]string{"cn"}, // Attributes []string
				nil,
			)
			sr1, _ := (*l).Search(msDSConstraint)
			for _, j1 := range sr1.Entries {
				for _, i1 := range j1.GetAttributeValues("cn") {
					a := sIDHistory{
						entry.DN,
						i1,
						creatorSID.String(),
					}
					listSIDHistorys = append(listSIDHistorys, a)
				}
			}

		}
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] sIDHistory:\n", outputFile)
		for _, j := range listSIDHistorys {
			process.OutFile("\t"+j.SIDHISTORYNAME+" -> "+j.SIDHISTORY+" ["+j.SIDNAME+"]\n", outputFile)
		}
		process.OutFile("\n", outputFile)
		fmt.Printf("[*] sIDHistory save file to:  %s\n", outputFile)
	} else {
		fmt.Printf("[*] sIDHistory:\n")
		for _, j := range listSIDHistorys {
			fmt.Println("\t"+j.SIDHISTORYNAME+" ->  "+j.SIDHISTORY+" ["+j.SIDNAME+"]\n", outputFile)
		}
	}
}
