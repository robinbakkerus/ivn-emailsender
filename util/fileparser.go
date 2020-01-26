package util

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/magiconair/properties"
)

// ReadEmailTemplate read
func ReadEmailTemplate() string {

	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile("C:/projects/paula-ivnmail/email-template.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	fmt.Println(text)
	return text
}

// ReadExcelFile read
func ReadExcelFile() [][]string {
	xlsx, err := excelize.OpenFile("C:/projects/paula-ivnmail/Test-dump.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("Blad1")
	return rows
}

// ReadProps read
func ReadProps() {
	p := properties.MustLoadFile("C:/projects/paula-ivnmail/config.properties", properties.UTF8)
	fmt.Println(p)
}
