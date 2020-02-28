package util

import (
	"fmt"

	model "jrb/ivn-emailsender/mailer/model"

	"gopkg.in/gomail.v2"
)

// SendEmail send email
func SendEmail(data model.EmailData, sendTo string, sendToName string, emailBody string) {

	m := gomail.NewMessage()
	m.SetHeader("From", data.SmtpUser)

	m.SetAddressHeader("To", sendTo, sendToName)
	m.SetHeader("Subject", data.Subject)
	m.Embed(data.TemplateDir + "/" + data.ImageName)
	m.SetBody("text/html", emailBody)

	for i := 0; i < len(data.Attachments); i++ {
		m.Attach(data.TemplateDir + model.AttachmentSubdir + "/" + data.Attachments[i].Name())
	}

	if len(sendTo) > 0 {
		if data.DryRun {
			fmt.Println("DRY-RUN: sending email to " + sendTo + " " + sendToName)
		} else {
			fmt.Println("sending email to " + sendTo + " " + sendToName)
			d := gomail.NewDialer("smtp.gmail.com", 587, data.SmtpUser, data.SmtpPwd)

			if err := d.DialAndSend(m); err != nil {
				fmt.Println(" Fout tijdens versturen van email naar " + sendTo + " " + err.Error())
			}
		}
	}
}
