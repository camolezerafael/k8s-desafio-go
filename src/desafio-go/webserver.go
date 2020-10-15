package main

import (
	"fmt"
	"log"
	"net/http"
)

func greeting( text string ) string{
	return fmt.Sprintf("<b>%s</b>", text)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, greeting("Code.education Rocks!"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}