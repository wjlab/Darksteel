package process

import (
	"fmt"
	"io"
	"os"
)

func OutFileRoast(content string, outputFile string) error {
	f, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("cannot create hash file: %s", err)
	}
	defer f.Close()

	_, err = f.Write([]byte(content))
	if err != nil {
		return err
	}

	fmt.Printf("[*] Save file to: %s\n", outputFile)
	return nil
}

func OutFile(content string, outputFile string) error {
	f, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}
	if _, err = io.WriteString(f, content); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}
	defer f.Close()
	return err
}
