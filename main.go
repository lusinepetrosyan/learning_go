package main

import (
	"fmt"
	"net/http"
)

func myHandlerFunc(w http.ResponseWriter, r *http.Request){

	if(r.URL.Path == "/"){
		fmt.Fprint(w, "welcome to my site")
	}else if (r.URL.Path == "/contact"){
		fmt.Fprint(w, "this is a contact page")
	}else{
		fmt.Fprint(w, "incorrect page")
		w.WriteHeader(http.StatusNotFound)
	}
}

func main(){
	http.HandleFunc("/", myHandlerFunc)
	http.ListenAndServe(":8000", nil)
}