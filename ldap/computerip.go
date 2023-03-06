package ldap

import (
	"bufio"
	"darksteel/conf"
	"darksteel/process"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/miekg/dns"
	"os"
	"time"
)

func fileRead(file string) *bufio.Scanner {
	f, err := os.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		return nil
	}
	datas := bufio.NewScanner(f)
	return datas
}

func SearchComputerIps(l **ldap.Conn, domain string, listDomain string, ldapSizeLimit int, ip string, outputFile string, file string) {
	var listComputerIps []string
	if len(file) != 0 {
		data := fileRead(file)
		for data.Scan() {
			listComputerIps = append(listComputerIps, data.Text())
		}
	} else {
		searchComputerIps := ldap.NewSearchRequest(listDomain,
			ldap.ScopeWholeSubtree,
			ldap.NeverDerefAliases,
			ldapSizeLimit,
			0,
			false,
			conf.LdapQueries["computers"],
			[]string{"name"},
			nil)
		//searchComputerIp, err := (*l).Search(searchComputerIps)
		searchComputerIp, err := (*l).SearchWithPaging(searchComputerIps, 10000)
		if err != nil {
			fmt.Println(err)
		}
		for _, entry := range searchComputerIp.Entries {
			listComputerIps = append(listComputerIps, entry.GetAttributeValue("name"))
		}
	}
	if len(outputFile) != 0 {
		process.OutFile(fmt.Sprintf("[*] ComputerIp save file to:  %s\n"), outputFile)
		for _, j := range listComputerIps {
			c := dns.Client{
				Timeout: 5 * time.Second,
			}
			m := dns.Msg{}
			m.SetQuestion(fmt.Sprintf("%s.%s.", j, domain), dns.TypeA)
			r, _, err := c.Exchange(&m, ip+":53")
			if err != nil {
				fmt.Println("dns error")
				return
			}

			for _, ans := range r.Answer {
				record, isType := ans.(*dns.A)
				if isType {
					process.OutFile(fmt.Sprintf("\t%s   ————> A: %s\n", j, record.A), outputFile)
				}
			}
		}
		fmt.Printf("[*] ComputerIp save file to: %s\n", outputFile)
	} else {
		fmt.Printf("[*] Computer correspondence iP:\n")
		for _, j := range listComputerIps {
			c := dns.Client{
				Timeout: 5 * time.Second,
			}
			m := dns.Msg{}
			m.SetQuestion(j+"."+domain+".", dns.TypeA)
			r, _, err := c.Exchange(&m, ip+":53")
			if err != nil {
				fmt.Println("dns error")
				return
			}
			for _, ans := range r.Answer {
				record, isType := ans.(*dns.A)
				if isType {
					fmt.Println("\t"+j+"  ————> A:", record.A)
				}
			}
		}
	}
}
