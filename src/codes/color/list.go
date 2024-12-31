package color

const (
	BLACK   = 30
	RED     = 31
	GREEN   = 32
	YELLOW  = 33
	BLUE    = 34
	MAGENTA = 35
	CYAN    = 36
	WHITE   = 37
	DEFAULT = 39
	RESET   = 0

	CUSTOM = "\033]38;5;%dm" // %d = [0-255]
)

const (
	B_BLACK   = 40
	B_RED     = 41
	B_GREEN   = 42
	B_YELLOW  = 43
	B_BLUE    = 44
	B_MAGENTA = 45
	B_CYAN    = 46
	B_WHITE   = 47
	B_DEFAULT = 49
	B_RESET   = 0

	B_CUSTOM = "\033]48;5;%dm" // %d = [0-255]
)
