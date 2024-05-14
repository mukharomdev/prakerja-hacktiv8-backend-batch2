package router


import (
	"prakerja-session-5/handler"
	"net/http"
)


func AppRouter(){


	http.HandleFunc("/products",handler.HandlerProducts)



}