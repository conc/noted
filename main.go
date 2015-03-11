package main

import (
	"errors"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	cmd := flag.Arg(0)
	key := flag.Arg(1)
	keyArg := flag.Arg(2)

	var err error
	var fileContent string

	switch cmd {
	case "add":
		err = addKey(key, keyArg)
	case "del":
		err = delKey(key)
	case "get":
		fileContent, err = getKey(key)
	case "append":
		err = appendKey(key, keyArg)
	case "ls":
		err = showKeys()
	default:
		showHelp()
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	showResult(cmd, fileContent)

	return
}

func showResult(cmd string, fileContent string) {
	if cmd == "get" {
		fmt.Println(fileContent)
	}
	return
}

func showHelp() {
	fmt.Println(helpStr)
}

func addKey(key string, addContent string) error {
	if checkFileIsExist(dataPath + key) {
		return errors.New("key is already exist!")
	}
	return saveToNewFile(dataPath+key, addContent)
}

func delKey(key string) error {
	delFile(dataPath + key)
	return nil
}

func getKey(key string) (string, error) {
	if !checkFileIsExist(dataPath + key) {
		return "", errors.New("key is not exist!")
	}
	return readFromFile(dataPath + key)
}

func appendKey(key string, addContent string) error {
	oriContent, err := getKey(key)
	if err != nil {
		return err
	}
	newContent := oriContent + addContent
	err = delKey(key)
	if err != nil {
		return err
	}
	return addKey(key, newContent)

}

func showKeys() error {
	keys, err := getFilelist(dataPath)
	if err != nil {
		return err
	}
	for i := 0; i < len(keys); i++ {
		key := getKeyFromFilePath(keys[i])
		fmt.Print(key + " ")
	}
	fmt.Println()
	return nil
}
