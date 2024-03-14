package file

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type Info struct {
	Path       string
	Line       int
	Preference string
}

type ReadLineFunc func(content string, info Info) bool

func ReadLine(filepath string, fn ReadLineFunc) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	info := Info{
		Path: filepath,
	}
	return readLine(bufio.NewReader(f), info, fn)
}

func readLine(buf *bufio.Reader, info Info, fn ReadLineFunc) error {
	for {
		lineBytes, _, err := buf.ReadLine()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		info.Line++
		info.Preference = fmt.Sprintf("%s:%d", info.Path, info.Line)
		if !fn(string(lineBytes), info) {
			return nil
		}
	}
}

func ReadLineByWalkDir(dir string, fn ReadLineFunc, ignore func(path string) bool) error {
	return filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if ignore != nil && ignore(path) {
			return nil
		}
		return ReadLine(path, fn)
	})
}
