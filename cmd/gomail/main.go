package main

import (
	"fmt"

	"github.com/mdhalse/lenslocked/models"
)

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 2525
	username = "****"
	password = "****"
)

func main() {
	email := models.Email{
		From:      "test@lenslocked.com",
		To:        "matt@example.com",
		Subject:   "This is a test email",
		Plaintext: "This is the body of the email",
		HTML:      `<h1>Hello there!</h1><p>This is the email</p>`,
	}
	es := models.NewEmailService(models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})
	err := es.Send(email)
	if err != nil {
		panic(err)
	}
	fmt.Print("Message sent!")
}
