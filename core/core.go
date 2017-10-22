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
	"time"
)

func selectMaterialName() string {
	var materialsDir = getOrCreateMaterialsDirectory()
	var items, err = ioutil.ReadDir(materialsDir.Name())
	utils.Perror(err)

	var filepaths []string
	for _, item := range items {
		if !item.IsDir() {
			var fp = filepath.Join(materialsDir.Name(), item.Name())
			filepaths = append(filepaths, fp)
		}
	}

	rand.Seed(time.Now().UnixNano())
	return filepaths[rand.Intn(len(filepaths))]
}

func Exercise() {
	//comm.DownloadFromGitHub(materialsDir)
	rand.Seed(time.Now().UnixNano())
	var langs = getCurrentLanguages()
	var lang = langs[rand.Intn(len(langs))]
	fmt.Println(langs, lang)

	var filepath = DownloadWikipediaArticle(lang)
	//var material, err = os.Open(selectMaterialName())
	var material, err = os.Open(filepath)
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

	var width, _ = termbox.Size()

	fmt.Println(msg)
	//utils.HLine()
	utils.Decorate("-*--*--*--*-", 4)

	var scanner = bufio.NewScanner(&material)
	utils.Perror(scanner.Err())
	var lines = make([]string, 0)
	for scanner.Scan() {
		var line = scanner.Text()
		for _, limitedLine := range utils.SeparateByLength(line, width-1) {
			//lines = append(lines, line)
			lines = append(lines, limitedLine)
		}
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
		//log.Println("[" + string(ev.Ch) + "]")

		switch ev.Key {
		case termbox.KeyBackspace,
			termbox.KeyBackspace2,
			termbox.KeyDelete,
			termbox.KeyArrowLeft,
			termbox.KeyArrowRight:

		case termbox.KeyEsc:
			break loop

			/*
		case termbox.KeySpace:
			fmt.Println("<space>")
			utils.MyPrint(" ")
			utils.MyPrintWithBlink("|", utils.LightGray)
			utils.MyPrint("\b")
			*/

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
			var input string
			switch ev.Key {
			case termbox.KeySpace:
				input = " "
			default:
				input = string(ev.Ch)
			}

			if charIndex < utf8.RuneCountInString(line) {
				var ansChar = string([]rune(line)[charIndex])
				//log.Println("[" + input + "," + ansChar + "]")

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
					utils.Routine(2, func() {
						utils.MyPrint("\b")
					})
				}
			}
		}

		termbox.Flush()
	}

	termbox.Sync()
	stats.End()

	return stats
}
