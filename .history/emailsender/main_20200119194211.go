package main

import (
	"fmt"

	utl "github.com/jrb/ivn-emailsender/emailsender/util"
)

func main() {
	template := utl.ReadEmailTemplate()
	fmt.Println(template)

	fmt.Println("done")
}
