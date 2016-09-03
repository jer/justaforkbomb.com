package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, ":(){ :|: & };:")
}

func main() {
	listen := fmt.Sprintf(":%s", os.Getenv("PORT"))
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(listen, nil))
}
