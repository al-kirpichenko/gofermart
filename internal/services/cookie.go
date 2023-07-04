package services

import "net/http"

func setCookie(w http.ResponseWriter, token string) *http.Cookie {

	newCookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		MaxAge:   10800,
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, newCookie)
	return newCookie

}
