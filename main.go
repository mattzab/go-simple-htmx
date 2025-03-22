package main

import (
	"html/template"
	"net/http"
	"strings"
	"sync"
)

var tmpl *template.Template
var counter int = 0
var counterMutex sync.Mutex

// Form data structure
type FormData struct {
	Name  string
	Email string
}

func main() {
	// Parse all templates in the templates directory
	tmpl = template.Must(template.ParseGlob("templates/*.html"))

	// Serve static files from the static directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Main page handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

	// Dynamic HTML handler
	http.HandleFunc("/html/", handleDynamicHtml)

	// Start the server
	println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleDynamicHtml(w http.ResponseWriter, r *http.Request) {
	// Extract the template name from the URL path
	path := strings.TrimPrefix(r.URL.Path, "/html/")
	if path == "" {
		http.Error(w, "Template not specified", http.StatusBadRequest)
		return
	}

	// Handle special routes
	switch path {
	case "increment":
		handleIncrement(w, r)
		return
	case "reset":
		handleReset(w, r)
		return
	case "submitform":
		handleFormSubmission(w, r)
		return
	case "counter":
		// Initialize the counter view
		err := tmpl.ExecuteTemplate(w, "counter", counter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Render standard templates
	err := tmpl.ExecuteTemplate(w, path, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleIncrement(w http.ResponseWriter, r *http.Request) {
	counterMutex.Lock()
	counter++
	currentCount := counter
	counterMutex.Unlock()

	err := tmpl.ExecuteTemplate(w, "counter", currentCount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleReset(w http.ResponseWriter, r *http.Request) {
	counterMutex.Lock()
	counter = 0
	counterMutex.Unlock()

	err := tmpl.ExecuteTemplate(w, "counter", 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleFormSubmission(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get form values
	formData := FormData{
		Name:  r.FormValue("name"),
		Email: r.FormValue("email"),
	}

	// Render the form submitted template with the form data
	err = tmpl.ExecuteTemplate(w, "formsubmitted", formData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
