package reflection


import (
	"reflect"
	"fmt"
)


func Ref(c any){
	 d := reflect.ValueOf(c)
	 switch d.Kind(){
	 case reflect.Int:
	 	fmt.Println(d.Type(),d.Int())
	 case reflect.String:
	 	fmt.Println(d.Type(),d.String())
	 default:
	 	fmt.Println("Bukan String dan Integer")
	 }
}