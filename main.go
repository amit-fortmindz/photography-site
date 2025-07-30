package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	// ✅ Serve static files
	fs := http.FileServer(http.Dir("public/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	imgFs := http.FileServer(http.Dir("public/images"))
	http.Handle("/images/", http.StripPrefix("/images/", imgFs))

	jsFs := http.FileServer(http.Dir("public/js"))
	http.Handle("/js/", http.StripPrefix("/js/", jsFs))

	// ✅ Page routes (AFTER static handlers)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index.html", nil)
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "about.html", nil)
	})

	http.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "services.html", nil)
	})

	http.HandleFunc("/gallery", func(w http.ResponseWriter, r *http.Request) {
		photos := []string{
			"./images/photo1.webp",
			"./images/photo2.webp",
		}
		renderTemplate(w, "gallery.html", photos)
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()
			name := r.FormValue("name")
			email := r.FormValue("email")
			message := r.FormValue("message")

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
	t := template.Must(template.ParseFiles(
		"public/"+tmpl,
		"public/header.html",
		"public/footer.html"))
	t.Execute(w, data)
}
