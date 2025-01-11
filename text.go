package iterm

import (
	"fmt"

	"github.com/normatov07/iterm/color"
)

type Text struct {
	color  color.Color
	tRGB   color.RGB
	bColor color.BColor
	bRGB   color.RGB
	bold   bool
	uline  bool
	italic bool
	text   string
}

func NewText(text string) *Text {
	return &Text{text: text}
}

func (t *Text) Color(color color.Color) *Text {
	t.color = color

	return t
}

func (t *Text) BackgroudColor(color color.BColor) *Text {
	t.bColor = color

	return t
}

func (t *Text) Bold() *Text {
	t.bold = true

	return t
}

func (t *Text) Italic() *Text {
	t.italic = true

	return t
}

func (t *Text) UnderLine() *Text {
	t.uline = true

	return t
}

func (t *Text) GetStyledText() string {
	return t.parseStyle()
}

func (t *Text) RGBColor(r, g, b int) *Text {
	t.tRGB = color.RGB{R: r, G: g, B: b}

	return t
}
func (t *Text) BackgroundRGBColor(r, g, b int) *Text {
	t.bRGB = color.RGB{R: r, G: g, B: b}

	return t
}

func (t *Text) Print() {
	fmt.Print(t.parseStyle())
}

func (t *Text) Println() {
	fmt.Println(t.parseStyle())
}

func (t *Text) parseStyle() string {
	text := t.text

	switch {
	case t.bold:
		text = fmt.Sprintf("\033[1m%s", text)
		fallthrough
	case t.italic:
		text = fmt.Sprintf("\033[3m%s", text)
		fallthrough
	case t.uline:
		text = fmt.Sprintf("\033[4m%s", text)
		fallthrough
	case t.color > 0:
		text = fmt.Sprintf("\033[%dm%s", t.color, text)
		fallthrough
	case t.bColor > 0:
		text = fmt.Sprintf("\033[%dm%s", t.bColor, text)
		fallthrough
	case t.tRGB.R > 0 && t.tRGB.G > 0 && t.tRGB.B > 0:
		text = fmt.Sprintf("\033[38;2;%d;%d;%dm%s", t.tRGB.R, t.tRGB.G, t.tRGB.B, text)
		fallthrough
	case t.bRGB.R > 0 && t.bRGB.G > 0 && t.bRGB.B > 0:
		text = fmt.Sprintf("\033[48;2;%d;%d;%dm%s", t.bRGB.R, t.bRGB.G, t.bRGB.B, text)
	}

	text = fmt.Sprintf("%s\033[0m", text)

	return text
}

func Colorful(text string, color color.Color) string {
	return fmt.Sprintf("\033[%dm%s\033[0m", color, text)
}

func RGBColor(text string, r, g, b int) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, text)
}

func BackgroundRGBColor(text string, r, g, b int) string {
	return fmt.Sprintf("\033[48;2;%d;%d;%dm%s\033[0m", r, g, b, text)
}
