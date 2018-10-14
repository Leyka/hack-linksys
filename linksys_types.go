package main

// Incoming Entries
type IncomingEntries struct {
	Entries []Entry `json:"output.entries"`
}

type Entry struct {
	DestinationPort uint   `json:"destinationPort"`
	Source          string `json:"source"`
}

// Wireless
type RadioInfo struct {
	Result    string          `json:"result"`
	Responses []RadioResponse `json:"responses"`
}

type RadioResponse struct {
	Result string      `json:"result"`
	Output RadioOutput `json:"output"`
}

type RadioOutput struct {
	Radios []Radio `json:"radios"`
}

type Radio struct {
	RadioId  string        `json:"radioID"`
	Settings RadioSettings `json:"settings"`
}

type RadioRequest struct {
	Radios []Radio `json:"radios"`
}

type RadioSettings struct {
	SSID                string              `json:"ssid"`
	BroadcastSSID       bool                `json:"broadcastSSID"`
	Channel             int                 `json:"channel"`
	ChannelWidth        string              `json:"channelWidth"`
	Enabled             bool                `json:"isEnabled"`
	Mode                string              `json:"mode"`
	Security            string              `json:"security"`
	WPAPersonalSettings WPAPersonalSettings `json:"wpaPersonalSettings"`
}

type WPAPersonalSettings struct {
	Passphrase string `json:"passphrase"`
}
