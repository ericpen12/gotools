package file

import (
	"bufio"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type FileInfo struct {
	Path string
	Line int
}

type ReadLineFunc func(content string, info FileInfo) bool

func ReadLine(filepath string, fn ReadLineFunc) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	info := FileInfo{
		Path: filepath,
	}
	return readLine(bufio.NewReader(f), info, fn)
}

func readLine(buf *bufio.Reader, info FileInfo, fn ReadLineFunc) error {
	for {
		lineBytes, _, err := buf.ReadLine()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		info.Line++
		if !fn(string(lineBytes), info) {
			return nil
		}
	}
}

func ReadLineByWalkDir(dir string, fn ReadLineFunc, ignorePath ...string) error {
	return filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		for _, v := range ignorePath {
			if path == v {
				return nil
			}
		}
		return ReadLine(path, fn)
	})
}
