package renpyGraphviz

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// FileHandler opens all renpy files and transform them into a string
func fileHandler(rootPath string) []string {

	files, err := walkMatch(rootPath, "*.rpy")
	if err != nil {
		log.Fatalf("failed to find root folder: %s", err)
	}

	var fileTextLines []string

	for _, file := range files {
		readFile, err := os.Open(file)

		if err != nil {
			log.Fatalf("failed to open file: %s", err)
		}

		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)

		for fileScanner.Scan() {
			fileTextLines = append(fileTextLines, fileScanner.Text())
		}
		fileTextLines = append(fileTextLines, "# renpy-graphviz: BREAK")

		readFile.Close()
	}

	return fileTextLines
}

func walkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func writeFile(filename, content string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(content)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
