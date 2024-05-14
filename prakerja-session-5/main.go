package main


import (
	//"io"
	"log"
	"net/http"
	"prakerja-session-5/router"
)


func main(){
	Port := ":9001"
	router.AppRouter()
	log.Fatal(http.ListenAndServe(Port,nil))
}
