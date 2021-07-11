package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", ViewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadpage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

//used for rendering HTTP
