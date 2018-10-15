package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const port = 8888

var linksys *Linksys

func main() {
	// Configure .env
	host, user, password := readCredentials()
	linksys = NewLinksys(host, user, password)

	// Start scanner
	go ScanIncConnections()

	// Start web server
	fmt.Println("[*] Starting server on port", port)
	SetupRoutes()
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}

func readCredentials() (string, string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("LINKSYS_HOST")
	user := os.Getenv("LINKSYS_USER")
	password := os.Getenv("LINKSYS_PW")

	return host, user, password
}
