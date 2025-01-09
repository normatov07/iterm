package iterm

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/normatov07/iterm/cursor"
	"github.com/normatov07/iterm/keypass"
)

type MenuData interface {
	Add(title string)
	Del(index int)
}

type Menu struct {
	items       []string
	border      bool
	title       string
	activeIndex int
	tmRowLength int
}

func NewMenu(title string) *Menu {
	return &Menu{
		title:       title,
		tmRowLength: 50,
	}
}

func (m *Menu) SetTitle(title string) {
	m.title = title
}

func (m *Menu) SetBorder() {
	m.border = true
}

func (m *Menu) SetList(list []string) {
	m.items = list
}

func (m *Menu) Add(title string) {
	m.items = append(m.items, title)
}

func (m *Menu) Del(index int) {
	m.items = append(m.items[:index], m.items[index:]...)
}

func (m *Menu) Length() int {
	return len(m.items)
}

func (m *Menu) lineCount() int {
	line := m.Length()

	if m.title != "" {
		line += 2
	}

	return line
}

func (m *Menu) Render() {
	if m.border {
		m.borderedRender()
	} else {
		m.simpleRender()
	}
}

func (m *Menu) simpleRender() {
	m.MoveCursorStart()

	if m.title != "" {
		fmt.Printf("    %s\n%s", m.title, cursor.BEGINING_OF_NEXT_LINE)

	}

	for i, title := range m.items {
		if i == m.activeIndex {
			title += " <"
		}
		fmt.Printf("  %d) %s%v", i+1, title, cursor.BEGINING_OF_NEXT_LINE)
	}
}

func (m *Menu) borderedRender() {
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

func (m *Menu) drawBorder(content string, line int) {
	switch {
	case line == 0:
		fmt.Printf("┌%s %s %s┐", strings.Repeat("─", 3), RGBColor(m.title, 42, 161, 179), strings.Repeat("─", m.tmRowLength-7-len(content)))
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

		fmt.Printf("│ %s %s %s│", arrow, RGBColor(content, 122+color, 121+color, 133+color), strings.Repeat(" ", m.tmRowLength-len(content)-6))
	}
}

func (m *Menu) MoveCursorStart() {
	fmt.Printf("%v%v", fmt.Sprintf(cursor.UP, m.lineCount()), "\033[J")
}

func (m *Menu) DefineMenuOperation(keyPass string) {
	switch keyPass {
	case cursor.DOWN_ONE:
		m.activeIndex = (m.activeIndex + 1) % m.Length()
		m.Render()
	case cursor.UP_ONE:
		m.activeIndex = (m.activeIndex - 1 + m.Length()) % m.Length()
		m.Render()
	}
}

func (m *Menu) DrawMenu() (*int, error) {
	m.tmRowLength = 50
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
