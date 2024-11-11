package main_test

import (
	wc "github.com/Adedunmol/go-wc"
	"reflect"
	"strings"
	"testing"
	"testing/fstest"
)

func TestWordCount(t *testing.T) {
	width := strings.Repeat(" ", 8)

	t.Run("test counting", func(t *testing.T) {
		fileSystem := fstest.MapFS{
			"test.txt": {Data: []byte(string("Some random text"))},
		}
		files := []string{"test.txt"}

		counts := wc.NewCountFromFile(fileSystem, files)
		wanted := wc.Count{FileName: "test.txt", Lines: 1, Words: 3, Characters: 16}

		assertCountEqual(t, counts[0], wanted)
	})

	t.Run("test line format output", func(t *testing.T) {
		fileSystem := fstest.MapFS{
			"test.txt": {Data: []byte(string("Some random text"))},
		}
		files := []string{"test.txt"}

		result := "1 test.txt"
		wanted := strings.Join([]string{width, result}, "")
		count := wc.NewCountFromFile(fileSystem, files)

		got := count[0].Format(wc.Options{Line: true})
		assertFormatOutputEqual(t, got, wanted)
	})

	t.Run("test word format output", func(t *testing.T) {
		fileSystem := fstest.MapFS{
			"test.txt": {Data: []byte(string("Some random text"))},
		}
		files := []string{"test.txt"}

		result := "3 test.txt"
		wanted := strings.Join([]string{width, result}, "")
		count := wc.NewCountFromFile(fileSystem, files)

		got := count[0].Format(wc.Options{Word: true})
		assertFormatOutputEqual(t, got, wanted)
	})

	t.Run("test character format output", func(t *testing.T) {
		fileSystem := fstest.MapFS{
			"test.txt": {Data: []byte(string("Some random text"))},
		}
		files := []string{"test.txt"}

		result := "16 test.txt"
		wanted := strings.Join([]string{width, result}, "")
		count := wc.NewCountFromFile(fileSystem, files)

		got := count[0].Format(wc.Options{Character: true})
		assertFormatOutputEqual(t, got, wanted)
	})

	t.Run("test no option format output", func(t *testing.T) {
		fileSystem := fstest.MapFS{
			"test.txt": {Data: []byte(string("Some random text"))},
		}
		files := []string{"test.txt"}

		result := "1 3 16 test.txt"
		wanted := strings.Join([]string{width, result}, "")
		count := wc.NewCountFromFile(fileSystem, files)

		got := count[0].Format(wc.Options{})
		assertFormatOutputEqual(t, got, wanted)
	})
}

func assertCountEqual(t *testing.T, got, want wc.Count) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v, got: %v\n", want, got)
	}
}

func assertFormatOutputEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("wanted: %v, got: %v\n", want, got)
	}
}
