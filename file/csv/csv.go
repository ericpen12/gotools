package csv

import (
	"encoding/csv"
	"github.com/ericpen12/gotools/log"
	"io"
	"os"
)

type Csv interface {
	TitleIndex(title string) int
	DeleteColumn(index int)
	AddColumn(title string, index int)
	Range(fn func(record []string) bool)
	ToCsv(filename string) error
	ExchangeColumn(i, j int)
	MoveColumn(current, target int)
	Add(data []string)
}

type Buffer struct {
	data [][]string
}

func (r *Buffer) TitleIndex(title string) int {
	if len(r.data) == 0 {
		return -1
	}
	for i, t := range r.data[0] {
		if t == title {
			return i
		}
	}
	return -1
}

func (r *Buffer) DeleteColumn(index int) {
	if len(r.data)-1 < index {
		return
	}
	for i := range r.data {
		r.data[i] = append(r.data[i][:index], r.data[i][index+1:]...)
	}
}

func (r *Buffer) AddColumn(title string, index int) {
	content := title
	for i := range r.data {
		if i > 0 {
			content = ""
		}
		r.data[i] = append(r.data[i][:index], append([]string{content}, r.data[i][index:]...)...)
	}
}

func (r *Buffer) Range(fn func(record []string) bool) {
	for _, record := range r.data {
		if !fn(record) {
			break
		}
	}
}

func (r *Buffer) Add(data []string) {
	r.data = append(r.data, data)
}

func NewCsvFile(filename string) Csv {
	file, err := os.Open(filename)
	if err != nil {
		log.Info("无法打开 Operator 文件:", err)
	}
	r := csv.NewReader(file)
	var list [][]string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		list = append(list, record)
	}
	return &Buffer{data: list}
}

func (r *Buffer) ExchangeColumn(i, j int) {
	if len(r.data)-1 < i || len(r.data)-1 < j {
		return
	}
	for k := range r.data {
		r.data[k][i], r.data[k][j] = r.data[k][j], r.data[k][i]
	}
}

func (r *Buffer) MoveColumn(current, target int) {
	if len(r.data)-1 < current || len(r.data)-1 < target {
		return
	}
	for k, v := range r.data {
		tempValue := v[current]
		temp := append(v[:current], v[current+1:]...)
		temp = append(temp[:target], append([]string{tempValue}, temp[target:]...)...)
		r.data[k] = temp
	}
}

func NewReaderStringSlice(content [][]string) Csv {
	return &Buffer{data: content}
}

func (r *Buffer) ToCsv(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	w := csv.NewWriter(file)
	if err := w.WriteAll(r.data); err != nil {
		return err
	}
	w.Flush()
	return nil
}
