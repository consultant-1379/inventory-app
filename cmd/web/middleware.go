package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// JUST TO CHECK IF SITE RESPONDS
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

// NoSurf csrf protection
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
		Path:     "/",
	})

	return csrfHandler
}

// SessionLoad Starts session
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
