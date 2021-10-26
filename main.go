package main

import (
	"log"
	"net/http"
)

func main()  {
	mux := http.NewServeMux()
	mux.Handle("/", handler())

	log.Println("Listening on port :3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}

func handler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing Handler")
		w.Write([]byte("OK"))
	})
}