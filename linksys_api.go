package main

import (
	"encoding/json"
	"math/rand"
	"time"
)

const defaultBody = `{"firstEntryIndex": 1,"entryCount": 255}`

var (
	desiredChannels = []int{1, 6, 11}
)

func (l *Linksys) GetIncomingEntries() *[]Entry {
	var incEntries IncomingEntries
	res := l.MakeRequestWithBody("routerlog", "GetIncomingLogEntries", defaultBody)
	json.Unmarshal([]byte(res), &incEntries)
	return &incEntries.Entries
}

// RadioIndex : 0 => 2.4Ghz, 1 => 5 Ghz
func (l *Linksys) GetRadioSettings(radioIndex int) *Radio {
	var radioInfo RadioInfo
	res := l.MakeRequestTransaction("wirelessap", "GetRadioInfo", "{}")
	json.Unmarshal([]byte(res), &radioInfo)
	return &radioInfo.Responses[0].Output.Radios[radioIndex]
}

func (l *Linksys) GetCurrentChannel() int {
	radio := l.GetRadioSettings(0) // 2.4GHz
	return radio.Settings.Channel
}

// Change WLAN 2.4GHz channel
func (l *Linksys) AutoSwitchChannel() int {
	// Change Channel
	newChannel := pickUnusedChannel(l.GetCurrentChannel())
	currentRadio := l.GetRadioSettings(0)
	currentRadio.Settings.Channel = newChannel

	// Send request
	output := new(RadioOutput)
	output.Radios = []Radio{*currentRadio}
	bytes, _ := json.Marshal(output)

	l.MakeRequestTransaction("wirelessap", "SetRadioSettings", string(bytes))

	return newChannel
}

// Choose a random channel inside the desired channels
func pickUnusedChannel(currentChannel int) int {
	rand.Seed(time.Now().Unix())
	channelsRemain := remove(desiredChannels, currentChannel)
	randIndex := rand.Int() % len(channelsRemain)
	return desiredChannels[randIndex]
}

// Remove element int from array
func remove(arr []int, val int) []int {
	for i, v := range arr {
		if v == val {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}
