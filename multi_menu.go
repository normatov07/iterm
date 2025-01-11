package iterm

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/normatov07/iterm/box"
	"github.com/normatov07/iterm/cursor"
	"github.com/normatov07/iterm/keypass"
)

type multiMenu struct {
	menu
	selected map[int]bool
}

func NewMultiMenu() *multiMenu {
	menu := multiMenu{
		selected: make(map[int]bool),
	}
	menu.terminalScreenSize()

	return &menu
}
func (m *multiMenu) GetActiveMenu() []string {
	selected := make([]string, 0, len(m.selected))
	for i := range m.selected {
		selected = append(selected, m.items[i])
	}

	return selected
}

func (m *multiMenu) GetActiveMenuIndex() []int {
	selected := make([]int, 0, len(m.selected))
	for i := range m.selected {
		selected = append(selected, i)
	}

	return selected
}

func (m *multiMenu) Render() {
	if m.border {
		m.borderedRender()
	} else {
		m.simpleRender()
	}
}

func (m *multiMenu) simpleRender() {
	m.MoveCursorStart()

	if m.title != "" {
		fmt.Printf("  %s%s", RGBColor(m.title, 42, 161, 179), cursor.BEGINING_OF_NEXT_LINE)
	}

	for i, title := range m.items {
		arrow := RGBColor(box.GEO_LARGE_INVERSE_CIRCLE, 86, 84, 114)
		color := 0

		if _, ok := m.selected[i]; ok {
			arrow = RGBColor(box.GEO_LARGE_DIAMOND, 42, 161, 179)
		}

		if i == m.activeIndex {
			color = 100
		}

		fmt.Printf("    %s %s%v", arrow, RGBColor(title, 122+color, 121+color, 133+color), cursor.BEGINING_OF_NEXT_LINE)
	}

	fmt.Print(RGBColor("\nUse the space bar to select options.", 122, 121, 133), cursor.BEGINING_OF_NEXT_LINE)
}

func (m *multiMenu) borderedRender() {
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

	fmt.Print(RGBColor("Use the space bar to select options.", 122, 121, 133), cursor.BEGINING_OF_NEXT_LINE)
}

func (m *multiMenu) drawBorder(content string, line int) {
	switch {
	case line == 0:
		fmt.Printf("┌─ %s %s┐", RGBColor(m.title, 42, 161, 179), strings.Repeat(box.HORIZONTAL, m.tmRowLength-5-utf8.RuneCountInString(content)))
	case line == m.Length()+1:
		fmt.Printf("└%s┘", strings.Repeat(box.HORIZONTAL, m.tmRowLength-2))
	default:
		arrow := ""
		color := 0
		if content != "" {
			arrow = RGBColor(box.GEO_LARGE_INVERSE_CIRCLE, 86, 84, 114)

			if _, ok := m.selected[line-1]; ok {
				arrow = RGBColor(box.GEO_LARGE_DIAMOND, 42, 161, 179)
			}

			if line-1 == m.activeIndex {
				color = 100
			}
		}

		fmt.Printf("│ %s %s %s│", arrow, RGBColor(content, 122+color, 121+color, 133+color), strings.Repeat(" ", m.tmRowLength-utf8.RuneCountInString(content)-6))
	}
}

func (m *multiMenu) DefineMenuOperation(keyPass string) {

	switch keyPass {
	case cursor.DOWN_ONE:
		m.activeIndex = (m.activeIndex + 1) % m.Length()
		m.Render()
	case cursor.UP_ONE:
		m.activeIndex = (m.activeIndex - 1 + m.Length()) % m.Length()
		m.Render()
	}
}

func (m *multiMenu) selectOrDeselect() {
	if _, ok := m.selected[m.activeIndex]; ok {
		delete(m.selected, m.activeIndex)
		return
	}

	m.selected[m.activeIndex] = true
}

func (m *multiMenu) lineCount() int {
	line := m.Length() + 2

	if m.title != "" {
		line += 1
	}

	return line
}

func (m *multiMenu) CleanMenu() {
	fmt.Printf("%v%v", fmt.Sprintf(cursor.UP, m.lineCount()+1), "\033[J")
}

func (m *multiMenu) MoveCursorStart() {
	fmt.Printf("%v%v", fmt.Sprintf(cursor.UP, m.lineCount()), "\033[J")
}

func (m *multiMenu) DrawMenu() (*int, error) {
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
				m.CleanMenu()
				break
			}

			if buffer[0] == keypass.CTRLC || buffer[0] == keypass.CTRLX {
				restoreAndExit(0)
				break
			}

			if buffer[0] == keypass.PROBEL {
				m.selectOrDeselect()
				m.Render()
				continue
			}

			m.DefineMenuOperation(string(buffer))
		}
	}

	return &m.activeIndex, nil
}
