package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Println("Key:", k, "Value:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello", r.Form.Encode())
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {

	}
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
