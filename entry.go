package main

import (
	"os/user"
	"os"
	"path/filepath"
	"github.com/nsf/termbox-go"
	"fmt"
	"github.com/ukaznil/metipo/utils"
	"time"
	"strconv"
	"bufio"
	"unicode/utf8"
	"bytes"
)

const baseDirname = ".metipo"
const materialsDirname = "materials"

var buffer bytes.Buffer

func bufferPrint(str string) {
	fmt.Print(str)
	buffer.WriteString(str)
}

func bufferPrintln(str string) {
	fmt.Println(str)
	buffer.WriteString(str + "\n")
}

func bufferPrintWithBlink(str string, color utils.Color) {
	buffer.WriteString(utils.PrintWithBlink(str, color))
}

func bufferPrintWithColor(str string, color utils.Color) {
	buffer.WriteString(utils.PrintWithColor(str, color))
}

func bufferDeleteUntilLineEnd(newline bool) {
	buffer.WriteString(utils.DeleteUntilLineEnd(newline))
}

func main() {
	var materialsDir = createMetipoDirectory()
	//comm.DownloadFromGitHub(materialsDir)

	var material, err = os.Open(
		filepath.Join(materialsDir.Name(),
			"saying.txt"))
	defer material.Close()
	utils.Perror(err)

	countDown(0)
	waitKeyInputUntilESC(*material, "[ Please 'ESC' key to quit ]")

	fmt.Println(buffer.String())
}

func countDown(sec int) {
	for i := 0; i < sec; i++ {
		fmt.Print(strconv.Itoa(sec-i) + "\r")
		time.Sleep(1 * time.Second)
	}

	fmt.Println("!! MeTipo !!")
	time.Sleep(1 * time.Second)
}

func waitKeyInputUntilESC(material os.File, msg string) {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	fmt.Println(msg)
	utils.HLine()

	var scanner = bufio.NewScanner(&material)
	utils.Perror(scanner.Err())
	var lines = make([]string, 0)
	for scanner.Scan() {
		var line = scanner.Text()
		lines = append(lines, line)
	}
	//fmt.Println(lines)

	var lineIndex = 0
	var line = lines[lineIndex]
	fmt.Println(line)
	var charIndex = 0

loop:
	for {
		line = lines[lineIndex]
		var ev = termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyBackspace,
				termbox.KeyBackspace2,
				termbox.KeyDelete:
				//fmt.Println("[" + string(ev.Ch) + "]")

			case termbox.KeyEsc:
				break loop

			case termbox.KeySpace:
				//fmt.Print(" ")
				bufferPrint(" ")
				//utils.PrintWithBlink("|", utils.White)
				bufferPrintWithBlink("|", utils.White)
				//fmt.Print("\b")
				bufferPrint("\b")

			case termbox.KeyEnter:
				var isLineEnd = charIndex == utf8.RuneCountInString(line)

				if isLineEnd {
					// カーソルから行末まで削除
					//utils.DeleteUntilLineEnd(true)
					bufferDeleteUntilLineEnd(true)
					lineIndex += 1

					if lineIndex == len(lines) {
						break loop
					}

					line = lines[lineIndex]
					//fmt.Println(line)
					bufferPrintln(line)
					//utils.PrintWithBlink("|", utils.White)
					bufferPrintWithBlink("|", utils.White)
					//fmt.Print("\b")
					bufferPrint("\b")

					charIndex = 0
				}

			default:
				var input = string(ev.Ch)
				var ansChar = string([]rune(line)[charIndex])

				if input == ansChar {
					//fmt.Print(input)
					bufferPrint(input)
					charIndex += 1
					//utils.PrintWithBlink("|", utils.White)
					bufferPrintWithBlink("|", utils.White)
					//fmt.Print("\b")
					bufferPrint("\b")
				} else {
					//utils.PrintWithBlink("|", utils.White)
					bufferPrintWithBlink("|", utils.White)
					//utils.PrintWithColor(input, utils.Red)
					bufferPrintWithColor(input, utils.Red)
					//utils.DeleteUntilLineEnd(false)
					bufferDeleteUntilLineEnd(false)
					utils.Routine(3, func() {
						//fmt.Print("\b")
						bufferPrint("\b")
					})
				}
			}
		}
	}
	termbox.Sync()
}

func createMetipoDirectory() os.File {
	var usr, err = user.Current()
	utils.Perror(err)

	var dir = filepath.Join(usr.HomeDir, baseDirname, materialsDirname)
	if _, err := os.Stat(dir); err != nil {
		if err := os.Mkdir(dir, 0744); err != nil {
			panic(err)
		}
	}

	if ret, err := os.Open(dir); err != nil {
		panic(err)
	} else {
		return *ret
	}
}
