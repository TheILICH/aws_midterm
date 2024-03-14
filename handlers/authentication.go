package handlers

import (
	"fmt"
	"time"
	"net/http"	
)

const authenticatedKey = "authenticated"

func clearAuthenticatedCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    authenticatedKey,
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	})
}

func setAuthenticated(w http.ResponseWriter, username string) {
	fmt.Printf("\nSETAUTHENTICATED WITH USERNAME = %v\n\n", username)
	http.SetCookie(w, &http.Cookie{
		Name:  authenticatedKey,
		Value: username,
		Path:  "/",
	})
}

func isAuthenticated(r *http.Request) bool {
	// fmt.Printf("\nISAUTHENTICATED WITH USERNAME = %v\n\n", username)
	cookie, err := r.Cookie(authenticatedKey)
	return err == nil && cookie != nil
}

func getCurrentUser(r *http.Request) string {
	if cookie, err := r.Cookie(authenticatedKey); err == nil && cookie != nil {
		fmt.Printf("\nISAUTHENTICATED WITH USERNAME = %v\n\n", cookie.Value)
		return cookie.Value
	}
	return ""
}
