package util

import (
	"fmt"
	"net/smtp"
)

// smtpServer data to smtp server.
type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server.
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

// send email
func SendEmail(sendTo string, message []byte) {
	// Sender data.
	from := "robin.bakkerus@gmail.com"
	password := "Mets2233"

	// Receiver email address.
	to := []string{sendTo}

	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}

	// Authentication.
	auth := smtp.PlainAuth("", "robin.bakkerus@gmail.com", password, smtpServer.host)

	// Sending email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email Sent!")
}
