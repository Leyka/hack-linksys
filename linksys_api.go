package main

import "fmt"

func (l *Linksys) GetIncomingEntries() {
	respStr := l.MakeRequest("routerlog", "GetIncomingLogEntries", true)
	fmt.Println(respStr)
}

func (l *Linksys) GetCurrentWlanChannel() {

}

// Change WLAN 2.4GHz channel between 1,6 and 11
func (l *Linksys) AutoSwitchChannel() {

}
