/*
	Simple License control server system
	You will need to replace the MongoDB connection string, admin-secret and other sensitive data with your own.
	Author @Weesli
*/

package main

import (
	"LicenseChecker/controller"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var err = godotenv.Load()

/*
setupHandlers sets up all the routes for the application.
It includes GET /, POST /api/v1/public/checkLicense, POST /api/v1/private/createLicense,
GET /api/v1/private/getAll, and DELETE /api/v1/private/removeLicense routes.
It also checks for admin-secret header in the requests and returns unauthorized status if it's missing or incorrect.
*/
func setupHandlers() {
	http.HandleFunc("/", controller.HelloHandler)
	http.HandleFunc("/api/v1/public/checkLicense", controller.GetLicenseHandler)
	http.HandleFunc("/api/v1/private/createLicense", controller.CreateLicenseHandler)
	http.HandleFunc("/api/v1/private/getAll", controller.GetLicensesHandler)
	http.HandleFunc("/api/v1/private/removeLicense", controller.DeleteLicenseHandler)
}

/*
main function initializes the application by loading environment variables, setting up handlers,
and starting the server on port 8080. If there's an error loading the.env file, it prints an error message and exits the program.
It also checks for admin-secret header in the requests and returns unauthorized status if it's missing or incorrect.
The server listens for incoming HTTP requests and serves them accordingly.
It returns a 401 status code if the admin-secret header is missing or incorrect.
It prints a message indicating that the server is listening on port 8080.
*/
func main() {
	if err != nil {
		fmt.Println("Error loading.env file")
		os.Exit(1)
	}
	setupHandlers()
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
