package main

import (
	"fmt"
)

func main(){

	fmt.Println("1")
	defer func() {
		if err := recover(); err != nil {
			//做异常处理
			if err != nil {
				fmt.Println("2",err)
			}
		}
	}()

	panic("222")



	fmt.Println("3")


}



