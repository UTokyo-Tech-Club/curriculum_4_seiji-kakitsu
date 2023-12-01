package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	dir := os.Args[1]
	err := NewMyError(cd(dir), "fail: パス関連のエラーが発生しました: %s, 指定可能なサブディレクトリ: %s\n", ReadDir())

	if err.original != nil {
		var pathError *os.PathError
		if errors.As(err.original, &pathError) {
			if ReadDir() != nil {
				dirs := strings.Join(ReadDir(), ", ")
				fmt.Printf(err.format, err.original.Error(), dirs)
			} else {
				dirs := "なし"
				fmt.Printf(err.format, err.original.Error(), dirs)
			}
		} else {
			fmt.Printf("fail: エラーが発生しました： %s\n", err.original.Error())
		}
		return
	} else {
		fmt.Println("success!")
	}
}

func cd(dir string) error {
	return os.Chdir(dir)
}

func ReadDir() []string {
	files, _ := os.ReadDir("./")
	var a []string
	for _, file := range files {
		if file.IsDir() {
			a = append(a, file.Name())
		}
	}
	return a
}

type MyError struct {
	original error
	format   string
	values   []any
}

func NewMyError(original error, format string, values ...any) *MyError {
	return &MyError{
		original: original,
		format:   format,
		values:   values,
	}
}
