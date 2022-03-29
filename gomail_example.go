package main

import (
	"log"

	gomail "gopkg.in/gomail.v2"
)

func TestGomail() {
	m := gomail.NewMessage()

	// Get value from env
	emailSender := LoadEnv("EMAIL_SENDER")
	emailReceiver := LoadEnv("EMAIL_RECEIVER")
	emailPassword := LoadEnv("EMAIL_PASS")

	// Set email content
	m.SetHeader("From", emailSender)
	m.SetHeader("To", emailReceiver)
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Test Gomail")
	m.SetBody("text/html", "Hello, this email was send from <b>Golang</b>!")
	m.Attach("Photo 1.jpeg")

	d := gomail.NewDialer("smtp.gmail.com", 465, emailSender, emailPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	log.Println("Send email success")
}
