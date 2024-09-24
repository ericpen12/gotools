package csv

import "testing"

func TestBuffer_autoWidth(t *testing.T) {
	data := [][]string{
		{
			"name",
			"年龄",
			"身高（单位厘米）",
			"性别",
		},
		{
			"Tom",
			"21",
			"178",
			"1",
		},
	}
	c := NewReaderStringSlice(data)
	c.SetColumnWidth(10)
	c.ToCsv("./person.csv")
}
