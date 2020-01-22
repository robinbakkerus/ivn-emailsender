package main

import (
	"fmt"
	utl "github.com/jrb/ivn-emailsender/util"
)

func main() {
	template := utl.ReadEmailTemplate()
	fmt.Println(template)

	// cells := utl.ReadExcelFile()
	// fmt.Println(cells)

	utl.ReadProps()

	fmt.Println("done")
}
