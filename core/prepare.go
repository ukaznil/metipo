package core

import (
	"fmt"
	"strconv"
	"time"
	"os"
	"os/user"
	"github.com/ukaznil/metipo/utils"
	"path/filepath"
)

func countDown(sec int) {
	for i := 0; i < sec; i++ {
		fmt.Print(strconv.Itoa(sec-i) + "\r")
		time.Sleep(1 * time.Second)
	}

	fmt.Println("!! MeTipo !!")
	time.Sleep(1 * time.Second)
}

const baseDirname = ".metipo"
const materialsDirname = "materials"

func getOrCreateMaterialsDirectory() os.File {
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
