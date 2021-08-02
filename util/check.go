package util

import "fmt"

// Check 处理错误
func Check(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
