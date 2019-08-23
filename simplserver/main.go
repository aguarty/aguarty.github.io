package main

import (
	"io/ioutil"
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

	http.HandleFunc("/", fooHandler)
	log.Println("Server started...")
	log.Fatal(s.ListenAndServe())

}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	body := openFile("../aguarty.github.io/index.html")
	w.Write(body)

}

func openFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}
