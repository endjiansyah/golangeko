package helper

import (
	"fmt"
	"runtime"
)

func PanicIfError(err error) {
	if err != nil {
		fmt.Println(err)
		// panic(err)
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("Error at %s:%d\n", file, line)
		fmt.Println(err.Error())
	}
}
