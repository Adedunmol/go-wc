package main_test

import (
	wc "github.com/Adedunmol/go-wc"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestWordCount(t *testing.T) {
	fileSystem := fstest.MapFS{
		"test.txt": {Data: []byte(string("Some random text"))},
	}
	files := []string{"test.txt"}

	counts := wc.NewCountFromFile(fileSystem, files)
	wanted := wc.Count{FileName: "test.txt", Lines: 1, Words: 3, Characters: 16}

	assertCountEqual(t, counts[0], wanted)
}

func assertCountEqual(t *testing.T, got, want wc.Count) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v, got: %v\n", want, got)
	}
}
