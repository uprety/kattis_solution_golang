package main

import "fmt"

func main() {
	var n, current int;
	total := 0
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &current)
		total += current
	}
	fmt.Println(total)
}
