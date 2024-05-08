package handler

import (
	// "fmt"
	"os"
	"strconv"
)



func getArgs() any {
	args := os.Args
	arg1 := args[1]
	var arg int64 = 0
	num,err := strconv.ParseInt(arg1,10,8)

	if err == nil {
		arg = num
	}



  return arg
}


func Get() any {
	var arg = getArgs()
	return arg
}