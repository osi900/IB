package main

import (
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := os.Getenv("RESPONSE")
	if len(response) == 0 {
		response = "Hello OpenShift for Developers!"
	}
	
	fmt.Fprintln(w, response)
	fmt.Println("Servicing an impatient beginner's request.")
}