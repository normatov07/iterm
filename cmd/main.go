package main

import (
	"fmt"

	"github.com/normatov13/iterm/src/codes/color"
	"github.com/normatov13/iterm/src/command"
)

func main() {

	menu := command.Menu{}
	menu.Add("Vue")
	menu.Add("React")
	menu.Add("Js")
	menu.Add("Ts")

	// command.Print("Hello", color.Black)
	// menu.DrawMenu()
	command.NewText("Hello World").Color(color.RED).Println()
	fmt.Println("sdgfsd")
}
