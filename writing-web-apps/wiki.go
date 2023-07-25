package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

// The function template.Must is a convenience wrapper that panics when passed a non-nil error value,
// and otherwise returns the *Template unaltered. A panic is appropriate here; if the templates can't
// be loaded the only sensible thing to do is exit the program.
var TEMPLATES = template.Must(template.ParseFiles("edit.html", "view.html"))

// The function regexp.MustCompile will parse and compile the regular expression, and return a regexp.Regexp.
// MustCompile is distinct from Compile in that it will panic if the expression compilation fails, while Compile
// returns an error as a second parameter.
var VALID_PATH *regexp.Regexp = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, error := os.ReadFile(filename)

	// Propagate error if file isn't found
	if error != nil {
		return nil, error
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := TEMPLATES.ExecuteTemplate(w, tmpl+".html", p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)

	// If page DNE, redirect to edit a new page
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	} else {
		renderTemplate(w, "view", p)
	}

}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	// Load page if exist. Otherwise, create an empty page struct
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := Page{Title: title, Body: []byte(body)}

	err := p.save()
	if err != nil { // Report any errors with saving
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/view/"+title, http.StatusFound)
	}
}

// Wrapper around the pre-existing handler functions. This turns those into a handlerFunc,
// while adding some extra functionality!
// NOTE: This is the decorator pattern. It's basically a python decorator, but a little
// more transparent! :)
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := VALID_PATH.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
