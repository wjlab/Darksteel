package main

import (
	"flag"
	"fmt"
	"os"
	"tdk/blast"
	"tdk/conf"
	"tdk/kerberos"
	"tdk/ldap"
)

var (
	target        string
	domain        string
	user          string
	userFile      string
	password      string
	passwordFile  string
	allDelegate   string
	searchValue   string
	outputContent string
	integrate     string
	roastModule   string
	targetUser    string
	format        string
	enctype       string
	outputFile    string
	ldapSizeLimit int
	allLdap       bool
	ticket        string
	fuzz          string
	blastModule   string
	threads       int
	verbose       bool
	userPass      string
)

type MyFlagSet struct {
	*flag.FlagSet
	cmdComment string // 二级子命令本身的注释
}

func main() {
	// 6 本地账号认证
	// ldap
	ldapCmd := &MyFlagSet{
		FlagSet:    flag.NewFlagSet("ldap", flag.ExitOnError),
		cmdComment: "Interact with LDAP server",
	}
	ldapCmd.StringVar(&target, "dc", "", "* Please enter the IP of the domain control")
	ldapCmd.StringVar(&domain, "domain", "", "* Please enter the domain name")
	ldapCmd.StringVar(&user, "user", "", "* Username in the domain")
	ldapCmd.StringVar(&password, "pass", "", "* The corresponding password or hash user")
	ldapCmd.StringVar(&allDelegate, "w", "", "all \n  all delegate information \nuw \n  unconstrained delegation information \ncw \n  Constraint appointment information \nbw \n  Resource-based constraint delegation")
	ldapCmd.StringVar(&searchValue, "f", "", "Customize the field of LDAP")
	ldapCmd.StringVar(&outputContent, "n", "", "The field to query, you can write multiple")
	ldapCmd.StringVar(&integrate, "m", "", "user \n  Query all users in the domain \ncomputer \n  Query all computers in the domain \nscomputer \n  Query survival computer \ndc \n  Query all domain controls in the domain \nspn \n  Query all SPN in the domain \nou \n  All OU in the query domain \nmssql \n  Query all mssql services in the domain \nasreproast \n  Query users in the domain who can use as-rep roast \nmaq \n  Query the value of maq in the domain")
	ldapCmd.IntVar(&ldapSizeLimit, "ldapSizeLimit", 0, "Query LDAP maximum number (default 0)")
	ldapCmd.StringVar(&outputFile, "o", "", "Output file position, default current directory")
	ldapCmd.BoolVar(&allLdap, "all", false, "Query all content")
	ldapCmd.StringVar(&fuzz, "fuzz", "", "vague query content")

	// kerberos
	kerberosCmd := &MyFlagSet{
		FlagSet:    flag.NewFlagSet("kerberos", flag.ExitOnError),
		cmdComment: "Do some Kerberos stuff",
	}
	kerberosCmd.StringVar(&target, "dc", "", "* Please enter the IP of the domain control")
	kerberosCmd.StringVar(&domain, "domain", "", "* Please enter the domain name")
	kerberosCmd.StringVar(&user, "user", "", "* Username in the domain")
	kerberosCmd.StringVar(&password, "pass", "", "* The corresponding password or hash user")
	kerberosCmd.StringVar(&roastModule, "m", "", "asreproast \n  as-rep roast attack\nkerberoast \n  kerberoasting attack")
	kerberosCmd.StringVar(&targetUser, "tuser", "", "Enter the user to be utilized")
	kerberosCmd.StringVar(&format, "format", "hashcat", "format Output hash as John the Ripper or Hashcat format")
	kerberosCmd.StringVar(&enctype, "enctype", "rc4", "enctype Encryption type: rc4, aes128 or aes256")
	kerberosCmd.StringVar(&outputFile, "o", "", "Output file position, default current directory")
	kerberosCmd.StringVar(&ticket, "ticket", "", "Using ticket authentication, enter the path of the ticket")
	kerberosCmd.IntVar(&ldapSizeLimit, "ldapsizelimit", 0, "Query LDAP maximum number (default 0)")

	// blast
	blastCmd := &MyFlagSet{
		FlagSet:    flag.NewFlagSet("blast", flag.ExitOnError),
		cmdComment: "Blasting Domain User",
	}
	blastCmd.StringVar(&target, "dc", "", "* Please enter the IP of the domain control")
	blastCmd.StringVar(&domain, "domain", "", "* Please enter the domain name")
	blastCmd.StringVar(&user, "user", "", "Username in the domain")
	blastCmd.StringVar(&userFile, "userfile", "", "User dictionary")
	blastCmd.StringVar(&password, "pass", "", "Password in the domain")
	blastCmd.StringVar(&passwordFile, "passfile", "", "Password dictionary")
	blastCmd.StringVar(&userPass, "upfile", "", "The dictionary corresponding to the user name and password is split by:")
	blastCmd.IntVar(&threads, "t", 20, "Number of burst threads")
	blastCmd.StringVar(&blastModule, "m", "", "userenum -userfile user.txt\n  User enumeration\npassspray -userfile user.txt -pass password\n  Password spraying\nblastpass -user username -passfile password.txt\n  Single user burst password\nuserpass -upfile userpass.txt\n  User password combinations explode")
	blastCmd.StringVar(&outputFile, "o", "", "Output file position, default current directory")
	blastCmd.BoolVar(&verbose, "v", false, "Whether a failure message is displayed")

	// 用 map 保存所有的二级子命令，方便快速查找
	subcommands := map[string]*MyFlagSet{
		ldapCmd.Name():     ldapCmd,
		kerberosCmd.Name(): kerberosCmd,
		blastCmd.Name():    blastCmd,
	}

	// 整个命令行的帮助信息
	useAge := func() {
		conf.Banner()
		fmt.Println("Available Commands:")
		fmt.Println("  darksteel ldap [parameter]")
		fmt.Println("  darksteel kerberos [parameter]")
		fmt.Println("  darksteel blast [parameter]\n")
		for _, v := range subcommands {
			fmt.Printf("%s %s\n", v.Name(), v.cmdComment)
			v.PrintDefaults() // 使用 flag 库自带的格式输出子命令的选项帮助信息
			fmt.Println()
		}
		os.Exit(0)
	}

	// 即没有输入子命令
	if len(os.Args) < 2 {
		useAge()
	}

	// 第二个参数必须是我们支持的子命令
	cmd := subcommands[os.Args[1]]
	if cmd == nil {
		useAge()
	}

	// 注意这里是 cmd.Parse 不是 flag.Parse，且值是 Args[2:]
	err := cmd.Parse(os.Args[2:])
	if err != nil {
		return
	}
	if os.Args[1] == "kerberos" {
		if len(os.Args) < 10 {
			useAge()
		} else {
			conf.Banner()
			kerberos.KerberosInit(domain, target, password, user, roastModule, targetUser, format, enctype, outputFile, ldapSizeLimit, ticket)
		}
	} else if os.Args[1] == "ldap" {
		if len(os.Args) < 11 {
			useAge()
		} else {
			conf.Banner()
			ldap.LdapInit(domain, target, password, user, allDelegate, searchValue, integrate, outputContent, ldapSizeLimit, outputFile, allLdap, fuzz)
		}
	} else if os.Args[1] == "blast" {
		if len(os.Args) < 7 {
			useAge()
		} else {
			conf.Banner()
			blast.SetupSession(domain, target, user, userFile, password, passwordFile, userPass, threads, blastModule, outputFile, verbose)
		}
	} else {
		useAge()
	}
}
