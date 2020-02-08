package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"

	m "ivnmailer/model"
	utl "ivnmailer/util"
	"strings"
)

func main() {
	fmt.Println("Sending emails ...")

	data := m.EmailData{}
	data = utl.ReadProps()
	data.Attachment = utl.FindAttachment(data.TemplateDir)
	excelHeaders := utl.ReadExcelFileHeaders(data.TemplateDir + "/" + "Test-dump.xlsx")

	goon := true
	for goon {
		data.MailListIdx, data.MailList = askMailingList(excelHeaders)
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
			goon = !askBool("Wil je helemaal stoppen? ", true)
		}
	}
}

func processExcelfile(rows [][]string, data m.EmailData) {
	for _, row := range rows {
		body := strings.Replace(data.TemplateBody, "{naam}", row[1], -1)
		if !skip(data, row) {
			utl.SendEmail(data, row[0], row[1], body)
		}

	}
}

func skip(data m.EmailData, row []string) bool {
	return strings.ToUpper(row[data.MailListIdx]) != "X"
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

	textInput, _ := reader.ReadString(delim())
	textInput = strings.ToUpper(trimDelim(textInput))
	if len(textInput) == 0 {
		textInput = defTxt
	}

	return strings.HasPrefix(textInput, "J")
}

func showData(data m.EmailData) {
	fmt.Println()
	fmt.Println("Mail list = " + data.MailList)
	fmt.Println("Onderwerp  = " + data.Subject)
	fmt.Println("Attachment = " + data.Attachment)
	fmt.Println()
}

func askSubject(defSubject string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Geef onderwerp (" + defSubject + ") > ")
	textInput, _ := reader.ReadString(delim())
	r := trimDelim(textInput)
	if len(r) == 0 {
		return defSubject
	}
	return r
}

func delim() byte {
	if runtime.GOOS == "windows" {
		return '\r'
	} else {
		return '\n'
	}
}

func trimDelim(s string) string {
	if runtime.GOOS == "windows" {
		return strings.TrimRight(s, "\r")
	} else {
		return strings.TrimRight(s, "\n")
	}
}

func askMailingList(headers []string) (int, string) {
	fmt.Println("Welke mailinglist :")
	for index, hdr := range headers {
		if index > 1 && len(hdr) > 0 {
			nr := index - 2
			fmt.Printf("%v: %v\n", nr, hdr)
		}
	}

	var askNr int
	fmt.Scan(&askNr)
	return askNr + 2, headers[askNr+2]
}
