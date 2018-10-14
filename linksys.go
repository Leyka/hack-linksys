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

func (l *Linksys) GenURL(category string, action string) string {
	return fmt.Sprintf("http://linksys.com/jnap/%s/%s", category, action)
}

func (l *Linksys) prepareRequest(category string, action string) *gorequest.SuperAgent {
	request := gorequest.New()
	return request.Post(l.jnapUrl).
		Set("X-JNAP-ACTION", l.GenURL(category, action)).
		Set("X-JNAP-Authorization", l.headerAuth)
}

func (l *Linksys) MakeRequest(category string, action string) string {
	_, body, err := l.prepareRequest(category, action).End()
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func (l *Linksys) MakeRequestWithBody(category string, action string, body string) string {
	_, body, err := l.prepareRequest(category, action).
		Send(body).
		End()

	if err != nil {
		log.Fatal(err)
	}

	return body
}

func (l *Linksys) MakeRequestTransaction(category string, action string, request string) string {
	coreAction := l.GenURL(category, action)
	body := fmt.Sprintf(`[{"action": "%s", "request": %s}]`, coreAction, request)
	return l.MakeRequestWithBody("core", "Transaction", body)
}
