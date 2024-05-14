package main


import (
	"fmt"
	"log"
	"net/http"
	"prakerja-session-5/router"
)


func main(){
	Port := ":9001"
	router.AppRouter()
	fmt.Println("Listening on PORT 9001")

	log.Fatal(http.ListenAndServe(Port,nil))
}
