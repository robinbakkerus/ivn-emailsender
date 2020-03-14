package main

import (
	"fmt"
	m "jrb/ivn-emailsender/mailer/model"
	u "jrb/ivn-emailsender/mailer/util"
	"strconv"
	"strings"
)

func main() {
	data := m.EmailData{}
	data = u.ReadProps()
	fmt.Println("Sending emails ..., dryrun = " + strconv.FormatBool((data.DryRun)))
	data.Attachments = u.FindAttachments(data.TemplateDir)
	excelHeaders, sendtoCounts := u.ReadExcelFileHeaders(data.TemplateDir + "/" + data.ExcelFile)

	goon := true
	for goon {
		data.MailListIdx, data.MailList = askMailingList(excelHeaders, sendtoCounts)
		data.Subject = u.AskString(data.Subject)
		showData(data)

		// if u.CheckIfAlreadyProcessed(data) {
		// 	fmt.Println("Deze excel file is vandaag al verwerkt!!")
		// }

		if len(data.Attachments) == 0 {
			fmt.Println("LET OP, Ik heb geen attachments gevonden in " + data.TemplateDir + m.AttachmentSubdir)
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
	sendMap := make(map[string]bool)
	cnt := 0

	fmt.Println("Sending emails to ...")

	for _, row := range rows {
		mailAddr := row[m.EMAIL]
		aanhef := row[m.AANHEF]

		if len(aanhef) == 0 {
			aanhef = data.Aanhef
		}

		if !sendMap[mailAddr] {
			sendMap[mailAddr] = true
			body := strings.Replace(data.TemplateBody, m.ReplaceAanhef, aanhef, -1)
			if !skip(data, row) {
				u.SendEmail(data, mailAddr, aanhef, body)
				cnt++
			}
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
	fmt.Println("Standaard aanhef  = " + data.Aanhef)
	fmt.Println("Attachments ")
	for i, f := range data.Attachments {
		fmt.Println(i, f.Name())
	}
	fmt.Println()
}

func askMailingList(headers []string, counts []int) (int, string) {
	fmt.Println("Welke mailinglist :")
	for index, hdr := range headers {
		if index > m.AANHEF && len(hdr) > 0 {
			fmt.Printf("%v: %v \t%v \n", index, hdr, counts[index-(m.AANHEF+1)])
		}
	}

	var askNr int
	fmt.Scan(&askNr)
	return askNr, headers[askNr]
}
