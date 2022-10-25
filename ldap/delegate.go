package ldap

import (
	"darksteel/conf"
	"darksteel/process"
	"encoding/base64"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

var (
	listComputerUnconstrained    []structResourceBasedConstraints
	listUserUnconstrained        []structResourceBasedConstraints
	listComputerConstrained      []constrained
	listUserConstrained          []constrained
	listResourceBasedConstraints []structResourceBasedConstraints
)

type constrained struct {
	CN   string
	SPN  []string
	NAME string
}

type structResourceBasedConstraints struct {
	CN   string
	SID  string
	NAME string
}

func SearchUnconstrained(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	//查询非约束委派机器
	computerUnconstrained := ldap.NewSearchRequest(
		domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["unconstrained_computers"],
		[]string{"dn", "cn"},
		nil)
	computerUnconstraint, err := (*l).Search(computerUnconstrained)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range computerUnconstraint.Entries {
		a := structResourceBasedConstraints{
			entry.DN,
			"",
			entry.GetAttributeValue("cn"),
		}
		listComputerUnconstrained = append(listComputerUnconstrained, a)
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] 非约束委派机器：\n", outputFile)
		for _, j := range listComputerUnconstrained {
			process.OutFile("\t"+j.CN+" ["+j.NAME+"]\n", outputFile)
		}
	} else {
		fmt.Println("[*] 非约束委派机器：")
		for _, j := range listComputerUnconstrained {
			fmt.Println("\t" + j.CN + " [" + j.NAME + "]")
		}
	}

	//查询非约束委派用户
	userUnconstrained := ldap.NewSearchRequest(
		domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["unconstrained_users"],
		[]string{"dn", "cn"},
		nil)
	userUnconstraint, err := (*l).Search(userUnconstrained)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range userUnconstraint.Entries {
		a := structResourceBasedConstraints{
			entry.DN,
			"",
			entry.GetAttributeValue("cn"),
		}
		listUserUnconstrained = append(listUserUnconstrained, a)
	}
	if len(outputFile) != 0 {
		process.OutFile("[*] 非约束委派用户：\n", outputFile)
		for _, j := range listUserUnconstrained {
			process.OutFile("\t"+j.CN+" ["+j.NAME+"]", outputFile)
		}
		fmt.Printf("[*] Unconstrained save file to: %s\n", outputFile)
	} else {
		fmt.Println("[*] 非约束委派用户：")
		for _, j := range listUserUnconstrained {
			fmt.Println("\t" + j.CN + " [" + j.NAME + "]")
		}
	}

}

func SearchConstrained(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	//查询约束委派机器
	computerConstrained := ldap.NewSearchRequest(
		domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["constrained_computers"],
		[]string{"dn", "cn", "msDS-AllowedToDelegateTo"},
		nil)
	computerConstraint, err := (*l).Search(computerConstrained)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range computerConstraint.Entries {
		a := constrained{
			entry.DN,
			entry.GetAttributeValues("msDS-AllowedToDelegateTo"),
			entry.GetAttributeValue("cn"),
		}
		listComputerConstrained = append(listComputerConstrained, a)
	}

	if len(outputFile) != 0 {
		process.OutFile("\n[*] 约束委派机器：\n", outputFile)
		for _, j := range listComputerConstrained {
			process.OutFile("\t"+j.CN+" ["+j.NAME+"]", outputFile)
			for _, k := range j.SPN {
				process.OutFile("\n\t"+k, outputFile)
			}
		}
	} else {
		fmt.Println("[*] 约束委派机器：")
		for _, j := range listComputerConstrained {
			fmt.Println("\t" + j.CN + " [" + j.NAME + "]")
			for _, k := range j.SPN {
				fmt.Println("\t" + k)
			}
		}
	}

	//查询约束委派用户
	userConstrained := ldap.NewSearchRequest(
		domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		conf.LdapQueries["constrained_users"],
		[]string{"dn", "cn", "msDS-AllowedToDelegateTo"},
		nil)
	userConstraint, err := (*l).Search(userConstrained)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range userConstraint.Entries {
		a := constrained{
			entry.DN,
			entry.GetAttributeValues("msDS-AllowedToDelegateTo"),
			entry.GetAttributeValue("cn"),
		}
		listUserConstrained = append(listUserConstrained, a)
	}

	if len(outputFile) != 0 {
		process.OutFile("\n[*] 约束委派用户：\n", outputFile)
		for _, j := range listUserConstrained {
			process.OutFile("\t"+j.CN+" ["+j.NAME+"]", outputFile)
			for _, k := range j.SPN {
				process.OutFile("\n\t"+k+"", outputFile)
			}
			fmt.Printf("[*] Constrained save file to: %s\n", outputFile)
		}
	} else {
		fmt.Println("[*] 约束委派用户：")
		for _, j := range listUserConstrained {
			fmt.Println("\t" + j.CN + " [" + j.NAME + "]")
			for _, k := range j.SPN {
				fmt.Println("\t" + k)
			}
		}
	}
}

func SearchBasedConstraints(l **ldap.Conn, domain string, ldapSizeLimit int, outputFile string) {
	//基于资源约束委派 查询机器是由哪些用户加入到域的

	//查询computer
	resourceBasedConstraints := ldap.NewSearchRequest(
		domain,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		ldapSizeLimit,
		0,
		false,
		"(&(objectClass=computer))",
		[]string{"dn", "cn", "mS-DS-CreatorSID"},
		nil,
	)
	computerResourceBasedConstraints, err := (*l).Search(resourceBasedConstraints)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range computerResourceBasedConstraints.Entries {
		// 查询CreatorSID
		for _, i := range entry.GetAttributeValues("mS-DS-CreatorSID") {
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
					a := structResourceBasedConstraints{
						entry.DN,
						creatorSID.String(),
						i1,
					}
					listResourceBasedConstraints = append(listResourceBasedConstraints, a)
				}
			}
		}
	}
	if len(outputFile) != 0 {
		process.OutFile("\n[*] 基于资源约束委派：\n", outputFile)
		for _, i := range listResourceBasedConstraints {
			process.OutFile("\t"+i.CN+" -> creator  "+i.SID+"["+i.NAME+"]\n", outputFile)
		}
		fmt.Printf("[*] Based constraints save file to: %s\n", outputFile)
	} else {
		fmt.Println("[*] 基于资源约束委派：")
		for _, i := range listResourceBasedConstraints {
			fmt.Println("\t" + i.CN + " -> creator  " + i.SID + "[" + i.NAME + "]")
		}
	}

}
