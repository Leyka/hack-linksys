package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var host string
var user string
var password string

func main() {
	readCredentials()
	l := NewLinksys(host, user, password)
	l.GetIncomingEntries()
}

func readCredentials() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	host = os.Getenv("LINKSYS_HOST")
	user = os.Getenv("LINKSYS_USER")
	password = os.Getenv("LINKSYS_PW")
}
