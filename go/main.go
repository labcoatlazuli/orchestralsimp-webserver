package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main()  {
	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./dist"))))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./dist/index.html")
	})

	router.HandleFunc("/sheet-music.xml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./sheet-music.xml")
	})

	http.ListenAndServe(":8080", router)
}