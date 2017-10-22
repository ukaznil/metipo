package core

import (
	"os/user"
	"github.com/ukaznil/metipo/utils"
	"os"
	"path/filepath"
)

func getHomeDirpath() string {
	usr, err := user.Current()
	utils.Perror(err)

	dir, err := os.Open(usr.HomeDir)
	defer dir.Close()
	utils.Perror(err)

	return dir.Name()
}

func getMeTipoDirpath() string {
	const metipoDirname = ".metipo"
	return filepath.Join(getHomeDirpath(), metipoDirname)
}

func getMaterialsDirpath() string {
	const materialsDirname = "materials"
	return filepath.Join(getMeTipoDirpath(), materialsDirname)
}
