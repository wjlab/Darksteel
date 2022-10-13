package blast

import (
	"bufio"
	"fmt"
	"github.com/ropnop/kerbrute/util"
	"os"
	"sync"
	"time"
)

func PasswordSpray(domain string, userNameList string, password string, threads int, verbose bool, outFileName string) {
	usersChan := make(chan string, threads)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(threads)

	var scanner *bufio.Scanner

	file, err := os.Open(userNameList)
	if err != nil {
		fmt.Printf("Error! %s\n", err)
		return
	}
	defer file.Close()
	scanner = bufio.NewScanner(file)

	for i := 0; i < threads; i++ {
		go makeSprayWorker(ctx, usersChan, &wg, password, verbose, outFileName, domain)
	}

	start := time.Now()

Scan:
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			break Scan
		default:
			userNameLine := scanner.Text()
			username, err := util.FormatUsername(userNameLine)
			if err != nil {
				fmt.Printf("[!] %q - %v\n", userNameLine, err.Error())
				continue
			}
			usersChan <- username
		}
	}
	close(usersChan)
	wg.Wait()

	fmt.Printf("Done! Tested logins in %.3f seconds", time.Since(start).Seconds())

	if err := scanner.Err(); err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
