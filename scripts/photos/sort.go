package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var rootDir = flag.String("root", "", "Root directory of photos")
var targetDir = flag.String("target", "", "Target directory for sorting")

func main() {
	flag.Parse()

	if *rootDir == "" || *targetDir == "" {
		log.Fatal("Please provide the root and target directories as arguments")
	}

	err := filepath.Walk(*rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			timestamp := info.ModTime()
			dateStr := timestamp.Format("2006/01/02")
			targetPath := filepath.Join(*targetDir, dateStr)
			err = os.MkdirAll(targetPath, 0755)
			if err != nil {
				return err
			}
			err = os.Rename(path, filepath.Join(targetPath, filepath.Base(path)))
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Photos sorted successfully!")
}
