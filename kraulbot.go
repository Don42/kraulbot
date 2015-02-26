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

func main() {
	con := irc.IRC("kraulbot", "kraulbot")
	err := con.Connect("irc.hackint.eu:6669")
	if err != nil {
		fmt.Println("Failed connecting")
		fmt.Println(err)
		return
	}
	con.AddCallback("001", func(e *irc.Event) {
		con.Join("#kraulbot")
	})
	con.AddCallback("PRIVMSG", func(e *irc.Event) {
		message := e.Message()
		fmt.Println(message)
		if !strings.HasPrefix(message, "kraulbot") {
			return
		}
		if !strings.Contains(strings.ToLower(message), "miau") {
			return
		}
		con.Privmsg("#kraulbot", fmt.Sprintf("%s *kraul*", e.Nick))
	})
	con.Loop()
}
