# iTerm CLI Toolkit

## Overview

`iTerm` is a powerful Golang package designed to enhance terminal text styling and menu creation. It provides an intuitive, fluent interface for creating colorful and styled terminal outputs with ease.

See for more infromation: https://pkg.go.dev/github.com/normatov07/iterm

## Features

- 🎨 **Text Styling**
  - Change text color
  - Set background colors
  - Apply text formatting (bold, italic, underline)
  - Support for 8 standard and RGB colors

- 📋 **Menu Management**
  - Dynamic and responsive menu creation
  - Easy menu item addition
  - Single selection
  - Multi selection
  - Beauty bordered menu

- 🖥️ **Change terminal mode**
  - Raw mode
  - Normal mode
  - Disable output on terminal
  - Restore default mode

- **Helper terminal codes**
  - ANSI escape sequences
  - Box drawing codes

## Installation

```bash
go get github.com/normatov07/iterm
```


## Usage Examples

### Creating a menu

easy and simple creating menu and get the index of the active item

```go
import (
    "fmt"
    "github.com/normatov07/iterm"
)

func main() {
    menu := iterm.NewMenu()
    menu.SetTitle("Which framework do you like?")
    menu.SetBorder()
    menu.Add("Vue")
    menu.Add("React")
    menu.Add("Angular")
    menu.Add("Svelte")

    index, err := menu.DrawMenu()
    if err != nil {
        panic(err)
    }
    fmt.Println("Active item index:", index)
    fmt.Println("Active item:", menu.GetActiveMenu())
}
```

### Creating a Multi-Select Menu

You can also create a multi-select menu where users can select multiple items.

```go
import (
    "fmt"
    "github.com/normatov07/iterm"
)

func main() {
    multiMenu := iterm.NewMultiMenu()
    multiMenu.SetTitle("Select your favorite languages")
    multiMenu.SetBorder()
    multiMenu.Add("Go")
    multiMenu.Add("Python")
    multiMenu.Add("JavaScript")
    multiMenu.Add("Rust")

    _, err := multiMenu.DrawMenu()
    if err != nil {
        panic(err)
    }
    fmt.Println("Selected items:", multiMenu.GetActiveMenu())
}
```

### Text Styling

  We have lots of patters to give elegant seen to text on terminal

```go
import (
    "github.com/normatov07/iterm"
    "github.com/normatov07/iterm/color"
)

func main() {
    // Basic colored text
    iterm.NewText("Hello World").Color(color.RED).Println()

    // Advanced styling
    iterm.NewText("Styled Text").
        Color(color.GREEN).
        Bold().
        Italic().
        UnderLine().
        Println()
  
  // RGB color
    iterm.NewText("Hello World").RGBColor(42, 161, 179).BackgroundRGBColor(0, 0, 0).Println()

    println(iterm.RGBColor("RGB colored text",42, 161, 179))
	  println(iterm.BackgroundRGBColor("RGB back colored text",42, 161, 179))

}
```
### Ask Secure And Insecure Inputs
You can ask for secure and insecure input (e.g: text, passwords) displaying or without displaying the input or patterned on the terminal.

```go
    password := iterm.AskSecurePatterned("Enter your password: ")
    fmt.Println("\nPassword entered:", password)

    age := iterm.Ask("How old are you?: ")
    fmt.Println("Password entered:", age)
```

#### Change terminal mode

```go
iterm.SetRawMode()
defer iterm.RestoreTermMode()
```

### License
Distributed under the MIT License. See LICENSE for more information.

### Contact
O'tkir Normatov - normatov.milliy@gmail.com

