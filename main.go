package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	// Serve static files (CSS, JS, Images)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Home page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index.html", nil)
	})

	// About page
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "about.html", nil)
	})

	// Services page
	http.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "services.html", nil)
	})

	// Gallery page
	http.HandleFunc("/gallery", func(w http.ResponseWriter, r *http.Request) {
		photos := []string{
			"templates/static/images/photo1.webp",
			"templates/static/images/photo2.webp",
		}
		renderTemplate(w, "gallery.html", photos)
	})

	// Contact page (GET/POST)
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()
			name := r.FormValue("name")
			email := r.FormValue("email")
			message := r.FormValue("message")

			// Save submissions to file
			f, _ := os.OpenFile("submissions.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			defer f.Close()
			f.WriteString(fmt.Sprintf("Name: %s, Email: %s, Message: %s\n", name, email, message))

			http.Redirect(w, r, "/contact?success=true", http.StatusSeeOther)
			return
		}
		renderTemplate(w, "contact.html", nil)
	})

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Helper function to render templates
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t := template.Must(template.ParseFiles("templates/"+tmpl, "templates/header.html", "templates/footer.html"))
	t.Execute(w, data)
}
