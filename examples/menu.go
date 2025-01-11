package main

import (
	"fmt"

	"github.com/normatov07/iterm"
)

func main() {

	pass := iterm.AskSecure("Password?: ")
	println("\npassword: ", pass)

	menu := iterm.NewMultiMenu()
	menu.SetTitle("Which framework do you like?")
	menu.SetBorder()
	menu.Add("Vue")
	menu.Add("React")
	menu.Add("Angular")
	menu.Add("Svelte")
	menu.DrawMenu()
	fmt.Println("  You choosed: ", menu.GetActiveMenu())

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
	fmt.Println("\n You choosed: ", menu1.GetActiveMenu())
}
