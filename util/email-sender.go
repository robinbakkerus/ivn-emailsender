package util

import (
	"fmt"
	"gopkg.in/gomail.v2"
)


// SendEmail send email
func SendEmail(
	smtpUser string,
	smtpPwd string,
	sendTo string,
	subject string,
	message string,
	attFilename string) {

	fmt.Println("send to = " + sendTo)

	m := gomail.NewMessage()
	m.SetHeader("From", smtpUser)
	m.SetHeader("To", sendTo, sendTo)
	m.SetAddressHeader("Cc", sendTo, "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)
	m.Attach(attFilename)
	
	d := gomail.NewDialer("smtp.gmail.com", 587, smtpUser, smtpPwd)
	
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}