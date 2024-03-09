package handlers

import (
	"html/template"
	"net/http"
	"aws/models"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func Home(w http.ResponseWriter, r *http.Request) {
	// Check if user is authenticated
	if isAuthenticated(r) {
		user := getCurrentUser(r)
		templates.ExecuteTemplate(w, "home.html", struct {
			Username string
		}{
			Username: user,
		})
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	// Check if user is authenticated
	if isAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Validate user credentials (in a real-world application, you would hash and compare passwords)
		if models.ValidateUser(username, password) {
			setAuthenticated(w, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	templates.ExecuteTemplate(w, "login.html", nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Clear the authentication cookie
	clearAuthenticatedCookie(w)

	// Redirect to the login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request) {
	// Check if user is authenticated
	if isAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		// Handle user registration logic (in a real-world application, you would store user data in the database)
		username := r.FormValue("username")
		password := r.FormValue("password")
		// Save user to the database (you need to implement this function in models/user.go)
		models.CreateUser(username, password)

		// Set user as authenticated
		setAuthenticated(w, username)

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	templates.ExecuteTemplate(w, "register.html", nil)
}
