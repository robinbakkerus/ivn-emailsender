package main

import (
	"fmt"

	m "ivnmailer/model"
	u "ivnmailer/util"
	"strings"
)

func main() {
	fmt.Println("Sending emails ...")

	data := m.EmailData{}
	data = u.ReadProps()
	data.Attachments = u.FindAttachments(data.TemplateDir)
	excelHeaders, sendtoCounts := u.ReadExcelFileHeaders(data.TemplateDir + "/" + data.ExcelFile)

	goon := true
	for goon {
		data.MailListIdx, data.MailList = askMailingList(excelHeaders, sendtoCounts)
		data.Subject = u.AskString(data.Subject)
		showData(data)

		if u.CheckIfAlreadyProcessed(data) {
			fmt.Println("Deze excel file is vandaag al verwerkt!!")
		}
		if u.AskBool("Okay om door te gaan? ", true) {
			u.MakeHistory(data)
			data.TemplateBody = u.ReadEmailTemplate(data)
			excelRows := u.ReadExcelFile(data.TemplateDir + "/" + data.ExcelFile)
			processExcelfile(excelRows, data)
			goon = u.AskBool("Wil je meer emails versturen ", true)
		} else {
			goon = !u.AskBool("Wil je helemaal stoppen? ", true)
		}
	}
}

func processExcelfile(rows [][]string, data m.EmailData) {
	sendMap := make(map[string]int)
	cnt := 0

	for _, row := range rows {
		mailAddr := row[0]
		name := row[1]
		if sendMap[mailAddr] == 0 {
			sendMap[mailAddr] = 1
			body := strings.Replace(data.TemplateBody, "{naam}", name, -1)
			if !skip(data, row) {
				u.SendEmail(data, mailAddr, name, body)
			}
			cnt++
		}
	}
	fmt.Println("In totaal " + u.ToStr(cnt) + " verstuurd")
}

func skip(data m.EmailData, row []string) bool {
	return strings.ToUpper(row[data.MailListIdx]) != "X"
}

func showData(data m.EmailData) {


	fmt.Println()
	fmt.Println("Mail list = " + data.MailList)
	fmt.Println("Onderwerp  = " + data.Subject)
	fmt.Println("Attachments ")
	for i, f := range data.Attachments {
		fmt.Println(i, f.Name())
	}
	fmt.Println()
}

func askMailingList(headers []string, counts []int) (int, string) {
	fmt.Println("Welke mailinglist :")
	for index, hdr := range headers {
		if index > 1 && len(hdr) > 0 {
			nr := index - 2
			fmt.Printf("%v: %v \t%v \n", nr, hdr, counts[index-2])
		}
	}

	var askNr int
	fmt.Scan(&askNr)
	return askNr + 2, headers[askNr+2]
}
