package pkg

import (
	"fmt"
	"time"
)

// PrintRunCost 打印运行时间
var PrintRunCost = func() func() {
	timeNow := time.Now()
	return func() {
		fmt.Printf("耗时: %0.3f\n", time.Since(timeNow).Seconds())
	}
}()
