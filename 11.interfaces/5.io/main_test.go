package main

import (
	"fmt"
	"io"
	"io/fs"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockedFileSystem struct{}

func (m mockedFileSystem) Open(name string) (fs.File, error) {
	return &mockedFile{}, nil
}

type mockedFile struct{}

func (m *mockedFile) Stat() (fs.FileInfo, error) {
	fmt.Println("Stat called on mocked file")
	return nil, nil
}

func (m *mockedFile) Read(b []byte) (int, error) {
	fmt.Println("Read called on mocked file", string(b))
	fakeFileContent := "This is a test file."
	for i := 0; i < len(fakeFileContent) && i < len(b); i++ {
		b[i] = fakeFileContent[i]
	}
	return len(fakeFileContent), io.EOF
}

func (m *mockedFile) Close() error {
	return nil
}

type mockedFileInfo struct{}

func (m *mockedFileInfo) Name() string {
	return "mocked_file.txt"
}

func (m *mockedFileInfo) Size() int64 {
	return 1234
}

func (m *mockedFileInfo) Mode() fs.FileMode {
	return 0o644
}

func (m *mockedFileInfo) ModTime() time.Time {
	return time.Now()
}

func (m *mockedFileInfo) IsDir() bool {
	return false
}

func (m *mockedFileInfo) Sys() any {
	return nil
}

func TestReadFile(t *testing.T) {
	mockedFileSystem := mockedFileSystem{}
	ioHandler := NewIOHandler(mockedFileSystem)
	output, err := ioHandler.readFile("testfile.txt")
	assert.NoError(t, err, "should not return an error when reading a file")
	assert.Equal(t, "This is a test file.", output, "should return the correct file content")
}

type mockedFileSystemThatFailsToOpen struct{}

func (m mockedFileSystemThatFailsToOpen) Open(name string) (fs.File, error) {
	return nil, fmt.Errorf("failed to open file %s", name)
}

func TestReadFileOpenError(t *testing.T) {
	mockedFileSystemThatFailsToOpen := mockedFileSystemThatFailsToOpen{}
	ioHandler := NewIOHandler(mockedFileSystemThatFailsToOpen)
	_, err := ioHandler.readFile("testfile.txt")
	fmt.Println("ERROR:", err)
	assert.Error(t, err, "should return an error when failing to open a file")
}

type mockedFileSystemFailedToRead struct{}

func (m mockedFileSystemFailedToRead) Open(name string) (fs.File, error) {
	return &mockedFileFailedToRead{}, nil
}

type mockedFileFailedToRead struct{}

func (m *mockedFileFailedToRead) Stat() (fs.FileInfo, error) {
	fmt.Println("Stat called on mocked file")
	return nil, nil
}

func (m *mockedFileFailedToRead) Read(b []byte) (int, error) {
	fmt.Println("Read called on mocked file", string(b))
	// Simulate a read error
	return 0, fmt.Errorf("failed to read from mocked file")
}

func (m *mockedFileFailedToRead) Close() error {
	return nil
}

func TestReadFileReadError(t *testing.T) {
	mockedFileSystemFailedToRead := mockedFileSystemFailedToRead{}
	ioHandler := NewIOHandler(mockedFileSystemFailedToRead)
	_, err := ioHandler.readFile("testfile.txt")
	fmt.Println("ERROR:", err)
	assert.Error(t, err, "should return an error when failing to read a file")
}
