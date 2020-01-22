package util

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/magiconair/properties"
	"github.com/360EntSecGroup-Skylar/excelize"
)

// ReadEmailTemplate read
func ReadEmailTemplate() []string {
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

// ReadExcelFile read
func ReadExcelFile() []string {
	xlsx, err := excelize.OpenFile("c:/temp/email-template.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    // Get all the rows in the Sheet1.
    rows := xlsx.GetRows("Blad1")
    for _, row := range rows {
        for _, colCell := range row {
            fmt.Print(colCell, "\t")
        }
        fmt.Println()
	}
	
	return rows
}


func ReadProps() {
	p := properties.MustLoadFile("c:/temp/config.properties", properties.UTF8)
}