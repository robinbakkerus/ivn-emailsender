package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readEmailTemplate() []string {
	file, err := os.Open("c:/temp/email-template.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}

	return txtlines
}
