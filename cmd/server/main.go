package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My Awesome Go App")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
}

func main() {
	cmd := exec.Command("/bin/sh", "refresh.sh")
	fmt.Println("Go Web App Started on Port 3000")
	setupRoutes()

	go cmd.Run()
	go http.ListenAndServe(":3000", nil)
}
