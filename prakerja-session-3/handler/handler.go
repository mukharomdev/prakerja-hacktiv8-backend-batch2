package handler

import (
	"fmt"
	"os"
	"strconv"
	//"errors"
  . "prakerja-session-3/database"
  . "prakerja-session-3/datatype"
  //. "prakerja-session-3/reflection"

)




func getArguments() (arg int,lenArgs int,lenDb int){
	    arg     = 0
		args := &Arguments{
			 		ListArgs : os.Args,
			 		Arg 	 : os.Args[1:],
			 		LenArg	 : len(os.Args),
			 		LenDb	 : len(Database),
 					}



 		num,err := strconv.Atoi(args.Arg[0])

 		if err == nil {
 			arg  = num
 		}
 		lenArgs  = args.LenArg
 		lenDb    = args.LenDb

 	return


}

func ConnectDatabase(arg int){

			fmt.Printf("Nama 		: %s\n",Database[arg-1].Nama)
			fmt.Printf("Alamat  	: %s\n",Database[arg-1].Alamat)
			fmt.Printf("Pekerjaan 	: %s\n",Database[arg-1].Pekerjaan)
			fmt.Printf("Alasan 		: %s\n",Database[arg-1].Alasan)
}


func ShowDatabase(){
	arg,lenargs,lendb := getArguments()
	// fmt.Printf("%v\n%d\n%d\n%d\n",os.Args,arg,lenargs,lendb)

    var show bool
    NolOrString := lenargs == 2 && arg == 0
    limitArg	:=  arg > lendb  && lenargs == 2
    limitParam	:= lenargs > 2
	switch show {
	case !NolOrString:
		fmt.Println("Jangan pakai string atau nol,harus pakai angka tapi bukan nol")
	case  !limitArg:
		fmt.Println("tidak ada database")
	case !limitParam:
		fmt.Println("terlalu banyak argument ")
	default:
			ConnectDatabase(arg)
		}


}