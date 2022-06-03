package server

import (
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/", handleHome())
	http.HandleFunc("/allgames", handleAllGames())

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
