package handlers

import (
	"fmt"
	"time"
	"net/http"	
)

const authenticatedKey = "authenticated"

func clearAuthenticatedCookie(w http.ResponseWriter) {
	// Clear the authentication cookie by setting an empty cookie with an expired date
	http.SetCookie(w, &http.Cookie{
		Name:    authenticatedKey,
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	})
}

func setAuthenticated(w http.ResponseWriter, username string) {
	// Set a cookie or session variable to mark the user as authenticated
	fmt.Printf("\nSETAUTHENTICATED WITH USERNAME = %v\n\n", username)
	http.SetCookie(w, &http.Cookie{
		Name:  authenticatedKey,
		Value: username,
		Path:  "/",
	})
}

func isAuthenticated(r *http.Request) bool {
	// Check if the user is marked as authenticated using the cookie or session variable
	// fmt.Printf("\nISAUTHENTICATED WITH USERNAME = %v\n\n", username)
	cookie, err := r.Cookie(authenticatedKey)
	return err == nil && cookie != nil
}

func getCurrentUser(r *http.Request) string {
	// Retrieve the username from the cookie or session variable
	if cookie, err := r.Cookie(authenticatedKey); err == nil && cookie != nil {
		fmt.Printf("\nISAUTHENTICATED WITH USERNAME = %v\n\n", cookie.Value)
		return cookie.Value
	}
	return ""
}
