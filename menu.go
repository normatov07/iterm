package iterm

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/normatov07/iterm/cursor"
	"github.com/normatov07/iterm/keypass"
	"github.com/normatov07/iterm/screen"
)

type MenuData interface {
	Add(title string)
	Del(index int)
}

type menu struct {
	items       []string
	border      bool
	title       string
	activeIndex int
	tmRowLength int
}

func NewMenu() *menu {
	menu := menu{}
	menu.terminalScreenSize()

	return &menu
}

func (m *menu) SetTitle(title string) {
	m.title = m.responsiveText(title)
}

func (m *menu) SetBorder() {
	m.border = true
}

func (m *menu) SetList(list []string) {
	m.items = list
}

func (m *menu) Add(title string) {
	m.items = append(m.items, m.responsiveText(title))
}

func (m *menu) Del(index int) {
	m.items = append(m.items[:index], m.items[index:]...)
}

func (m *menu) Get(index int) string {
	if index >= m.Length() {
		panic("index out of range")
	}

	return m.items[index]
}
func (m *menu) GetActiveMenu() string {
	return m.items[m.activeIndex]
}

func (m *menu) Length() int {
	return len(m.items)
}

func (m *menu) lineCount() int {
	line := m.Length()

	if m.title != "" {
		if m.border {
			line += 1
		}
		line += 1
	}

	return line
}

func (m *menu) Render() {
	if m.border {
		m.borderedRender()
	} else {
		m.simpleRender()
	}
}

func (m *menu) simpleRender() {
	m.MoveCursorStart()

	if m.title != "" {
		fmt.Printf("  %s%s", RGBColor(m.title, 42, 161, 179), cursor.BEGINING_OF_NEXT_LINE)
	}

	for i, title := range m.items {
		arrow := RGBColor("○", 86, 84, 114)
		color := 0
		if i == m.activeIndex {
			arrow = RGBColor("◉", 42, 161, 179)
			color = 100
		}
		fmt.Printf("    %s %s%v", arrow, RGBColor(title, 122+color, 121+color, 133+color), cursor.BEGINING_OF_NEXT_LINE)
	}
}

func (m *menu) borderedRender() {
	m.MoveCursorStart()
	length := m.Length() + 1

	for i := 0; i <= length; i++ {
		if i == 0 {
			m.drawBorder(m.title, i)
		} else if length == i {
			m.drawBorder("", i)
		} else {
			m.drawBorder(m.items[i-1], i)
		}

		fmt.Print(cursor.BEGINING_OF_NEXT_LINE)
	}
}

func (m *menu) drawBorder(content string, line int) {
	switch {
	case line == 0:
		fmt.Printf("┌─ %s %s┐", RGBColor(m.title, 42, 161, 179), strings.Repeat("─", m.tmRowLength-5-utf8.RuneCountInString(content)))
	case line == m.Length()+1:
		fmt.Printf("└%s┘", strings.Repeat("─", m.tmRowLength-2))
	default:
		arrow := ""
		color := 0
		if content != "" {
			arrow = RGBColor("○", 86, 84, 114)
			if line-1 == m.activeIndex {
				arrow = RGBColor("◉", 42, 161, 179)
				color = 100
			}
		}

		fmt.Printf("│ %s %s %s│", arrow, RGBColor(content, 122+color, 121+color, 133+color), strings.Repeat(" ", m.tmRowLength-utf8.RuneCountInString(content)-6))
	}
}

func (m *menu) MoveCursorStart() {
	fmt.Printf("%v%v", fmt.Sprintf(cursor.UP, m.lineCount()), "\033[J")
}

func (m *menu) DefineMenuOperation(keyPass string) {
	switch keyPass {
	case cursor.DOWN_ONE:
		m.activeIndex = (m.activeIndex + 1) % m.Length()
		m.Render()
	case cursor.UP_ONE:
		m.activeIndex = (m.activeIndex - 1 + m.Length()) % m.Length()
		m.Render()
	}
}

func (m *menu) DrawMenu() (*int, error) {
	err := SetRawMode()
	if err != nil {
		return nil, err
	}

	defer func() {
		RestoreTermMode()
		print(cursor.SHOW_CURSOR)
	}()

	print(cursor.HIDE_CURSOR, strings.Repeat("\n", m.lineCount()+1))

	m.Render()

	buffer := make([]byte, 3)
	for {
		len, err := os.Stdin.Read(buffer)
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, err
		}

		if len > 0 {
			if buffer[0] == keypass.ENTER {
				break
			}
			if buffer[0] == keypass.CTRLC || buffer[0] == keypass.CTRLX {
				// --- TODO: exit interface
				break
			}

			m.DefineMenuOperation(string(buffer))
		}

	}

	return &m.activeIndex, nil
}

func (m *menu) terminalScreenSize() {
	screen, err := screen.GetTerminalSize()
	if err != nil || screen.Width > 66 {
		m.tmRowLength = 66
		return
	}

	m.tmRowLength = screen.Width
}

func (m *menu) responsiveText(text string) string {
	if m.tmRowLength-10 < utf8.RuneCountInString(text) {
		return text[:m.tmRowLength-9] + "\u2026"
	}

	return text
}
