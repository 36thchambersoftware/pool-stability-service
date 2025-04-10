package main

import (
	"log"
	"net/http"
	"pool-stability-service/pkg/api"
)



func init() {

}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome."))
		return
	})

	http.HandleFunc("/relays", func(w http.ResponseWriter, r *http.Request) {
		api.GetPoolRelays(w , r )
		return
	})

	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}