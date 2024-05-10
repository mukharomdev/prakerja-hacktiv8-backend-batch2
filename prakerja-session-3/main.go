package main


import (
	"fmt"
	"os"
	"prakerja-session-3/handler"
)

func main(){
	defer func(){
		if r:= recover();r != nil{
			fmt.Println("Harus pakai argument,gunakan perintah,contoh : go run . 1")
		}else if len(os.Args) == 0{
			panic("Error")
		}
	}()


	handler.ShowDatabase()

}