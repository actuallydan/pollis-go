package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	initClerk()

	// Initialize Turso connection
	initTurso()

	// Initialize AWS session and S3 client
	initAWS()

	// Set up the route for presign
	// http.HandleFunc("/presign", Middleware(handlePresign, CORSMiddleware, AuthMiddleware))
	// http.HandleFunc("/analyze", Middleware(handleAnalyze, CORSMiddleware, AuthMiddleware))
	// http.HandleFunc("/search", Middleware(handleSearch, CORSMiddleware, AuthMiddleware))

	// http.HandleFunc("/webhooks/sign-up", handleSignUp)

	// http.HandleFunc("/ping", Middleware(handlePing, CORSMiddleware))
	// http.HandleFunc("/organizations/create", Middleware(handleCreateOrganization, CORSMiddleware, AuthMiddleware))
	// http.HandleFunc("/organizations", Middleware(handleGetUserOrganizations, CORSMiddleware, AuthMiddleware))
	// http.HandleFunc("/organizations/{orgID}", Middleware(handleGetOrganizationBuckets, CORSMiddleware, AuthMiddleware))

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
