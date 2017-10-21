package utils

import "fmt"

func Perror(err error) {
	if err != nil {
		panic(err)
	}
}

func HLine() {
	const hline = "--------------"
	fmt.Println(hline)
}

func Routine(num int, routine func()) {
	for i := 0; i < num; i++ {
		routine()
	}
}

func DeleteUntilLineEnd(newline bool) string {
	var commandDelete = "\u001B[0K"
	if newline {
		commandDelete += "\n"
	}

	fmt.Print(commandDelete)
	return commandDelete
}

type Color string

func (color Color) string() string {
	return string(color)
}

const (
	base   Color = "\u001B["
	reset  Color = base + "0m"
	Black  Color = "0"
	Red    Color = "1"
	Green  Color = "2"
	Yellow Color = "3"
	Blue   Color = "4"
	Purple Color = "5"
	Cyan   Color = "6"
	White  Color = "7"
)

func PrintWithBlink(msg string, color Color) string {
	//fmt.Print("\u001B[37;5m" + msg + "\u001B[0m")
	var str = base.string() + "3" + color.string() + ";5m" + msg + reset.string()
	fmt.Print(str)

	return str
}

func PrintWithColor(msg string, color Color) string {
	//fmt.Print("\u001B[31;1m" + msg + "\u001B[0m")
	var str = base.string() + "3" + color.string() + "m" + msg + reset.string()
	fmt.Print(str)

	return str
}
