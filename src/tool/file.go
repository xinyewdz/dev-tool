package tool

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func Rename(pathStr string, prefix string) {
	files, err := ioutil.ReadDir(pathStr)
	var fileList []string
	if err != nil {
		return
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		fileList = append(fileList, f.Name())
	}
	for i := 0; i < len(fileList); i++ {
		name := fileList[i]
		fullPath := filepath.Join(pathStr, name)
		os.Rename(fullPath, prefix+"_"+strconv.Itoa(i))
	}
}
