package server

import (
	"net/http"
)

func handleHome() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Game Recommender!"))
	})
}
