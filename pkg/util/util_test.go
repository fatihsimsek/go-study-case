package util

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestFileExists(t *testing.T) {
	filename := "file" + strconv.Itoa(rand.Int())
	want := false

	got := FileIsExists(filename)
	if got != want {
		t.Errorf("u.FileIsExists(%s) == %v, want %v", filename, got, want)
	}
}
