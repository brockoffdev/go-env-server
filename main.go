package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func giveEnvVars(w http.ResponseWriter, r *http.Request) {
	println("path", r.URL.Path)
	println("scheme", r.URL.Scheme)

	fmt.Fprintf(w, "LIST OF ENVIRONMENT VARIABLES: \n")
	fmt.Fprintf(w, "==================================================== \n \n")
	for _, e := range os.Environ() {
		fmt.Fprintf(w, "%s \n", e)
	}
}

func main() {
	http.HandleFunc("/", giveEnvVars)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
