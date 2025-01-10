package screen

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type TerminalSize struct {
	Width  int
	Height int
}

func GetTerminalSize() (*TerminalSize, error) {
	if runtime.GOOS == "windows" {
		return getTerminalSizeWindows()
	}

	return getTerminalSizeUnix()

}

func getTerminalSizeWindows() (*TerminalSize, error) {
	// TODO: add support for windows

	return nil, nil
}

func getTerminalSizeUnix() (*TerminalSize, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	fields := strings.Fields(string(output))
	if len(fields) == 2 {
		height, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, err
		}
		width, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, err
		}

		return &TerminalSize{
			Width:  width,
			Height: height,
		}, nil
	}
	return nil, nil
}
