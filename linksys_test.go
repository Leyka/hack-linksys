package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	host, user, password := readCredentials()
	linksys = NewLinksys(host, user, password)
}

func TestSwitchChannel(t *testing.T) {
	before := linksys.GetCurrentChannel()
	// Switch channel auto
	linksys.AutoSwitchChannel()
	assert.NotEqual(t, before, linksys.GetCurrentChannel())
}
