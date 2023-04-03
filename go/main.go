package main

import (
	"net/http"
)

func main()  {

	fs := http.FileServer(http.Dir("../node/dist/"))

	http.Handle("../node/dist/", http.StripPrefix("../node/dist/", fs))

	http.ListenAndServe(":8080", nil)
}
