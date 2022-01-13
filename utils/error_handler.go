package utils

import (
	"fmt"
	"runtime"
)

func HandleException(err error) {
	if err != nil {
		fmt.Println(err)
		runtime.Goexit()
	}
}
