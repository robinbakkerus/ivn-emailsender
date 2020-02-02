package util

import (
	"fmt"

	model "ivnmailer/model"

	"gopkg.in/gomail.v2"
)

// SendEmail send email
func SendEmail(data model.EmailData, sendTo string, sendToName string, emailBody string) {

	fmt.Println("send to = " + sendTo + " " + sendToName)

	m := gomail.NewMessage()
	m.SetHeader("From", data.SmtpUser)
	m.SetAddressHeader("To", sendTo, sendToName)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", emailBody)
	if len(data.Attachment) > 0 {
		m.Attach(data.TemplateDir + model.AttachmentSubdir + "/" + data.Attachment)
	}

	fmt.Println("sending email to " + sendTo + " " + sendToName)
	d := gomail.NewDialer("smtp.gmail.com", 587, data.SmtpUser, data.SmtpPwd)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
