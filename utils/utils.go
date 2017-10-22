package utils

import (
	"fmt"
	"bytes"
)

func Perror(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func Decorate(component string, num int) {
	var buffer bytes.Buffer
	for i := 0; i < num; i++ {
		buffer.WriteString(component)
	}

	fmt.Println(buffer.String())
}

func HLine() {
	//const hline = "--------------"
	//fmt.Println(hline)
	Decorate("----", 4)
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
	base  Color = "\u001B["
	reset Color = base + "0m"

	Black     Color = "0;30"
	Red       Color = "0;31"
	Green     Color = "0;32"
	Orange    Color = "0;33"
	Blue      Color = "0;34"
	Purple    Color = "0;35"
	Cyan      Color = "0;36"
	LightGray Color = "0;37"

	DarkGray    Color = "1;30"
	LightRed    Color = "1;31"
	LightGreen  Color = "1;32"
	Yello       Color = "1;33"
	LightBlue   Color = "1;34"
	LightPurple Color = "1;35"
	LightCyan   Color = "1;36"
	White       Color = "1;37"
)

func PrintWithBlink(msg string, color Color) string {
	var str = base.string() + color.string() + ";5m" + msg + reset.string()
	fmt.Print(str)

	return str
}

func PrintWithColor(msg string, color Color) string {
	var str = base.string() + color.string() + "m" + msg + reset.string()
	fmt.Print(str)

	return str
}

func PrintlnWithColor(msg string, color Color) string {
	return PrintWithColor(msg+"\n", color)
}
