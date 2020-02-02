package util

import (
	"fmt"
	"io/ioutil"
	m "ivnmailer/model"
	"log"
	"os"
	"sort"
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
	return data
}

func rundir() string {
	dir, _ := os.Getwd()
	return dir
}

// FindExcelFiles ..
func FindExcelFiles(templateDir string) []m.Excel {
	files, err := ioutil.ReadDir(templateDir)
	if err != nil {
		log.Fatal(err)
	}

	var excelFiles []m.Excel

	for _, f := range files {
		if strings.HasSuffix(strings.ToUpper(f.Name()), ".XLSX") && !strings.HasPrefix(f.Name(), "~") {
			rows := ReadExcelFile(templateDir + "/" + f.Name())
			excelFiles = append(excelFiles, m.Excel{f.Name(), len(rows)})
		}
	}

	sort.Sort(m.BySize(excelFiles))
	return excelFiles
}

// FindAttachment ..
func FindAttachment(templateDir string) string {
	files, err := ioutil.ReadDir(templateDir + m.AttachmentSubdir)
	if err != nil {
		log.Fatal(err)
	}

	if len(files) > 1 {
		fmt.Println("Meer dan 1 attachment gevonden in " + templateDir + m.AttachmentSubdir)
		fmt.Println("Op dit moment mag er 1 of 0 attachment worden gestuurd")
		os.Exit(0)
	} else if len(files) == 1 {
		return files[0].Name()
	}
	return ""
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

	fromAtt := attDir(data) + "/" + data.Attachment
	toAtt := histDir(data) + "/" + data.Attachment
	Copy(fromAtt, toAtt)
}

func histDir(data m.EmailData) string {
	return data.TemplateDir + m.HistorySubdir + "/" + DateToStr(time.Now())
}

func attDir(data m.EmailData) string {
	return data.TemplateDir + m.AttachmentSubdir
}
