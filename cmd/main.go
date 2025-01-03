package main

import (
	"fmt"

	"github.com/normatov07/iterm"
	"github.com/normatov07/iterm/color"
)

func main() {

	menu := iterm.Menu{}
	menu.Add("Vue")
	menu.Add("React")
	menu.Add("Js")
	menu.Add("Ts")

	// command.Print("Hello", color.Black)
	// menu.DrawMenu()
	iterm.NewText("Hello World").Color(color.RED).Println()
	fmt.Println("sdgfsd")
}
