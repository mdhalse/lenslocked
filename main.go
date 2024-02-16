package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
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

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)
}
