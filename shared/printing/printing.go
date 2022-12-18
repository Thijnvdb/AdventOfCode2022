package printing

func HideCursor() {
	print("\033[?25l")
}

// Clear screen and move cursor to 0,0
func ResetScreen() {
	print("\x1b[2J\033[0;0H")
}

func ShowCursor() {
	print("\033[?25h")
}
