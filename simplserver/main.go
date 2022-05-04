package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	s := &http.Server{
		Addr:           ":8181",
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	log.Println("Server started...")
	log.Fatal(s.ListenAndServe())

}
