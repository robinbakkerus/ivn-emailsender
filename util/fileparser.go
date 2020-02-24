package util

import (
	"fmt"
	"io/ioutil"
	m "ivnmailer/model"
	"log"
	"os"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/magiconair/properties"
)

// ReadEmailTemplate read
func ReadEmailTemplate(data m.EmailData) string {

	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(data.TemplateDir + "/" + data.TemplateName)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	return text
}

// ReadExcelFile read
func ReadExcelFile(filename string) [][]string {
	xlsx, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("Blad1")

	rows[0] = rows[len(rows)-1] // Copy last element to index i.
	rows[len(rows)-1] = nil     // Erase last element (write zero value).
	rows = rows[:len(rows)-1]   // Truncate slice.

	return rows
}

// ReadExcelFileHeaders ..
func ReadExcelFileHeaders(filename string) ([]string, []int) {
	xlsx, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("Blad1")

	counts := make([]int, 0)

	for i := 2; i < len(rows[0]); i++ {
		counts = append(counts, getSenttoCount(rows, i))
	}

	return rows[0], counts
}

func getSenttoCount(rows [][]string, col int) int {
	cnt := 0
	for _, row := range rows {
		if "X" == strings.ToUpper(row[col]) {
			cnt++
		}
	}
	return cnt
}

// ReadProps read
func ReadProps() m.EmailData {
	p := properties.MustLoadFile(rundir()+"/config.properties", properties.UTF8)
	// fmt.Println(p)
	data := m.EmailData{}

	data.SmtpUser = p.GetString("smtpUser", "")
	data.SmtpPwd = p.GetString("smtpPwd", "")
	data.SendFrom = p.GetString("sendFrom", "")
	data.TemplateDir = p.GetString("templateDir", "")
	data.TemplateName = p.GetString("templateName", "email-template.html")
	data.ExcelFile = p.GetString("excelFile", "")
	data.DryRun = p.GetBool("dryrun", false)
	return data
}

func rundir() string {
	dir, _ := os.Getwd()
	return dir
}

// FindAttachments ..
func FindAttachments(templateDir string) []os.FileInfo {
	files, err := ioutil.ReadDir(templateDir + m.AttachmentSubdir)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

// CheckIfAlreadyProcessed ..
func CheckIfAlreadyProcessed(data m.EmailData) bool {
	chkfile := histDir(data) + "/" + data.ExcelFile
	if _, err := os.Stat(chkfile); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

// MakeHistory dir & copy all file that were used
func MakeHistory(data m.EmailData) {
	fmt.Println("Bestanden die gebruikt worden naar " + histDir(data) + " gekopieerd")
	os.Mkdir(histDir(data), 0700)

	fromExcel := data.TemplateDir + "/" + data.ExcelFile
	toExcel := histDir(data) + "/" + data.ExcelFile
	Copy(fromExcel, toExcel)

	fromTemplate := data.TemplateDir + "/" + data.TemplateName
	toTemplate := histDir(data) + "/" + data.TemplateName
	Copy(fromTemplate, toTemplate)

	//fromAtt := attDir(data) + "/" + data.Attachments
	//toAtt := histDir(data) + "/" + data.Attachments
	//Copy(fromAtt, toAtt)
}

func histDir(data m.EmailData) string {
	return data.TemplateDir + m.HistorySubdir + "/" + DateToStr(time.Now())
}

