package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var linksys *Linksys

func incomingConnections(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(linksys.GetIncomingEntries())
	w.Write(json)
}

func wirelessInfo(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(linksys.GetCurrentChannel())
	w.Write(json)
}

func autoChannel(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(linksys.AutoSwitchChannel())
	w.Write(json)
}

func main() {
	host, user, password := readCredentials()
	linksys = NewLinksys(host, user, password)

	// Web server
	fmt.Println("Starting server on port 8080...")

	http.HandleFunc("/incoming", incomingConnections)
	http.HandleFunc("/wifi", wirelessInfo)
	http.HandleFunc("/autochan", autoChannel)

	if err := http.ListenAndServe(":8080", nil); err != nil {
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
