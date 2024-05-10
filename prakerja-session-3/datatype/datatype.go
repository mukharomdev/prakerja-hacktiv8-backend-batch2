package datatype

// import (
// 	"fmt"
// )

const DefaultLenArgs = 2

type Biodata struct{
	Nama string
	Alamat string
	Pekerjaan string
	Alasan string
}


type Arguments struct{
	ListArgs []string
	Arg   	 []string
	LenArg   int
	LenDb    int
}


// func (bio *Biodata)Save(bd []Biodata) []Biodata{
// 	var b []Biodata = bd

// 	return b
// }


// func (bio *Biodata)Show(){
//     ListData bdata
//     var b []Biodata = ListData

//     for i := 0 ;i <= len(b);i++{
//     	fmt.Printf("%s,%s,%s,%s",b[i].Nama,b[i].Alamat,b[i].Pekerjaan,b[i].Alasan)
//     }
// }