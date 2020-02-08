package model

const AttachmentSubdir = "/attachment"
const HistorySubdir = "/verstuurd"

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
	Attachment   string
	MailListIdx  int
	MailList     string
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
