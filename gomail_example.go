package main

import (
	gomail "gopkg.in/gomail.v2"
)

func TestGomail() {
	m := gomail.NewMessage()
	m.SetHeader("From", "if-20043@students.ithb.ac.id")
	m.SetHeader("To", "if-20043@students.ithb.ac.id")
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Test Gomail")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 465, "if-20043@students.ithb.ac.id", "EducationEmail07")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
