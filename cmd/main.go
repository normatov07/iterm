package main

import (
	"fmt"

	"github.com/normatov07/iterm"
)

func main() {

	// fmt.Println(screen.GetTerminalSize())
	// os.Exit(9)

	menu := iterm.NewMenu()
	menu.SetTitle("Which framework do you like?")
	menu.SetBorder()
	menu.Add("it contain info about your terminal")
	menu.Add("React")
	menu.Add("Angular")
	menu.Add("Svelte")

	// command.Print("Hello", color.Black)
	menu.DrawMenu()
	fmt.Println("\n You choosed: ", menu.GetActiveMenu(), "\n")

	menu = iterm.NewMenu()
	menu.SetTitle("Which language do you like?")
	menu.SetBorder()
	menu.Add("Go")
	menu.Add("PHP")
	menu.Add("Python")
	menu.Add("Js")
	menu.Add("Ts")
	menu.Add("Java")
	menu.Add("C#")
	menu.DrawMenu()
	fmt.Println("\n You choosed: ", menu.GetActiveMenu(), "\n")
	// iterm.NewText("Hello World").Color(color.RED).Println()
	// iterm.NewText("Salom Do'stlar").Color(color.RED).BackgroudColor(color.B_GREEN).Println()
	// fmt.Println("sdgfsd")
}
