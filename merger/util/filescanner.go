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

	fmt.Println("Op zoek naar dumpfiles in " + data.DownloadDir)

	files, err := ioutil.ReadDir(data.DownloadDir)
	if err != nil {
		log.Fatal(err)
	}

	result := make([]string, 0)

	for _, f := range files {

		if strings.HasSuffix(f.Name(), ".csv") &&
			(!strings.HasSuffix(f.Name(), data.ExcelFile)) {
			result = append(result, f.Name())
		}
	}

	return result
}
