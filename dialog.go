package iterm

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/normatov07/iterm/cursor"
	"github.com/normatov07/iterm/keypass"
	"github.com/normatov07/iterm/screen"
)

func AskSecure(question string) (respond string) {

	err := DisableOutput()

	if err != nil {
		panic(err)
	}

	defer RestoreTermMode()

	fmt.Print(question)

	reader := bufio.NewReader(os.Stdin)

	respond, err = reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	return
}

func Ask(question string) (respond string) {

	fmt.Print(question)

	fmt.Scan(&respond)

	return
}

func AskLine(question string) (respond string) {

	fmt.Print(question)

	reader := bufio.NewReader(os.Stdin)

	respond, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	return
}

func AskSecurePatterned(question string) string {
	err := SetRawMode()

	if err != nil {
		panic(err)
	}

	defer RestoreTermMode()

	fmt.Print(question)

	reader := make([]byte, 3)

	var respond []rune
	for {

		_, err := os.Stdin.Read(reader)

		if err != nil && !errors.Is(err, io.EOF) {
			panic(err)
		}

		if reader[0] == keypass.ENTER {
			if len(respond) <= 0 {
				continue
			}

			break
		}

		if reader[0] == keypass.CTRLC || reader[0] == keypass.CTRLX {
			restoreAndExit(0)
			break
		}

		if reader[0] == keypass.BACKSPACE || reader[0] == keypass.CTRLBACKSPACE {
			if len(respond) > 0 {
				respond = respond[:len(respond)-1]
				fmt.Print(screen.CLEAR_LINE)
				fmt.Printf(cursor.LEFT, len(respond)+len(question)+1)
				fmt.Print(question, strings.Repeat("*", len(respond)))
			}

			continue
		}

		respond = append(respond, rune(reader[0]))
		fmt.Print("*")
	}

	return string(respond)
}
