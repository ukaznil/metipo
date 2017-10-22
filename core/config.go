package core

import (
	"gopkg.in/yaml.v2"
	"path/filepath"
	"io/ioutil"
	"github.com/ukaznil/metipo/utils"
	"os"
	"bufio"
	"fmt"
)

type Language string

const (
	Japanese = "ja"
	English  = "en"
)

type config struct {
	Langs []Language `yaml:"langs"`
}

func getConfigFilepath() string {
	const configFilename = ".mtp-config.yml"
	return filepath.Join(getMeTipoDirpath(), configFilename)
}

func InitConfig() {
	// todo
}

func changeConfig() {
	// todo
}

func resetConfig() {
	var cfg = config{
		Langs: []Language{Japanese, English},
	}
	buf, err := yaml.Marshal(cfg)
	utils.Perror(err)

	// output
	configFile, err := os.OpenFile(getConfigFilepath(), os.O_WRONLY|os.O_CREATE, 0600)
	defer configFile.Close()
	var writer = bufio.NewWriter(configFile)
	fmt.Println(buf)
	writer.Write(buf)
	writer.Flush()
}

func getCurrentLanguages() []Language {
	var configFilepath = getConfigFilepath()
	if _, err := os.Stat(configFilepath); err != nil {
		resetConfig()
	}

	buf, err := ioutil.ReadFile(configFilepath)
	utils.Perror(err)

	var cfg config
	err = yaml.Unmarshal(buf, &cfg)
	utils.Perror(err)

	return cfg.Langs
}
