package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var err = godotenv.Load()

func setupHandlers() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/api/v1/public/checkLicense", getLicenseHandler)
	http.HandleFunc("/api/v1/private/createLicense", createLicenseHandler)
	http.HandleFunc("/api/v1/private/getAll", getLicensesHandler)
}

func main() {
	if err != nil {
		fmt.Println("Error loading.env file")
		os.Exit(1)
	}
	setupHandlers()
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
