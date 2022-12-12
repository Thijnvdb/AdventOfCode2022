package colors

import "fmt"

const reset = "\033[0m"
const red = "\033[0;31m"
const green = "\033[0;32m"
const yellow = "\033[0;33m"
const blue = "\033[0;34m"
const magenta = "\033[0;35m"
const cyan = "\033[0;36m"
const white = "\033[0;37m"

func Red(input string) string {
	return red + input + reset
}

func Green(input string) string {
	return green + input + reset
}

func Yellow(input string) string {
	return yellow + input + reset
}

func Blue(input string) string {
	return blue + input + reset
}

func Magenta(input string) string {
	return magenta + input + reset
}

func Cyan(input string) string {
	return cyan + input + reset
}

func White(input string) string {
	return white + input + reset
}

func Rgb(str string, r int, g int, b int) string {
	return fmt.Sprintf("\x1b[38;2;%v;%v;%vm", r, g, b) + str + "\x1b[0m"
}
