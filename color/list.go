package color

type Color int

type BColor int

type RGB struct {
	R int
	G int
	B int
}

const (
	BLACK     Color = 30
	RED       Color = 31
	GREEN     Color = 32
	YELLOW    Color = 33
	BLUE      Color = 34
	MAGENTA   Color = 35
	CYAN      Color = 36
	WHITE     Color = 37
	DEFAULT   Color = 39
	SECONDARY Color = 85
	RESET     Color = 0
)

const (
	B_BLACK   BColor = 40
	B_RED     BColor = 41
	B_GREEN   BColor = 42
	B_YELLOW  BColor = 43
	B_BLUE    BColor = 44
	B_MAGENTA BColor = 45
	B_CYAN    BColor = 46
	B_WHITE   BColor = 47
	B_DEFAULT BColor = 49
	B_RESET   BColor = 0
)
