package core

import (
	"fmt"
	"strconv"
	"time"
	"os"
)

func countDown(sec int) {
	for i := 0; i < sec; i++ {
		fmt.Print(strconv.Itoa(sec-i) + "\r")
		time.Sleep(1 * time.Second)
	}

	fmt.Println("!! MeTipo !!")
	time.Sleep(1 * time.Second)
}

func getOrCreateMaterialsDirectory() os.File {
	var dir = getMaterialsDirpath()
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
