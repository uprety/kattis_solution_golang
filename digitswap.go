package main

import "fmt"

func main(){
	var digit string
	fmt.Scanln(&digit)

	for i := len(digit) - 1; i > -1; i-- {
		fmt.Printf("%c", digit[i])
	}
	fmt.Println("")
}
