package main


import "fmt"


func main() {
	var baris int = 10
	var batas int

	for i := 1; i <= baris; i++ {
		batas = 0
		for spasi := 1; spasi <= baris-i; spasi++ {
			fmt.Print("  ")
		}
		for {
			fmt.Print("* ")
			batas++
			if batas == 2*i-1 {
				break
			}
		}
		fmt.Println("")
	}

}


