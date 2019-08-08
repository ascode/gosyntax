package main

import (
	"fmt"
	"os"
)

func deleteFile(fileName string) error {
	ifexist, err := pathExists(fileName)
	if err != nil {
		return err
	}
	if ifexist {
		fmt.Println("文件存在", fileName)
		if err := os.Remove(fileName); err != nil {
			return err
		} else {
			fmt.Println("文件已经删除", fileName)
			return nil
		}
	} else {
		return nil
	}
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	deleteFile("/Users/alston/SourceCode/universe/gosyntax/aaa.txt")
}