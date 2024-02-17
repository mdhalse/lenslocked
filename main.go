package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmplPath := filepath.Join("templates", "home.gohtml")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	
	err = tmpl.Execute(w, nil) 
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:matthew.halse@hey.com\">matthew.halse@hey.com</a>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
		<h1>FAQs</h1>
		<p><strong>Q:</strong> Is there a free version?<p>
		<p><strong>A:</strong> Yes!  We offer a free trial for 30 days on any paid plans.<p>
		<p><strong>Q:</strong> What are your support hours?<p>
		<p><strong>A:</strong> We have support staff answering emails 24/7, though response times may be a bit slower on the weekend.<p>
		<p><strong>Q:</strong> How do I contact support?<p>
		<p><strong>A:</strong> Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a><p>
	`)
}

func greeterHandler(w http.ResponseWriter, r *http.Request) {
	nameParam := chi.URLParam(r, "name")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(fmt.Sprintf("<h1>Hello!<h2><p>Greetings to %s", nameParam)))
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/greeter/{name}", greeterHandler)
	r.NotFound(http.NotFound)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
