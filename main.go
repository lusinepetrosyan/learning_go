package main

import (
	"fmt"
	"net/http"
)

func myHandlerFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "welcome to my site")
}

func main(){
	http.HandleFunc("/", myHandlerFunc)
	http.ListenAndServe(":8000", nil)
}