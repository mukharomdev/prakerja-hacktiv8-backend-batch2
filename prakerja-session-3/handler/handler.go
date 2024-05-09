package handler

import (
	"fmt"
	"os"
	"strconv"
  . "prakerja-session-3/database"
  . "prakerja-session-3/datatype"
  //. "prakerja-session-3/reflection"

)




func getArguments() (arg int,lenArgs int,lenDb int){
	    arg     = 0
		args := &Arguments{
			 		ListArgs : os.Args,
			 		Arg 	 : os.Args[1:][0],
			 		LenArg	 : len(os.Args),
			 		LenDb	 : len(Database),
 					}



 		num,err := strconv.Atoi(args.Arg)

 		if err == nil {
 			arg  = num
 		}
 		lenArgs  = args.LenArg
 		lenDb    = args.LenDb

 	return


}


func ShowDatabase(){
	arg,lenargs,lendb := getArguments()
	fmt.Printf("%v\n%d\n%d\n%d\n",os.Args,arg,lenargs,lendb)

    var show bool
    NolOrString := lenargs == 2 || arg == 0
    limitArg	:= lendb > arg+1 && lenargs > 2
    limitParam	:= lenargs > 2 || lenargs < 2
	switch show {
	case !NolOrString:
		fmt.Println("Jangan pakai string atau nol,harus pakai angka tapi bukan nol")
	case  !limitArg:
		fmt.Println("tidak ada database")
	case !limitParam:
		fmt.Println("terlalu banyak argument ")
	default:
		fmt.Println("buka data")
	}


}