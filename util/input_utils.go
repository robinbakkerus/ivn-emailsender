package util

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// AskBool ..
func AskBool(askmsg string, defval bool) bool {
	reader := bufio.NewReader(os.Stdin)

	defTxt := ""
	if defval {
		defTxt = "J"
	} else {
		defTxt = "N"
	}

	msg := askmsg + " (" + defTxt + ") > "
	fmt.Println(msg)

	textInput, _ := reader.ReadString(Delim())
	textInput = strings.ToUpper(TrimDelim(textInput))
	if len(textInput) == 0 {
		textInput = defTxt
	}

	return strings.HasPrefix(textInput, "J")
}

// AskString
func AskString(defSubject string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Geef onderwerp (" + defSubject + ") > ")
	textInput, _ := reader.ReadString(Delim())
	r := TrimDelim(textInput)
	if len(r) == 0 {
		return defSubject
	}
	return r
}

func Delim() byte {
	if runtime.GOOS == "windows" {
		return '\r'
	} else {
		return '\n'
	}
}

func TrimDelim(s string) string {
	if runtime.GOOS == "windows" {
		return strings.TrimRight(s, "\r")
	} else {
		return strings.TrimRight(s, "\n")
	}
}
