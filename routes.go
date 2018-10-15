package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type IndexPageData struct {
	CountIncoming int
	Channel       int
}

var tmpl *template.Template

func InitTemplate() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func SetupRoutes() {
	InitTemplate()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", Index)
	http.HandleFunc("/incoming", IncomingConnections)
	http.HandleFunc("/wifi", WirelessInfo)
	http.HandleFunc("/autochan", AutoChannel)
}

func Index(w http.ResponseWriter, r *http.Request) {
	data := &IndexPageData{
		CountIncoming: len(*linksys.GetIncomingEntries()),
		Channel:       linksys.GetCurrentChannel(),
	}
	tmpl.ExecuteTemplate(w, "index.html", data)
}

func IncomingConnections(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(linksys.GetIncomingEntries())
	w.Write(json)
}

func WirelessInfo(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(linksys.GetRadioSettings(0))
	w.Write(json)
}

func AutoChannel(w http.ResponseWriter, r *http.Request) {
	linksys.AutoSwitchChannel()
	Index(w, r)
}
