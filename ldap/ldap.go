package ldap

import (
	"darksteel/process"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"log"
)

func LdapInit(domain string, target string, password string, user string, allDelegate string, searchValue string, integrate string, outputContent string, ldapSizeLimit int, outputFile string, allLdap bool, fuzz string) {
	var listDomain string
	listDomain = process.DcFormatConversion(domain)

	//连接
	l, err := ldap.DialURL("ldap://" + target + ":389")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	//判断hash验证
	if len(password) != 32 {
		_, err = l.SimpleBind(&ldap.SimpleBindRequest{
			Username: fmt.Sprintf("CN=%s,cn=Users,%s", user, listDomain),
			Password: password,
		})
		if err != nil {
			log.Fatalf("Failed to bind: %s\n", err)
		}
	} else {
		err = l.NTLMBindWithHash(target, user, password)
		if err != nil {
			log.Fatalf("Failed to bind: %s\n", err)
		}
	}

	//判断输出委派信息
	switch {
	case allDelegate == "all":
		SearchUnconstrained(&l, listDomain, ldapSizeLimit, outputFile)
		SearchConstrained(&l, listDomain, ldapSizeLimit, outputFile)
		SearchBasedConstraints(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case allDelegate == "uw":
		SearchUnconstrained(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case allDelegate == "cw":
		SearchConstrained(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case allDelegate == "bw":
		SearchBasedConstraints(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case len(searchValue) != 0:
		SearchLdap(&l, listDomain, searchValue, outputContent, ldapSizeLimit)
		break
	case integrate == "user":
		SearchUsers(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case integrate == "ou":
		SearchOU(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case integrate == "mssql":
		SearchMsSqlServer(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case integrate == "maq":
		SearchMaq(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case integrate == "dc":
		SearchDc(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case integrate == "computer":
		SearchComputers(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case integrate == "asreproast":
		SearchRoast(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case integrate == "spn":
		SearchSpn(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case integrate == "scomputer":
		SearchSurvivalComputer(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case integrate == "admins":
		SearchAdmins(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case integrate == "enterprise":
		SearchEnterprise(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case integrate == "exchangecomputer":
		SearchExchangeComputer(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case integrate == "exchangesystem":
		SearchExchangeTrustedSubsystem(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case integrate == "exchangeorgmanager":
		SearchExchangeOrganizationManagement(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case integrate == "trustdomain":
		SearchTrustDomain(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case integrate == "adminsdholder":
		SearchAdminSDHolder(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case integrate == "sidhistory":
		SearchSIDHistory(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case allLdap:
		SearchUsers(&l, listDomain, ldapSizeLimit, outputFile)
		SearchAdmins(&l, listDomain, ldapSizeLimit, outputFile)
		SearchAdminSDHolder(&l, listDomain, ldapSizeLimit, outputFile)
		SearchSIDHistory(&l, listDomain, ldapSizeLimit, outputFile)
		SearchEnterprise(&l, listDomain, ldapSizeLimit, outputFile)
		SearchOU(&l, listDomain, ldapSizeLimit, outputFile)
		SearchMsSqlServer(&l, listDomain, ldapSizeLimit, outputFile)
		SearchMaq(&l, listDomain, ldapSizeLimit, outputFile)
		SearchDc(&l, listDomain, ldapSizeLimit, outputFile)
		SearchTrustDomain(&l, listDomain, ldapSizeLimit, outputFile)
		SearchComputers(&l, listDomain, ldapSizeLimit, outputFile)
		SearchSurvivalComputer(&l, listDomain, ldapSizeLimit, outputFile)
		SearchExchangeComputer(&l, listDomain, ldapSizeLimit, outputFile)
		SearchExchangeTrustedSubsystem(&l, listDomain, ldapSizeLimit, outputFile)
		SearchExchangeOrganizationManagement(&l, listDomain, ldapSizeLimit, outputFile)
		SearchRoast(&l, listDomain, ldapSizeLimit, outputFile)
		SearchUnconstrained(&l, listDomain, ldapSizeLimit, outputFile)
		SearchConstrained(&l, listDomain, ldapSizeLimit, outputFile)
		SearchBasedConstraints(&l, listDomain, ldapSizeLimit, outputFile)
		SearchSpn(&l, listDomain, ldapSizeLimit, outputFile)
		break
	case len(fuzz) != 0:
		Fuzzy(&l, listDomain, ldapSizeLimit, outputFile, fuzz)
		break
	}
}
