package tools

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

func ReadConfig(path string, config any) error {
	fullPath := getBaseRoot(path)

	jsonFile, err := os.ReadFile(fullPath)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonFile, config)
}

var htmlCache = make(map[string]string)

func ReadHTML(path string) string {
	if val, ok := htmlCache[path]; ok {
		return val
	}

	fullPath := getBaseRoot(path)

	fileContent, err := os.ReadFile(fullPath)
	if err != nil {
		log.Fatal("Не удалось прочитать файл", err)
	}

	htmlCache[path] = string(fileContent)

	return htmlCache[path]
}

func GetTemplate(id int) string {
	return ReadHTML("templates/" + strconv.Itoa(id) + ".html")
}
