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
	if isAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if models.ValidateUser(username, password) {
			setAuthenticated(w, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	templates.ExecuteTemplate(w, "login.html", nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	clearAuthenticatedCookie(w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if isAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		models.CreateUser(username, password)

		setAuthenticated(w, username)

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	templates.ExecuteTemplate(w, "register.html", nil)
}
