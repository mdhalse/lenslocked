package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/mdhalse/lenslocked/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	host := os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(err)
	}
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
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
	err = es.Send(email)
	if err != nil {
		panic(err)
	}
	fmt.Print("Message sent!")
}
