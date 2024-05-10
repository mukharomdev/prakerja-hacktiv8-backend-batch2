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
	 case reflect.Interface:
	 	fmt.Println(d.Type(),d.Interface())
	 default:
	 	fmt.Println("Bukan String dan Integer")
	 }
}