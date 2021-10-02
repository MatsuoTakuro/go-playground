//go:build !windows && !plan9
// +build !windows,!plan9

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {

	crPath, _ := os.Getwd()
	filesPath := fmt.Sprintf("%v/files/", crPath)
	err := os.Chdir(filesPath)
	if err != nil {
		fmt.Printf("unable to change tmpDir path : %v\n", err)
	}

	var storePath string

	err = filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && path != "." {
			storePath = path
			fmt.Println(storePath)
		}
		if !info.IsDir() {
			filePath := path
			fmt.Println(filePath)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", filesPath, err)
		return
	}
}
