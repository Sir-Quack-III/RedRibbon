package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/marcusolsson/tui-go"
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
	sidebar := tui.NewVBox(
		tui.NewLabel("Congregations:"),
		tui.NewLabel("Private Assemblies:"),
	)

	ui, err := tui.New(sidebar)
	if err != nil {
		panic(err)
	}
	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
