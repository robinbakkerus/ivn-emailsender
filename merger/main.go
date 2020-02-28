package main

import (
	"fmt"

	"strings"

	utl "jrb/ivn-emailsender/merger/util"

	u "jrb/ivn-emailsender/mailer/util"

	m "jrb/ivn-emailsender/mailer/model"
)

func main() {
	fmt.Println("Merging IVN dumpfile with mailing list ...")

	data := m.EmailData{}
	data = u.ReadProps()
	dumpfile := askDumpFile(utl.FindDumpFiles(data))
	utl.UpdateNewAndModifiedRecords(data, dumpfile)
}

func skip(data m.EmailData, row []string) bool {
	return strings.ToUpper(row[data.MailListIdx]) != "X"
}

func showData(data m.EmailData) {
	fmt.Println("Onderwerp  = " + data.Subject)
	fmt.Println()
}

func askDumpFile(filenames []string) string {
	fmt.Println("Welke dump file gebruiken :")
	for index, name := range filenames {
		fmt.Printf("%v: %v  \n", index, name)
	}

	var askNr int
	fmt.Scan(&askNr)
	return filenames[askNr]
}
