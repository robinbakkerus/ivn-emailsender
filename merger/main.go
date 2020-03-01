package main

import (
	"fmt"
	"os"
	"strings"

	utl "jrb/ivn-emailsender/merger/util"

	u "jrb/ivn-emailsender/mailer/util"

	m "jrb/ivn-emailsender/mailer/model"
)

func main() {
	data := m.EmailData{}
	data = u.ReadProps()
	if data.DryRun {
		fmt.Println("\nDRYRUN: Merging IVN dumpfile with mailing list ...")
	} else {
		fmt.Println("\nMerging IVN dumpfile with mailing list ...")
	}

	dumpfile := askDumpFile(utl.FindDumpFiles(data), data.DownloadDir)
	utl.UpdateNewAndModifiedRecords(data, dumpfile)
}

func skip(data m.EmailData, row []string) bool {
	return strings.ToUpper(row[data.MailListIdx]) != "X"
}

func showData(data m.EmailData) {
	fmt.Println("Onderwerp  = " + data.Subject)
	fmt.Println()
}

func askDumpFile(filenames []string, downloadDir string) string {
	if len(filenames) > 0 {
		fmt.Println("Welke dump file gebruiken :")
		for index, name := range filenames {
			fmt.Printf("%v: %v  \n", index, name)
		}

		var askNr int
		fmt.Scan(&askNr)
		return filenames[askNr]
	} else {
		fmt.Println("Kan geen .csv files vinden in " + downloadDir)
		os.Exit(3)
		return ""
	}

}
