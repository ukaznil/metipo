package utils

import (
	"bytes"
	"fmt"
)

var MyBuffer bytes.Buffer

func MyPrint(str string) {
	fmt.Print(str)
	MyBuffer.WriteString(str)
}

func MyPrintln(str string) {
	fmt.Println(str)
	MyBuffer.WriteString(str + "\n")
}

func MyPrintWithBlink(str string, color Color) {
	MyBuffer.WriteString(PrintWithBlink(str, color))
}

func MyPrintWithColor(str string, color Color) {
	MyBuffer.WriteString(PrintWithColor(str, color))
}

func MyDeleteUntilLineEnd(newline bool) {
	MyBuffer.WriteString(DeleteUntilLineEnd(newline))
}
