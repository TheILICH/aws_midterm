package main

import (
	"net/http"
	"aws/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/logout", handlers.Logout)
	http.ListenAndServe(":80", nil)
}
