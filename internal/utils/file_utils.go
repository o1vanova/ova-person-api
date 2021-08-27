package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// OpenAndCloseFileInLoop opens and closes countLoop files
func OpenAndCloseFileInLoop(countLoop int, configPath string) {
	if countLoop < 1 {
		return
	}

	for i := 1; i <= countLoop; i++ {
		log.Printf("Reading of file %d times\n", i)
		err := UpdateConfig(configPath)
		if err != nil {
			log.Fatalf("%s\n", err)
			panic(err)
			return
		}
	}
}

// UpdateConfig reads and closes file by path
func UpdateConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Panic("Error when opening file: ", err)
		return err
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

	var lines map[string] interface{}
	err = json.Unmarshal(content, &lines)
	if err != nil {
		log.Fatal("Error when parsing: ", err)
	}

	for key, value := range lines {
		// this is the place to update the config in smw
		log.Printf("%s: %v", key, value)
	}

	return err
}
