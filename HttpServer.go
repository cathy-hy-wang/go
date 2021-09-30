package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("incoming request from IP: %s\n", getIP(r))

	//read request headers then set in reponse headers
	for k, v := range r.Header {
		for _, v1 := range v {
			fmt.Printf("Adding Response Header --> key: %s, value: %s\n", k, v1)
			w.Header().Add(k, v1)
		}
	}
	w.Header().Add("version", os.Getenv("VERSION"))

	w.WriteHeader(200)
	fmt.Fprintf(w, "hello world")

	fmt.Printf("#########################################\n\n")
}

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func main() {
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Listen and Serve", err)
	}
}
