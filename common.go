package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func readFromFile(fileName string) (string, error) {
	fi, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer fi.Close()

	fd, err := ioutil.ReadAll(fi)
	return string(fd), err
}

func saveToNewFile(fileName string, content string) error {
	if checkFileIsExist(fileName) {
		delFile(fileName)
	}

	fout, err := os.Create(fileName)
	defer fout.Close()
	if err != nil {
		return err
	}

	if err = changePermission(fileName); err != nil {
		return err
	}

	fout.WriteString(content)
	return nil
}

func checkFileIsExist(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

func changePermission(fileName string) error {
	in := bytes.NewBuffer(nil)
	cmd := exec.Command("sh")
	cmd.Stdin = in
	go func() {
		in.WriteString("chmod 755 " + fileName + "\n")
		in.WriteString("exit\n")
	}()
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func delFile(fileName string) {
	os.Remove(fileName)
	return
}

func mkFile(fileName string) error {
	_, err := os.Create(fileName)
	return err
}

func mkDir(dirPath string) error {
	return os.MkdirAll(dirPath, 0755)
}

func getFilelist(dirPath string) ([]string, error) {
	var result []string
	err := filepath.Walk(dirPath, func(dirPath string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		result = append(result, dirPath)
		return nil
	})
	return result, err
}

func getKeyFromFilePath(filePath string) string {
	lastIndex := strings.LastIndex(filePath, "/")
	return filePath[lastIndex+1:]
}
