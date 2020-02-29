package model

import "os"

const AttachmentSubdir = "/attachment"
const HistorySubdir = "/verstuurd"

// colums on mailing.xlsx
const NAME = 1
const EMAIL = 2
const AANHEF = 3

// EmailData ...
type EmailData struct {
	SmtpUser     string
	SmtpPwd      string
	SendFrom     string
	TemplateDir  string
	TemplateName string
	ExcelFile    string
	TemplateBody string
	Subject      string
	Attachments  []os.FileInfo
	MailListIdx  int
	MailList     string
	ImageName    string
	Aanhef       string
	DryRun       bool
}

// Excel ..
type Excel struct {
	Name string
	Size int
}

// BySize implements sort.Interface based on the Size field.
type BySize []Excel

func (a BySize) Len() int           { return len(a) }
func (a BySize) Less(i, j int) bool { return a[i].Size < a[j].Size }
func (a BySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
