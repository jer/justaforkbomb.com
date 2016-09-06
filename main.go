package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type forkbombserver string

func (f forkbombserver) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, f)
}

func main() {
	var f forkbombserver
	f = ":(){ :|: & };:"
	listen := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(listen, f))
}
