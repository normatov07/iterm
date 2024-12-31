package cursor

const (
	HOME                  = "\033[H"
	UP_ONE                = "\033[A"
	UP                    = "\033[%dA"
	DOWN                  = "\033[%dB"
	DOWN_ONE              = "\033[B"
	RIGHT                 = "\033[%dC"
	LEFT                  = "\033[D"
	MOVE                  = "\033[%d;%dH"
	TO_LINE               = "\033[%dG"
	BEGINING_OF_NEXT_LINE = "\033[E"

	HIDE_CURSOR = "\033[?25l"
	SHOW_CURSOR = "\033[?25h"
)
