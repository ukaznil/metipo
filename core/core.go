package core

import (
	"fmt"
	"github.com/ukaznil/metipo/utils"
	"github.com/nsf/termbox-go"
	"os"
	"bufio"
	"unicode/utf8"
	"io/ioutil"
	"math/rand"
	"path/filepath"
)

func selectMaterialName() string {
	var materialsDir = getOrCreateMaterialsDirectory()
	var items, err = ioutil.ReadDir(materialsDir.Name())
	utils.Perror(err)

	var files []os.File
	for _, item := range items {
		if !item.IsDir() {
			var file, err = os.Open(filepath.Join(materialsDir.Name(), item.Name()))
			utils.Perror(err)
			files = append(files, *file)
		}
	}

	var file = files[rand.Intn(len(files))]
	return file.Name()
}

func Exercise() {
	//comm.DownloadFromGitHub(materialsDir)
	var material, err = os.Open(selectMaterialName())
	defer material.Close()
	utils.Perror(err)

	countDown(0)
	var stats = waitKeyInputUntilESC(*material, "[ Please 'ESC' key to quit ]")

	utils.Decorate("------------", 4)
	fmt.Println(utils.MyBuffer.String())
	utils.Decorate("-*--*--*--*-", 4)
	fmt.Print(stats.String())
	utils.Decorate("-*--*--*--*-", 4)
}

func waitKeyInputUntilESC(material os.File, msg string) *utils.Stats {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	fmt.Println(msg)
	//utils.HLine()
	utils.Decorate("-*--*--*--*-", 4)

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
	//utils.Decorate("------------", 4)
	var charIndex = 0
	utils.MyPrintWithBlink("|", utils.LightGray)
	utils.MyPrint("\b")

	// 計測開始
	var stats = utils.NewStats()
	stats.Begin()

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
					utils.Decorate("------------", 4)
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
						stats.AddErrorCount(utils.CorrectWrong{Correct: ansChar, Wrong: input})

						utils.MyPrintWithBlink("|", utils.LightGray)
						utils.MyPrintWithColor(input, utils.Red)
						utils.MyDeleteUntilLineEnd(false)
						utils.Routine(3, func() {
							utils.MyPrint("\b")
						})
					}
				}
				/*
				else {
					errorCount += 1
				}
				*/
			}
		}
	}

	termbox.Sync()
	stats.End()

	return stats
}
