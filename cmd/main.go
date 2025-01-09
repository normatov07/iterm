package main

import (
	"github.com/normatov07/iterm"
)

func main() {

	menu := iterm.Menu{}
	menu.SetTitle("Which framework do you like?")
	menu.SetBorder()
	menu.Add("Vue")
	menu.Add("React")
	menu.Add("Js")
	menu.Add("Ts")

	// command.Print("Hello", color.Black)
	menu.DrawMenu()
	// iterm.NewText("Hello World").Color(color.RED).Println()
	// iterm.NewText("Salom Do'stlar").Color(color.RED).BackgroudColor(color.B_GREEN).Println()
	// fmt.Println("sdgfsd")
}
