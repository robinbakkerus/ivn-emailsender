package main

import (
	"fmt"
	utl "ivnmailer/util"
	"strings"
)

// Data d
type Data struct {
	template string
	smtpUser string
	smtpPwd  string
	sendFrom string
}

func main() {
	data := new(Data)
	data.template = utl.ReadEmailTemplate()

	data.smtpUser, data.smtpPwd, data.sendFrom = utl.ReadProps()
	excelRows := utl.ReadExcelFile()

	processExcelfile(excelRows, data)
	fmt.Println("done")
}

func processExcelfile(rows [][]string, d *Data) {
	for _, row := range rows {
		body := strings.Replace(d.template, "{naam}", row[8], -1)
		// bodyBytes := []byte(body)
		// utl.SendEmail(data.smtpUser, data.smtpPwd, data.sendFrom, row[7], bodyBytes)
		// tos := []string {row[8]}
		// utl.SendGMail(d.smtpPwd, d.smtpPwd, d.smtpUser, tos, body,
		// 	"C:/projects/paula-ivnmail", "Test-dump.xlsx")

		// smtpUser string,
		// smtpPwd string,
		// sendTo string,
		// subject string,
		// message []byte,
		// attFilename

		utl.SendEmail(d.smtpUser, d.smtpPwd, row[0], "Dit is een test", body,
			"C:/projects/paula-ivnmail/Test-dump.xlsx")
	}
}
