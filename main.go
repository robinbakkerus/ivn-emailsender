package main

import (
	"fmt"
	"strings"
	utl "ivnmailer/util"
)

func main() {
	template := utl.ReadEmailTemplate()
	fmt.Println(template)

	utl.ReadProps()
	excelRows := utl.ReadExcelFile()

	processExcelfile(excelRows, template)
	fmt.Println("done")
}

func processExcelfile(rows [][]string, template string) {
	for _, row := range rows {
		body := strings.Replace(template, "{naam}", row[8], -1)
		bodyBytes := []byte(body)
		utl.SendEmail(row[7], bodyBytes)
		fmt.Print(body, "\t")
		fmt.Println()
	}
}

