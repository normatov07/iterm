package main

import (
	"github.com/normatov13/go-iterm/src/command"
)

func main() {

	menu := command.Menu{}
	menu.Add("Vue")
	menu.Add("React")
	menu.Add("Js")
	menu.Add("Ts")

	menu.DrawMenu()
}
