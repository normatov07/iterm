package iterm

import (
	"os"

	"golang.org/x/sys/unix"
)

var defaultTermios unix.Termios

func SetRawMode() (err error) {
	fd := int(os.Stdin.Fd())

	term, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	if err != nil {
		return err
	}

	defaultTermios = *term

	term.Iflag &^= unix.IGNBRK | unix.BRKINT | unix.PARMRK | unix.ISTRIP | unix.INLCR | unix.IGNCR | unix.ICRNL | unix.IXON
	term.Oflag &^= unix.OPOST
	term.Lflag &^= unix.ECHO | unix.ECHONL | unix.ICANON | unix.ISIG | unix.IEXTEN
	term.Cflag &^= unix.CSIZE | unix.PARENB
	term.Cflag |= unix.CS8

	err = unix.IoctlSetTermios(fd, unix.TCSETS, term)
	if err != nil {
		return err
	}

	return nil
}

func RestoreTermMode() error {
	return unix.IoctlSetTermios(int(os.Stdin.Fd()), unix.TCSETS, &defaultTermios)
}
