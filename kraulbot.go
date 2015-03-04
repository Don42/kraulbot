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
import "github.com/belak/irc"

var nickName = "kraulbot"
var channelName = "#kraulbot"

var miauTags = []string{"miau", "maunz", "mrauw", "meow", "nya", "nyā", "mau", "mew", "mrew", "mauw"}

func containsAny(message string, tags []string) bool {
	message = strings.ToLower(message)
	for _, x := range tags {
		if strings.Contains(message, x) {
			return true
		}
	}
	return false
}

func handleIRCMessage(con *irc.Client, e *irc.Event) {
	message := e.Args[1]
	if !strings.HasPrefix(message, nickName) {
		fmt.Println("Not addressed to me")
		return
	}
	if !containsAny(message, miauTags) {
		fmt.Println("No cat")
		return
	}
	action := "kraul"
	switch action {
	case "kraul":
		con.MentionReply(e, "*kraul*")
	default:
		fmt.Println(e)
	}
}

func main() {
	handler := irc.NewBasicMux()
	client := irc.NewClient(irc.HandlerFunc(handler.HandleEvent), nickName, nickName, nickName, "")
	handler.Event("001", func(c *irc.Client, e *irc.Event) {
		c.Write("JOIN #kraulbot")
	})
	handler.Event("PRIVMSG", handleIRCMessage)
	err := client.Dial("irc.hackint.org:6667")
	if err != nil {
		fmt.Println("Failed connecting")
		fmt.Println(err)
		return
	}
}
