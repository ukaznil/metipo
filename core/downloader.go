package core

import (
	"net/http"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ukaznil/metipo/utils"
	"os"
	"io/ioutil"
	"path/filepath"
)

type fileInfo struct {
	Name        string `json:"name"`
	DownloadUrl string `json:"download_url"`
}

func getDownloadList() []fileInfo {
	const githubFileListUrl = "https://api.github.com/repos/ukaznil/metipo/contents/materials"
	var res, err = http.Get(githubFileListUrl)
	utils.Perror(err)
	defer res.Body.Close()

	var data []fileInfo
	var buf = new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	var resByte = buf.Bytes()
	var err2 = json.Unmarshal(resByte, &data)
	utils.Perror(err2)

	return data
}

func downloadFile(url string, dest string) string {
	var res, err1 = http.Get(url)
	if err1 != nil {
		fmt.Println(err1)
		return "failed"
	}

	var body, err2 = ioutil.ReadAll(res.Body)
	if err2 != nil {
		fmt.Println(err2)
		return "failed"
	}

	if _, err := os.Stat(dest); err == nil {
		return "skipped"
	} else {
		var file, err3 = os.Create(dest)
		defer file.Close()
		if err3 != nil {
			fmt.Println(err3)
			return "failed"
		}

		file.Write(body)

		return "success"
	}
}

func DownloadFromGitHub(dest os.File) {
	fmt.Println("[MeTipo] now start downloading...")
	utils.HLine()
	var list = getDownloadList()
	for _, fileInfo := range list {
		var success = downloadFile(
			fileInfo.DownloadUrl,
			filepath.Join(dest.Name(), fileInfo.Name))
		fmt.Println(fileInfo.Name + ": " + success)
	}
	utils.HLine()
	fmt.Println("[MeTipo] download done.")
}
