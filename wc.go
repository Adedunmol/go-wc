package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	line      = flag.Bool("l", false, "Show only line count")
	word      = flag.Bool("w", false, "Show only word count")
	character = flag.Bool("c", false, "Show only character count")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: wc [options] [files]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("wc: ")

	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		usage()
	}

	files := stat(args)

	if _, err := run(files, Options{}); err != nil {
		log.Println(err)
	}
}

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

func run(files []string, options Options) (string, error) {
	var output []string

	counts := NewCountFromFile(os.DirFS("."), files)

	for _, count := range counts {
		result := count.Format(options)
		output = append(output, result)
	}

	if len(files) > 1 {
		total := Total(counts)
		output = append(output, total)
	}

	return strings.Join(output, "\n"), nil
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

	result := strings.Join(output, " ")

	width := strings.Repeat(" ", 8)

	return strings.Join([]string{width, result}, "")
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

func stat(files []string) []string {
	filteredFiles := make([]string, len(files))
	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
		}

		if !info.IsDir() {
			filteredFiles = append(filteredFiles, file)
		}
	}

	return filteredFiles
}

func Total(counts []Count) string {
	var lines uint
	var characters uint
	var words uint

	for _, count := range counts {
		lines += count.Lines
		characters += count.Characters
		words += count.Words
	}

	var output []string

	output = append(output, strconv.Itoa(int(lines)))
	output = append(output, strconv.Itoa(int(words)))
	output = append(output, strconv.Itoa(int(characters)))
	output = append(output, "total")

	result := strings.Join(output, " ")

	width := strings.Repeat(" ", 8)

	return strings.Join([]string{width, result}, "")
}
