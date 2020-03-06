package main

import "fmt"

func main(){

	fmt.Println(div_zero(5,0)) 
	fmt.Println(div_zero(6,3))

}

func div_zero( a, b int) int {

	defer func(){
		fmt.Println(recover())
	}()

	sol:= a/b
	return sol
}
