package color

type Color int

const (
	BLACK   Color = 30
	RED     Color = 31
	GREEN   Color = 32
	YELLOW  Color = 33
	BLUE    Color = 34
	MAGENTA Color = 35
	CYAN    Color = 36
	WHITE   Color = 37
	DEFAULT Color = 39
	RESET   Color = 0

	CUSTOM = "\033]38;5;%dm" // %d = [0-255]
)

const (
	B_BLACK   Color = 40
	B_RED     Color = 41
	B_GREEN   Color = 42
	B_YELLOW  Color = 43
	B_BLUE    Color = 44
	B_MAGENTA Color = 45
	B_CYAN    Color = 46
	B_WHITE   Color = 47
	B_DEFAULT Color = 49
	B_RESET   Color = 0

	B_CUSTOM = "\033]48;5;%dm" // %d = [0-255]
)
