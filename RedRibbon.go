package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	. "github.com/gbin/goncurses"
)

type ChatMsg struct {
	Channel  string
	Msg      string
	Username string
}

func DecodeMsg(jsonstr string) ChatMsg {
	var msg ChatMsg
	json.Unmarshal([]byte(jsonstr), &msg) // basically converts the json string to a struct of type ChatMsg
	return msg
}

func httpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR. Termiating program...")
		os.Exit(3)
	}
	// We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR. Termiating program...")
		os.Exit(3)
	}
	// Convert the body to type string
	sb := string(body)
	return sb
}

func main() {
	// Curses init
	stdscr, err := Init()
	Raw(true)

	if err != nil {
		log.Fatal(err)
	}

	// <CODE>
	jsonreceived := httpGet("https://redribbonservers.jort57.repl.co/jsontest")
	stdscr.Println("Fetching JSON data...")
	stdscr.Printf("JSON received from redribbon servers: %s\n", jsonreceived)
	stdscr.Printf("Decoding JSON...\n")
	jsonDecoded := DecodeMsg(jsonreceived)
	stdscr.Printf("Username: %s\n", jsonDecoded.Username)
	stdscr.Printf("Channel: %s\n", jsonDecoded.Channel)
	stdscr.Printf("Message: %s\n", jsonDecoded.sg)

	// </CODE>

	// Wait for key to be pressed to end
	stdscr.Refresh()
	stdscr.GetChar()

	End()
}
