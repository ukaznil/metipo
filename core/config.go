package core

import (
	"gopkg.in/yaml.v2"
	"path/filepath"
	"io/ioutil"
	"github.com/ukaznil/metipo/utils"
	"os"
	"bufio"
)

type settings struct {
	ModelColor utils.Color
	YourColor  utils.Color
}

func getConfigFilepath() string {
	const configFilename = ".mtp-config.yml"
	return filepath.Join(getMeTipoDirpath(), configFilename)
}

func InitConfig() {
	resetConfig()
}

func changeConfig() {
	buf, err := ioutil.ReadFile(getConfigFilepath())
	utils.Perror(err)

	var stgs settings
	err = yaml.Unmarshal(buf, &stgs)
	utils.Perror(err)
}

func resetConfig() {
	var stgs = settings{
		ModelColor: utils.LightGray,
		YourColor:  utils.White,
	}
	buf, err := yaml.Marshal(&stgs)
	utils.Perror(err)

	// output
	configFile, err := os.OpenFile(getConfigFilepath(), os.O_WRONLY|os.O_CREATE, 0600)
	defer configFile.Close()
	var writer = bufio.NewWriter(configFile)
	writer.Write(buf)
	writer.Flush()
}
