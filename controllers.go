package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Weesli")
	}
}

func getLicenseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	key := r.URL.Query().Get("key")
	license := getLicense(key)
	if license == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"error": "License not found"}`)
		return
	}
	json.NewEncoder(w).Encode(license)
}

func createLicenseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("admin-secret") == "" || r.Header.Get("admin-secret") != os.Getenv("ADMIN_SECRET") {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"error": "Unauthorized"}`)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var license License
	err := json.NewDecoder(r.Body).Decode(&license)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "Invalid request body"}`)
		return
	}
	createLicense(license)
	json.NewEncoder(w).Encode(license)
}

func getLicensesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("admin-secret") == "" || r.Header.Get("admin-secret") != os.Getenv("ADMIN_SECRET") {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"error": "Unauthorized"}`)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	licenses := getLicenses()
	json.NewEncoder(w).Encode(licenses)
}

func deleteLicenseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("admin-secret") == "" || r.Header.Get("admin-secret") != os.Getenv("ADMIN_SECRET") {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, `{"error": "Unauthorized"}`)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	key := r.URL.Query().Get("key")
	deleteLicense(key)
	fmt.Fprintf(w, `{"message": "License deleted successfully"}`)
}
