package main

import (
	"testing"
	"io/fs"

	"github.com/stretchr/testify/assert"
)

func (i ioHandler) Open(name string) (File, error) {

}

func TestReadFile(t *testing.T) {
	ioHandler := NewIOHandler(fs fs.FS)
	output, err := ioHandler.readFile("testfile.txt")
}
