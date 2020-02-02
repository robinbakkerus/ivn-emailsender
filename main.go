package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	m "ivnmailer/model"
	utl "ivnmailer/util"
	"strings"
)

func main() {
	fmt.Println("Sending emails ...")

	data := m.EmailData{}
	data = utl.ReadProps()
	data.Attachment = utl.FindAttachment(data.TemplateDir)

	goon := true
	for goon {
		data.ExcelFile = askWhichExcel(data.TemplateDir)
		data.Subject = askSubject(data.Subject)
		showData(data)

		if utl.CheckIfAlreadyProcessed(data) {
			fmt.Println("Deze excel file is vandaag al verwerkt!!")
		}
		if askBool("Okay om door te gaan? ", true) {
			utl.MakeHistory(data)
			data.TemplateBody = utl.ReadEmailTemplate(data)
			excelRows := utl.ReadExcelFile(data.TemplateDir + "/" + data.ExcelFile)
			processExcelfile(excelRows, data)
			goon = askBool("Wil je meer emails versturen ", true)
		} else {
			goon = false
		}
	}
}

func processExcelfile(rows [][]string, data m.EmailData) {
	for _, row := range rows {
		body := strings.Replace(data.TemplateBody, "{naam}", row[1], -1)
		if strings.ToUpper(row[2]) != "N" {
			utl.SendEmail(data, row[0], row[1], body)
		}

	}
}

func askWhichExcel(templateDir string) string {
	files := utl.FindExcelFiles(templateDir)

	nr := 1
	fmt.Println("Welke Excel :")
	for _, f := range files {
		fmt.Println(strconv.Itoa(nr) + " : " + f.Name + " (" + strconv.Itoa(f.Size) + ")")
		nr++
	}

	reader := bufio.NewReader(os.Stdin)
	textInput, _ := reader.ReadString('\r')

	nr, err := strconv.Atoi(strings.TrimRight(textInput, "\r"))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	useExcel := files[nr-1]
	return useExcel.Name
}

func askBool(askmsg string, defval bool) bool {
	reader := bufio.NewReader(os.Stdin)

	defTxt := ""
	if defval {
		defTxt = "J"
	} else {
		defTxt = "N"
	}

	msg := askmsg + " (" + defTxt + ") > "
	fmt.Println(msg)
	textInput, _ := reader.ReadString('\r')
	textInput = strings.ToUpper(strings.TrimRight(textInput, "\r"))
	if len(textInput) == 0 {
		textInput = defTxt
	}

	return strings.HasPrefix(textInput, "J")
}

func showData(data m.EmailData) {
	fmt.Println()
	fmt.Println("Excel file = " + data.ExcelFile)
	fmt.Println("Onderwerp  = " + data.Subject)
	fmt.Println("Attachment = " + data.Attachment)
	fmt.Println()
}

func askSubject(defSubject string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Geef onderwerp (" + defSubject + ") > ")
	textInput, _ := reader.ReadString('\r')
	r := strings.TrimRight(textInput, "\r")
	if len(r) == 0 {
		return defSubject
	}
	return r

}
