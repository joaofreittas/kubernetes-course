package main

import (
	"io/ioutil"
	"net/http"
	"log"
	"fmt"
	"os"
)

func main() {
	http.HandleFunc("/config-map", Hello)
	http.HandleFunc("/", ConfigMap)
	http.ListenAndServe(":80", nil)

}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello, I'm %s. I'm %s.", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile("myfamily/family.txt")

	if err != nil {
		log.Fatalf("Error reading file", err)
	}

	fmt.Fprintf(w, "My Family: %s.", string(data))
}