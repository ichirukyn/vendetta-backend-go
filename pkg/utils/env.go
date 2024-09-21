package utils

import (
	"strings"
)

func GetEnvFilePath(path, filename string) string {
	index := strings.Index(path, "\\internal\\")
	if index != -1 {
		basePath := path[:index]
		return basePath + "\\" + filename
	} else {
		return path + "\\" + filename
	}
}
