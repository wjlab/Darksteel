package blast

import (
	"context"
	"fmt"
)

var (
	kSession KerbruteSession
	// 线程
	ctx, cancel = context.WithCancel(context.Background())
)

func SetupSession(domain string, domainController string, username string, userFile string, password string, passwordFile string, userPass string, threads int, blastModule string, outFileName string, verbose bool) {
	k, err := NewKerbruteSession(domain, domainController)
	if err != nil {
		fmt.Println("请输入-h或--help查看帮助信息")
		fmt.Printf("[!] %s", err)
		return
	}
	kSession = k
	switch {
	case blastModule == "userenum":
		if userFile == "" {
			fmt.Println("[!] Please enter -U to specify the user name dictionary")
		} else if domain == "" {
			fmt.Println("[!] Please enter -n domain name")
		} else {
			UserEnum(domain, userFile, threads, verbose, outFileName)
		}
		break
	case blastModule == "passspray":
		if userFile == "" {
			fmt.Println("[!] Please enter -U to specify the user name dictionary")
		} else if password == "" {
			fmt.Println("[!] Please enter -p to specify the password for spraying")
		} else {
			PasswordSpray(domain, userFile, password, threads, verbose, outFileName)
		}
		break
	case blastModule == "blastpass":
		if username == "" {
			fmt.Println("[!] Please enter -u to specify the user you want to blow up")
		} else if passwordFile == "" {
			fmt.Println("[!] Please enter -P to specify a password dictionary")
		} else {
			BlastPassword(domain, username, passwordFile, threads, verbose, outFileName)
		}
		break
	case blastModule == "userpass":
		if userPass == "" {
			fmt.Println("[!] Please enter -F to specify the password dictionary for the user name")
		} else {
			UserPass(domain, userPass, threads, verbose, outFileName)
		}
		break
	default:
		fmt.Println("[!] Please enter -m to specify the module you want to use")
	}
}
