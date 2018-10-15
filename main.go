package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const port = 8888

var linksys *Linksys
var tmpl *template.Template

func main() {
	host, user, password := readCredentials()
	linksys = NewLinksys(host, user, password)

	go ScanIncConnections()

	// Web server
	fmt.Println("[*] Starting server on port", port)
	initTemplate()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/incoming", incomingConnections)
	http.HandleFunc("/wifi", wirelessInfo)
	http.HandleFunc("/autochan", autoChannel)
	http.HandleFunc("/channel", currentChannel)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}

func initTemplate() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
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

// Routes
type Data struct {
	CountIncoming int
	Channel       int
}

func index(w http.ResponseWriter, r *http.Request) {
	data := &Data{
		CountIncoming: len(*linksys.GetIncomingEntries()),
		Channel:       linksys.GetCurrentChannel(),
	}
	tmpl.ExecuteTemplate(w, "index.html", data)
}

func incomingConnections(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(linksys.GetIncomingEntries())
	w.Write(json)
}

func wirelessInfo(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(linksys.GetRadioSettings(0))
	w.Write(json)
}

func autoChannel(w http.ResponseWriter, r *http.Request) {
	linksys.AutoSwitchChannel()
	index(w, r)
}

func currentChannel(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("[2.4GHz] Current channel => %d", linksys.GetCurrentChannel())))
}
