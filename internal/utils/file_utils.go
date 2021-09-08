package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

// OpenAndCloseFileInLoop opens and closes countLoop files
func OpenAndCloseFileInLoop(countLoop int, configPath string) {
	if countLoop < 1 {
		return
	}

	for i := 1; i <= countLoop; i++ {
		log.Printf("Reading of file %d times\n", i)
		_, err := GetConfig(configPath)
		if err != nil {
			log.Fatalf("%s\n", err)
			panic(err)
		}
	}
}

// GetConfig reads and closes file by path
func GetConfig(path string) (map[string]interface{}, error) {
	file, err := os.Open(path)
	var lines map[string]interface{}
	if err != nil {
		log.Panic("Error when opening file: ", err)
		return lines, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(content, &lines)
	if err != nil {
		log.Fatal("Error when parsing: ", err)
	}

	for key, value := range lines {
		// this is the place to update the config in smw
		log.Printf("%s: %v", key, value)
	}

	return lines, err
}

// GetConfigPath get path to config
func GetConfigPath() string {
	_, currentFile, _, _ := runtime.Caller(0)
	rootDir := filepath.Join(filepath.Dir(currentFile), "../..")
	configDir := path.Join(rootDir, "configs")
	configPath := filepath.Join(configDir, "config.json")

	return configPath
}

func Format(s string, v interface{}) string {
	t, b := new(template.Template), new(strings.Builder)
	err := template.Must(t.Parse(s)).Execute(b, v)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return b.String()
}
