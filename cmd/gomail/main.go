package main

import (
	"fmt"
	"os"

	"github.com/go-mail/mail/v2"
)

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 2525
	username = "****"
	password = "****"
)

func main() {
	from := "test@lenslocked.com"
	to := "matt@example.com"
	subject := "This is a test email"
	plaintext := "This is the body of the email"
	html := `<h1>Hello there!</h1><p>This is the email</p>`
	msg := mail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", plaintext)
	msg.AddAlternative("text/html", html)
	msg.WriteTo(os.Stdout)

	dialer := mail.NewDialer(host, port, username, password)
	err := dialer.DialAndSend(msg)
	if err != nil {
		panic(err)
	}
	fmt.Print("Message sent!")
}
