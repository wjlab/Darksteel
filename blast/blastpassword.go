package blast

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func BlastPassword(domain string, username string, passwordList string, threads int, verbose bool, outFileName string) {
	passwordsChan := make(chan string, threads)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(threads)

	var scanner *bufio.Scanner

	file, err := os.Open(passwordList)
	if err != nil {
		fmt.Printf("Error! %s\n", err)
		return
	}
	defer file.Close()
	scanner = bufio.NewScanner(file)

	for i := 0; i < threads; i++ {
		go makeBruteWorker(ctx, passwordsChan, &wg, username, verbose, outFileName, domain)
	}

	start := time.Now()

	var password string
Scan:
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			break Scan
		default:
			password = scanner.Text()
			passwordsChan <- password
		}
	}
	close(passwordsChan)
	wg.Wait()

	fmt.Printf("Done! Tested logins in %.3f seconds", time.Since(start).Seconds())

	if err := scanner.Err(); err != nil {
		fmt.Printf(fmt.Sprintf(err.Error()))
	}

}
