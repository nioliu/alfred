package util

import (
	"fmt"
	"log"
	"runtime"
)

func PrintError(err error) {
	errorPrint := GetErrorPrint(err)
	fmt.Println(errorPrint)

	stack := make([]byte, 4096)
	_ = runtime.Stack(stack, true)
	log.Fatal(fmt.Sprintf("errors: %s \n %s", errorPrint, string(stack)))
}

func GetErrorPrint(err error) string {
	return fmt.Sprintf(`{"items":[{"uid":"error","type":"text","title":"Something Error","subtitle":"%s"}]}`, err.Error())
}
