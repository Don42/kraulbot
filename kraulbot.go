// ----------------------------------------------------------------------------
// "THE SCOTCH-WARE LICENSE" (Revision 42):
// <DonMarco42@gmail.com> wrote this file. As long as you retain this notice you
// can do whatever you want with this stuff. If we meet some day, and you think
// this stuff is worth it, you can buy me a scotch whisky in return
// Marco 'don' Kaulea
// ----------------------------------------------------------------------------
package main

import "fmt"
import "strings"
import "github.com/thoj/go-ircevent"

var nickName = "kraulbot"
var channelName = "#kraulbot"

var miauTags = []string{"miau", "maunz", "mrauw", "meow"}

func containsAny(message string, tags []string) bool {
	message = strings.ToLower(message)
	for _, x := range tags {
		if strings.Contains(message, x) {
			return true
		}
	}
	return false
}

func main() {
	con := irc.IRC(nickName, nickName)
	err := con.Connect("irc.hackint.eu:6669")
	if err != nil {
		fmt.Println("Failed connecting")
		fmt.Println(err)
		return
	}
	con.AddCallback("001", func(e *irc.Event) {
		con.Join(channelName)
	})
	con.AddCallback("PRIVMSG", func(e *irc.Event) {
		message := e.Message()
		fmt.Println(message)
		if !strings.HasPrefix(message, nickName) {
			return
		}
		if !containsAny(message, miauTags) {
			return
		}
		con.Action(channelName, fmt.Sprintf("krault %s", e.Nick))
	})
	con.Loop()
}
