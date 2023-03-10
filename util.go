package util

import (
	"fmt"
	"log"
)

func PrintError(err error) {
	errorPrint := GetErrorPrint(err)
	fmt.Println(errorPrint)
	log.Fatal(errorPrint)
}

func GetErrorPrint(err error) string {
	return fmt.Sprintf(`{"items":[{"uid":"error","type":"text","title":"Something Error","subtitle":"%s"}]}`, err.Error())
}
