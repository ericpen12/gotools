package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// PrintRunCost 打印运行时间
var PrintRunCost = func() func() {
	timeNow := time.Now()
	return func() {
		fmt.Printf("耗时: %0.3fs\n", time.Since(timeNow).Seconds())
	}
}()

// ZeroTime 获取指定日期的零点
func ZeroTime(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// ZeroTimeToday 获取今天的零点
func ZeroTimeToday() time.Time {
	return ZeroTime(time.Now())
}

// GetCurrentAppName 获取当前可执行文件的名称
func GetCurrentAppName() string {
	programPath := os.Args[0]
	return filepath.Base(programPath)
}
