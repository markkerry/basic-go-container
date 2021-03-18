package main

import (
  "fmt"
  "net/http"
  "os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
  hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
  fmt.Fprintf(w, "Go App running on host: %s\n", hostname)
}

func setupRoutes() {
  http.HandleFunc("/", homePage)
}

func main() {
  fmt.Printf("Go Web App Started on Port 3000\n")
  setupRoutes()
  http.ListenAndServe(":3000", nil)
}