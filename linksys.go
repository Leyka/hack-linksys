package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/parnurzeal/gorequest"
)

type Linksys struct {
	jnapUrl    string
	headerAuth string
}

func NewLinksys(host string, user string, password string) *Linksys {
	l := new(Linksys)
	l.jnapUrl = fmt.Sprintf("http://%s/JNAP/", host)
	// Encode user and password to base64
	userPassword := fmt.Sprintf("%s:%s", user, password)
	encoded := base64.StdEncoding.EncodeToString([]byte(userPassword))
	l.headerAuth = fmt.Sprintf("Basic %s", encoded)
	return l
}

func (l *Linksys) genURL(category string, action string) string {
	return fmt.Sprintf("http://linksys.com/jnap/%s/%s", category, action)
}

func (l *Linksys) MakeRequest(category string, action string, withBody bool) string {
	var body string = ``
	if withBody {
		body = `{"firstEntryIndex": 1,"entryCount": 255}`
	}

	request := gorequest.New()
	_, body, err := request.
		Post(l.jnapUrl).
		Set("X-JNAP-ACTION", l.genURL(category, action)).
		Set("X-JNAP-Authorization", l.headerAuth).
		Send(body).
		End()

	if err != nil {
		log.Fatal(err)
	}

	return body
}
