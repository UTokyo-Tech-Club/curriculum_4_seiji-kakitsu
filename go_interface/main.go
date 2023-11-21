package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dir := os.Args[1]

	if err := cd(dir); err != nil {
		files, _ := ioutil.ReadDir("./")
		var a []string
		for _, file := range files {
			if file.IsDir() {
				a = append(a, file.Name())
			}
		}
		dirs := strings.Join(a, ", ")
		if a != nil {
			fmt.Printf("fail: %s, 指定可能なサブディレクトリ: %s\n", err.Error(), dirs)
		} else {
			fmt.Printf("fail: %s, 指定可能なサブディレクトリ: なし\n", err.Error())
		}
	} else {
		fmt.Println("success!")
	}
}

func cd(dir string) error {
	return os.Chdir(dir)
}
