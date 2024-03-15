package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
	"github.com/mdhalse/lenslocked/controllers"
	"github.com/mdhalse/lenslocked/models"
	"github.com/mdhalse/lenslocked/templates"
	"github.com/mdhalse/lenslocked/views"
)

func main() {
	r := chi.NewRouter()

	// TODO: Don't use secure==false in production envs
	csrfKey := "gFvi44Fy5xnbLNeEztQbfAvCyEIaux"
	r.Use(csrf.Protect([]byte(csrfKey), csrf.Secure(false)))
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

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userService := models.UserService{
		DB: db,
	}
	sessionService := models.SessionService{
		DB: db,
	}

	usersController := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}
	usersController.Templates.New = views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "signup.gohtml"))
	usersController.Templates.SignIn = views.Must(views.ParseFS(templates.FS, "tailwind.gohtml", "signin.gohtml"))
	r.Get("/signup", usersController.New)
	r.Post("/users", usersController.Create)
	r.Get("/signin", usersController.SignIn)
	r.Post("/signin", usersController.ProcessSignIn)
	r.Post("/signout", usersController.ProcessSignOut)
	r.Get("/users/me", usersController.CurrentUser)

	r.NotFound(http.NotFound)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
