package utils

import (
	"fmt"
	"runtime"
)

func MarkErrLoc() {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("Error occurred at %s:%d\n", file, line)
	} else {
		fmt.Println("Failed to get error location")
	}
}
