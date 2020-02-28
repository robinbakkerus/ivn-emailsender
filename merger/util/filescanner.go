package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	// "os"

	m "jrb/ivn-emailsender/mailer/model"
	//"strings"
	//"time"
)

// FindDumpFiles ...
func FindDumpFiles(data m.EmailData) []string {

	fmt.Println(" op zoek naar dumpfiles in " + data.TemplateDir)

	files, err := ioutil.ReadDir(data.TemplateDir)
	if err != nil {
		log.Fatal(err)
	}

	result := make([]string, 0)

	for _, f := range files {

		if strings.HasSuffix(f.Name(), ".xlsx") &&
			(!strings.HasSuffix(f.Name(), data.ExcelFile)) {
			result = append(result, f.Name())
		}
	}

	return result
}
