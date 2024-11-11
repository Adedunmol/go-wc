package main

import (
	"bufio"
	"flag"
	"io"
	"io/fs"
	"strconv"
	"strings"
)

var (
	line      = flag.Bool("l", false, "Show only line count")
	word      = flag.Bool("w", false, "Show only word count")
	character = flag.Bool("c", false, "Show only character count")
)

func main() {}

type Count struct {
	FileName   string
	Lines      uint
	Words      uint
	Characters uint
}

type Options struct {
	Line      bool
	Word      bool
	Character bool
}

func (c *Count) Format(options Options) string {
	var output []string

	if options.Line {
		output = append(output, strconv.Itoa(int(c.Lines)))
	}

	if options.Word {
		output = append(output, strconv.Itoa(int(c.Words)))
	}

	if options.Character {
		output = append(output, strconv.Itoa(int(c.Characters)))
	}

	if len(output) == 0 {
		output = append(output, strconv.Itoa(int(c.Lines)))
		output = append(output, strconv.Itoa(int(c.Words)))
		output = append(output, strconv.Itoa(int(c.Characters)))
	}

	output = append(output, c.FileName)

	return strings.Join(output, " ")
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
	lines := splitLines(string(fileData))

	count.Lines = uint(len(lines))

	count.Words = uint(splitWords(lines))

	count.Characters = uint(splitCharacters(lines))

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

func splitWords(lines []string) uint {
	var words uint
	for _, line := range lines {
		splitWords := strings.Split(line, " ")
		words += uint(len(splitWords))
	}

	return words
}

func splitCharacters(lines []string) uint {
	var chars uint
	for _, line := range lines {
		splitChars := strings.Split(line, "")
		chars += uint(len(splitChars))
	}
	return chars
}
