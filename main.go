package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mdhalse/lenslocked/controllers"
	"github.com/mdhalse/lenslocked/templates"
	"github.com/mdhalse/lenslocked/views"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.CleanPath)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	tpl := views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "faq.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	usersController := controllers.Users{}
	usersController.Templates.New = views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "signup.gohtml"))
	r.Get("/signup", usersController.New)

	r.NotFound(http.NotFound)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
