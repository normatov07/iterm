package command

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/normatov13/iterm/src/codes/cursor"
	"github.com/normatov13/iterm/src/codes/keypass"
)

type MenuData interface {
	Add(title string)
	Del(index int)
}

type Menu struct {
	items       []string
	activeIndex int
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

func (m *Menu) Render() {
	m.MoveCursorStart()

	for i, title := range m.items {
		if i == m.activeIndex {
			title += " <"
		}
		fmt.Printf("%d) %s%v", i+1, title, cursor.BEGINING_OF_NEXT_LINE)
	}
}

func (m *Menu) MoveCursorStart() {
	fmt.Printf("%v%v", fmt.Sprintf(cursor.UP, m.Length()), "\033[J")
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
	err := SetRawMode()
	if err != nil {
		return nil, err
	}

	defer func() {
		RestoreTermMode()
		print(cursor.SHOW_CURSOR)
	}()

	print(cursor.HIDE_CURSOR, strings.Repeat("\n", m.Length()))
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
