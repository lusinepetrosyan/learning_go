package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func home(w http.ResponseWriter, request *http.Request){
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1 color='green'>this is my home page</h1>")
}

func contact(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1 color='brown'>this is my contact page</h1>")
}

func fac(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h2>This is my FAQ page </h2>" +
		"<img src='/image.jpg' width='100' height='100'>")
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/fac", fac)

	http.ListenAndServe(":8000", r)
}