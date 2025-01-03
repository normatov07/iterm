# iTerm CLI Toolkit

## Overview

`iTerm` is a powerful Golang package designed to enhance terminal text styling and menu creation. It provides an intuitive, fluent interface for creating colorful and styled terminal outputs with ease.

See for more infromation: https://pkg.go.dev/github.com/normatov07/iterm

## Features

- 🎨 **Text Styling**
  - Change text color
  - Set background colors
  - Apply text formatting (bold, italic, underline)
  - Support for 8 standard terminal colors
  - Custom color support

- 📋 **Menu Management**
  - Dynamic menu creation
  - Easy menu item addition

- 🖥️ **Change terminal mode**
  - Raw mode
  - Normal mode


## Installation

```bash
go get github.com/normatov07/iterm
```


## Usage Examples

### Creating a menu

easy and simple creating menu and get the index of the active item

```go
menu := iterm.Menu{}
menu.Add("Vue")
menu.Add("React")
menu.Add("JavaScript")
menu.Add("TypeScript")

// Draw the menu (implementation details may vary)
index, err := menu.DrawMenu()
if err != nil {
   panic(err)
}
fmt.Println("Active item index:", index)
fmt.Println("Active item index:", menu.activeIndex)
```

Text Styling

```go
// Basic colored text
iterm.NewText("Hello World").Color(color.RED).Println()

// Advanced styling
iterm.NewText("Styled Text")
    .Color(color.GREEN)
    .Bold(true)
    .Italic(true)
    .UnderLine(true)
    .Println()
```

Change terminal mode

```go
iterm.SetRawMode()
iterm.RestoreTermMode()
```

### License
Distributed under the MIT License. See LICENSE for more information.

### Contact
O'tkir Normatov - normatov.milliy@gmail.com

