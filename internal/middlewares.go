package internal

import (
	"log"
	"net/http"

	"github.com/masterhorn/XM-Golang-Exercise-v21.0.0/cfg"
)

func AccessHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if ipRegionDetection(r) != "Cyprys" {
			errorMessage := "Access header is invalid"
			log.Println("You have no access not to make this operation")
			http.Error(w, errorMessage, http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		accessHeader := r.Header.Get("AccessToken")

		if accessHeader != cfg.JWT_Token {
			errorMessage := "Access header is invalid"
			log.Println(errorMessage)
			http.Error(w, errorMessage, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
