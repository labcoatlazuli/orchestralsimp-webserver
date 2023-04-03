package main

import (
	"net/http"
	"log"
)

func main()  {

	fs := http.FileServer(http.Dir("../node/dist/"))

	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
