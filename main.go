package main

import (
	"fmt"
	"net/http"
	// ("html/template")
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome demo</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8080", nil)
}

// func html(){
// 	<h1>editing {{title}}</h1>
// 	<form action="/save/{{.Title}}" method="POST">
// 	<div><textarea name="body" rows="20" cols="80">{{printf "%s" .Body}}</textarea></div>
// 	<div><input type="submit" value="Save"></div>
// 	</form>
// }

// func editHandler(w http.ResponseWriter, r *http.Request) {
//     title := r.URL.Path[len("/edit/"):]
//     p, err := loadPage(title)
//     if err != nil {
//         p = &Page{Title: title}
//     }
//     t, _ := template.ParseFiles("edit.html")
//     t.Execute(w, p)
// }
