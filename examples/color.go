package main

import (
	"fmt"

	"github.com/normatov07/iterm"
	"github.com/normatov07/iterm/color"
)

func main() {
	fmt.Println(iterm.RGBColor("RGB colored text", 42, 161, 179))

	fmt.Println(iterm.BackgroundRGBColor("RGB colored text", 42, 161, 179))

	iterm.NewText("Hello World").Color(color.RED).Println()

	iterm.NewText("Salom Do'stlar").Color(color.RED).BackgroudColor(color.B_GREEN).Println()

	iterm.NewText("Hello World").RGBColor(42, 161, 179).BackgroundRGBColor(0, 0, 0).Println()
}
