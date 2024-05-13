package main

import (
	"fmt"
	"sync"
)


func main(){

			List := []int{1,2,3,4,5,6,7,8,9,10}

			var wg sync.WaitGroup

		 	func(in []int){

				for _,L := range List{
				wg.Add(1)
				go Print(&wg,L)

				}

			}(List)



			wg.Wait()
}

func Print(wg *sync.WaitGroup,i int){
	defer wg.Done()
	fmt.Println(i)
}