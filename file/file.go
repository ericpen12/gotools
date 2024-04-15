package file

import (
	"bufio"
	"io"
	"os"
)

func ReadLine(filepath string, fn func(content string, line int) bool) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()
	return readLine(bufio.NewReader(f), fn)
}

func readLine(buf *bufio.Reader, fn func(content string, line int) bool) error {
	var line int
	for {
		lineBytes, _, err := buf.ReadLine()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		line++
		if !fn(string(lineBytes), line) {
			return nil
		}
	}
}
