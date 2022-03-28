package main

import (
	"log"

	gomail "gopkg.in/gomail.v2"
)

func TestGomail() {
	m := gomail.NewMessage()
	m.SetHeader("From", "if-20043@students.ithb.ac.id")
	m.SetHeader("To", "if-20043@students.ithb.ac.id")
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Test Gomail")
	m.SetBody("text/html", "Hello, this email was send from <b>Golang</b>!")
	// m.Attach("/home/Alex/lolcat.jpg")

	emailPassword := LoadEnv("EMAIL_PASS")
	log.Println(emailPassword)
	d := gomail.NewDialer("smtp.gmail.com", 465, "if-20043@students.ithb.ac.id", emailPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
