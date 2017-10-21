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
)

const baseDirname = ".metipo"
const materialsDirname = "materials"

func main() {
	var materialsDir = createMetipoDirectory()
	//comm.DownloadFromGitHub(materialsDir)

	var material, err = os.Open(
		filepath.Join(materialsDir.Name(),
			//"saying.txt"))
			"aiueo.txt"))
	defer material.Close()
	utils.Perror(err)

	countDown(0)
	waitKeyInputUntilESC(*material, "[ Please 'ESC' key to quit ]")

	utils.Decorate("-*--*--*--*-", 4)
	fmt.Println(utils.MyBuffer.String())
	utils.Decorate("-*--*--*--*-", 4)
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

	var lineIndex = 0
	var line = lines[lineIndex]
	utils.PrintlnWithColor(line, utils.DarkGray)
	var charIndex = 0
	utils.MyPrintWithBlink("|", utils.LightGray)
	utils.MyPrint("\b")

loop:
	for {
		line = lines[lineIndex]
		var ev = termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyBackspace,
				termbox.KeyBackspace2,
				termbox.KeyDelete,
				termbox.KeyArrowLeft,
				termbox.KeyArrowRight:
				//fmt.Println("[" + string(ev.Ch) + "]")

			case termbox.KeyEsc:
				break loop

			case termbox.KeySpace:
				utils.MyPrint(" ")
				utils.MyPrintWithBlink("|", utils.LightGray)
				utils.MyPrint("\b")

			case termbox.KeyEnter:
				var isLineEnd = charIndex == utf8.RuneCountInString(line)

				if isLineEnd {
					lineIndex += 1
					if lineIndex == len(lines) {
						// カーソルから行末まで削除
						utils.MyDeleteUntilLineEnd(false)
						break loop
					} else {
						// カーソルから行末まで削除
						utils.MyDeleteUntilLineEnd(true)
					}

					line = lines[lineIndex]
					utils.PrintlnWithColor(line, utils.DarkGray)
					utils.MyPrintWithBlink("|", utils.LightGray)
					utils.MyPrint("\b")

					charIndex = 0
				}

			default:
				var input = string(ev.Ch)
				if charIndex < utf8.RuneCountInString(line) {
					var ansChar = string([]rune(line)[charIndex])

					if input == ansChar {
						utils.MyPrint(input)
						charIndex += 1

						utils.MyPrintWithBlink("|", utils.LightGray)
						utils.MyPrint("\b")
					} else {
						// ビープ音
						fmt.Print("\a")

						utils.MyPrintWithBlink("|", utils.LightGray)
						utils.MyPrintWithColor(input, utils.Red)
						utils.MyDeleteUntilLineEnd(false)
						utils.Routine(3, func() {
							utils.MyPrint("\b")
						})
					}
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
