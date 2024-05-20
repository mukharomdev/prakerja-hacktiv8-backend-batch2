package main



import (
	"fmt"
)

// Interface
type Interface interface{
	Interface() any
}

//struct
type Struct1 struct{

}

// method struct1
func(s1 *Struct1)Interface()any{
    var s = "dari struct1 =>"
    var u = "bisa dipanggil asal punya metod sama dengan Interface"
 return s+u
}

// Fungsi membuat struct1
func NewS1(s1 Interface)*Struct1{
	return &Struct1{

	}
}

// Struct dua
type Struct2 struct{

}
//method struct2
func(s2 Struct2)Interface()any{
	var s = "dari struct2 =>"
	var u = "bisa dipanggil asal punya metod sama dengan Interface"
	return s + u
}

// Fungsi membuat struct1
func NewS2(s2 *Interface)*Struct2{
	return &Struct2{

	}
}

/* fungsi yang berisi parameter Interface dan
bisa di isi oleh struct yang mempunyai method
sama dengan yang difenisikan di Interface
*/
func ShowInterface1(K Interface){
	fmt.Println(K.Interface())
}

func ShowInterface2(K Interface){
	fmt.Println(K.Interface())
}



func main(){
 var interfacee Interface
 var struct1 Struct1
 var struct11 Struct1
 var struct2 Struct2

 fmt.Printf("interface:%t\n",interfacee)

 fmt.Printf("struct 1  :%t\n",struct1)
 fmt.Printf("struct 2  :%t\n",struct2)


 fmt.Printf("panggil struct 1 dengan fungsi dengan pointer:%v\n",NewS1(&struct1))
 fmt.Printf("panggil struct 1 dengan fungsi tanpa pointer:%v\n",NewS2(&interfacee))

 fmt.Printf("membanding struct1 <=> struct2 hasilnya :%v\n",struct11 == struct1)

 st1 := &Struct1{}

 ShowInterface1(st1)

 fmt.Println("=============")

 st2 := Struct2{}

 ShowInterface2(st2)

 fmt.Println("=============")








}