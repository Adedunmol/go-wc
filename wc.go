package main

import (
	"bufio"
	"flag"
	"io"
	"io/fs"
	"strings"
)

var (
	line      = flag.Bool("l", false, "Show only line count")
	word      = flag.Bool("w", false, "Show only word count")
	character = flag.Bool("c", false, "Show only character count")
)

func main() {}

type Count struct {
	FileName string
}

func NewCountFromFile(fileSystem fs.FS, files []string) []Count {
	var counts []Count

	for _, file := range files {
		count := Count{FileName: file}
		_ = getFile(fileSystem, file, &count)
		counts = append(counts, count)
	}
	return counts
}

func getFile(fileSystem fs.FS, filePath string, count *Count) error {
	//path, _ := filepath.Abs(filePath)
	file, _ := fileSystem.Open(filePath)

	defer file.Close()
	fileData, _ := io.ReadAll(file)
	_ = splitLines(string(fileData))

	return nil
}

func splitLines(line string) []string {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(line))

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
