package main

import (
	"fmt"
	utl "/util"
)

func main() {
	template := utl.ReadEmailTemplate()
	fmt.Println(template)

	fmt.Println("done")
}
