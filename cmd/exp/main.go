package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
	Age  int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Matt Halse",
		Bio:  `<script>alert("Haha! You've been pwned!")</script>`,
		Age:  35,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
