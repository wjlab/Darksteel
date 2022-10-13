package blast

import (
	"bufio"
	"fmt"
	"github.com/ropnop/kerbrute/util"
	"os"
	"sync"
	"time"
)

func UserPass(domain string, userPass string, threads int, verbose bool, outFileName string) {
	combosChan := make(chan [2]string, threads)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(threads)

	var scanner *bufio.Scanner

	file, err := os.Open(userPass)
	if err != nil {
		fmt.Printf("Error! %s\n", err)
		return
	}
	defer file.Close()
	scanner = bufio.NewScanner(file)

	for i := 0; i < threads; i++ {
		go makeBruteComboWorker(ctx, combosChan, &wg, verbose, outFileName, domain)
	}

	start := time.Now()

Scan:
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			break Scan
		default:
			comboLine := scanner.Text()
			if comboLine == "" {
				continue
			}
			username, password, err := util.FormatComboLine(comboLine)
			if err != nil {
				fmt.Printf("[!] Skipping: %q - %v\n", comboLine, err.Error())
				continue
			}
			combosChan <- [2]string{username, password}
		}
	}
	close(combosChan)
	wg.Wait()

	fmt.Printf("Done! Tested logins in %.3f seconds", time.Since(start).Seconds())

	if err := scanner.Err(); err != nil {
		fmt.Printf(err.Error())
	}
}
