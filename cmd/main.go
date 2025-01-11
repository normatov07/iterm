package main

import (
	"fmt"

	"github.com/normatov07/iterm"
)

func main() {

	pass := iterm.Ask("How old are you?: ")
	// fmt.Println(screen.GetTerminalSize())
	println("\nresponse: ", pass)
	// os.Exit(9)
	// mMenu := iterm.NewMultiMenu()
	// mMenu.SetBorder()
	// return
	// pass = iterm.AskSecure("How old are you?: ")
	// // fmt.Println(screen.GetTerminalSize())
	// println("\nresponse: ", pass)

	menu := iterm.NewMultiMenu()
	menu.SetTitle("Which framework do you like?")
	menu.SetBorder()
	menu.Add("Vue")
	menu.Add("React")
	menu.Add("Angular")
	menu.Add("Svelte")

	// command.Print("Hello", color.Black)
	menu.DrawMenu()
	fmt.Println("\n You choosed: ", menu.GetActiveMenu(), "\n")

	menu1 := iterm.NewMenu()
	menu1.SetTitle("Which language do you like?")
	menu1.SetBorder()
	menu1.Add("Go")
	menu1.Add("PHP")
	menu1.Add("Python")
	menu1.Add("Js")
	menu1.Add("Ts")
	menu1.Add("Java")
	menu1.Add("C#")
	menu1.DrawMenu()

	fmt.Println(iterm.RGBColor(42, 161, 179))
	fmt.Println("\n You choosed: ", menu1.GetActiveMenu(), "\n")
	// iterm.NewText("Hello World").Color(color.RED).Println()
	// iterm.NewText("Salom Do'stlar").Color(color.RED).BackgroudColor(color.B_GREEN).Println()
	// fmt.Println("sdgfsd")
	iterm.NewText("Hello World").RGBColor(42, 161, 179).BackgroundRGBColor(0, 0, 0).Println()
}
