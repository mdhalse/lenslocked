package views

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/mdhalse/lenslocked/context"
	"github.com/mdhalse/lenslocked/models"
)

type public interface {
	Public() string
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	tmpl := template.New(patterns[0])
	tmpl = tmpl.Funcs(
		template.FuncMap{
			"csrfField": func() (template.HTML, error) {
				return "", fmt.Errorf("csrfField not implemented")
			},
			"currentUser": func() (template.HTML, error) {
				return "", fmt.Errorf("currentUser not implemented")
			},
			"errors": func() []string {
				return nil
			},
		},
	)

	tmpl, err := tmpl.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing embed/fs: %w", err)
	}
	return Template{html: tmpl}, nil
}

type Template struct {
	html *template.Template
}

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}, errs ...error) {
	tmpl, err := t.html.Clone()
	if err != nil {
		log.Printf("cloning template: %v", err)
		http.Error(w, "There was an error rendering the page.", http.StatusInternalServerError)
		return
	}
	errMsgs := errMessages(errs...)
	tmpl = tmpl.Funcs(
		template.FuncMap{
			"csrfField": func() (template.HTML, error) {
				return csrf.TemplateField(r), nil
			},
			"currentUser": func() *models.User {
				return context.User(r.Context())
			},
			"errors": func() []string {
				return errMsgs
			},
		},
	)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		log.Printf("executing template %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}

func errMessages(errs ...error) []string {
	var messages []string
	for _, err := range errs {
		var pubErr public
		if errors.As(err, &pubErr) {
			messages = append(messages, pubErr.Public())
		} else {
			messages = append(messages, "Something went wrong.")
		}
	}
	return messages
}
