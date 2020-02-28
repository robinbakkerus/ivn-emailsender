package util

import (
	"fmt"

	//"strings"
	//"time"
	m "jrb/ivn-emailsender/mailer/model"
	u "jrb/ivn-emailsender/mailer/util"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func openExcel(filename string) *excelize.File {
	xlsx, err := excelize.OpenFile(filename)
	if err != nil {
		panic(err)
	}
	return xlsx
}

// UpdateNewAndModifiedRecords ...
func UpdateNewAndModifiedRecords(data m.EmailData, dumpFilename string) {

	dumpfile, err := excelize.OpenFile(data.TemplateDir + "/" + dumpFilename)
	if err != nil {
		panic(err)
	}
	dumpRecs := dumpfile.GetRows(dumpfile.GetSheetName(1))

	mailFilename := data.TemplateDir + "/" + data.ExcelFile
	mailfile, err := excelize.OpenFile(mailFilename)
	if err != nil {
		panic(err)
	}
	mailRecs := mailfile.GetRows(mailfile.GetSheetName(1))

	processUpdatedAndNew(dumpRecs, mailRecs, mailfile)
	processDeleted(dumpRecs, mailRecs, mailfile)

	if err := mailfile.SaveAs(mailFilename); err != nil {
		fmt.Println(err)
	}

}

func processUpdatedAndNew(dumpRecs [][]string, mailRecs [][]string, mailfile *excelize.File) {

	maxRow := len(mailRecs)

	updStyle, _ := mailfile.NewStyle(`{"fill":{"type":"pattern","color":["#f2f542"],"pattern":1}}`)
	newStyle, _ := mailfile.NewStyle(`{"fill":{"type":"pattern","color":["#37b03d"],"pattern":1}}`)

	for rowNr, row := range dumpRecs {
		id := row[0]
		name := row[1]
		mail := row[2]

		if rowNr > 4 {
			mailRow, mailRowNr := corrRow(id, mailRecs)
			if mailRow != nil {
				if name != mailRow[1] || mail != mailRow[2] {
					fmt.Printf("Aangepast: %v - %v id= %v : %v %v => %v %v \n", rowNr, mailRowNr, id, mailRow[1], mailRow[2], name, mail)
					mailfile.SetCellValue(mailfile.GetSheetName(1), getCellname("B", mailRowNr+1), name)
					mailfile.SetCellValue(mailfile.GetSheetName(1), getCellname("C", mailRowNr+1), mail)
					mailfile.SetCellStyle(mailfile.GetSheetName(1), getCellname("B", mailRowNr+1), getCellname("C", mailRowNr+1), updStyle)
				}
			} else {
				fmt.Printf("Nieuw record: %v - %v id= %v : %v %v \n", rowNr, mailRowNr, id, name, mail)
				fmt.Printf("max = %v \n", maxRow)
				mailfile.SetCellValue(mailfile.GetSheetName(1), getCellname("A", maxRow+1), id)
				mailfile.SetCellValue(mailfile.GetSheetName(1), getCellname("B", maxRow+1), name)
				mailfile.SetCellValue(mailfile.GetSheetName(1), getCellname("C", maxRow+1), mail)
				mailfile.SetCellStyle(mailfile.GetSheetName(1), getCellname("B", maxRow+1), getCellname("D", maxRow+1), newStyle)
				maxRow = maxRow + 1
			}
		}
	}
}

func processDeleted(dumpRecs [][]string, mailRecs [][]string, mailfile *excelize.File) {

	delStyle, _ := mailfile.NewStyle(`{"fill":{"type":"pattern","color":["#cf481b"],"pattern":1}}`)

	for rowNr, row := range mailRecs {
		id := row[0]
		name := row[1]
		mail := row[2]

		if rowNr > 4 {
			dumpRow, _ := corrRow(id, dumpRecs)
			if dumpRow == nil && string(id) != "0" {
				fmt.Printf("Verwijderd record: %v id = %v : %v %v \n", rowNr, id, name, mail)
				mailfile.SetCellValue(mailfile.GetSheetName(1), getCellname("E", rowNr+1), "D")
				mailfile.SetCellValue(mailfile.GetSheetName(1), getCellname("F", rowNr+1), "D")
				mailfile.SetCellStyle(mailfile.GetSheetName(1), getCellname("B", rowNr+1), getCellname("C", rowNr+1), delStyle)
			}
		}

	}
}

func corrRow(dumpID string, mailRecs [][]string) ([]string, int) {
	for rownum, row := range mailRecs {
		if string(dumpID) == string(row[0]) {
			return row, rownum
		}
	}
	return nil, 0
}

func getCellname(x string, y int) string {
	return x + u.ToStr(y)
}
