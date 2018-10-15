// Linksys Guard
// Continuous check of incoming connections
// IF incoming connection found, send mail for warning to the subscriber
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"time"
)

var (
	from        string
	password    string
	to          string
	alreadySent bool = false
)

const sleepTime = 30 * time.Minute

// Mainloop to scan for inc. connections
func Scan() {
	fmt.Println("[*] Launching Scan for incoming connections...")
	var incEntries []Entry
	for {
		incEntries = *linksys.GetIncomingEntries()

		if len(incEntries) == 0 {
			// Positive
			if !alreadySent {
				bytes, _ := json.Marshal(incEntries)
				sendWarningMail(string(bytes))
				alreadySent = true
				fmt.Println("[!] Found incoming connections. Warning Mail sent.")
			}
		} else {
			// Reset alreadSent variable if incoming connections is back to normal
			if alreadySent {
				alreadySent = false
			}
		}
		// Wait until next scan
		time.Sleep(sleepTime)
	}
}

func sendWarningMail(body string) {
	from = os.Getenv("EMAIL")
	password = os.Getenv("PASSWORD")
	to = os.Getenv("TO_EMAIL")

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: [LINKSYS] Incoming connections found\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, password, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("[X] SMTP mail error: %s", err)
	}
}
