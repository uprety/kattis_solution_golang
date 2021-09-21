package main

import (
	"fmt"
)

func main() {

	var n, t, d, pred, pret, current_speed, max_speed  int;
	fmt.Scanln(&n)

	pret = 0
	pred = 0
	max_speed = 0

	fmt.Scanln(&t, &d)
	for i := 1; i < n; i++ {
		fmt.Scanln(&t, &d)
		current_speed = (d - pred) / (t - pret)
		if current_speed > max_speed {
			max_speed = current_speed
		}

		pret = t
		pred = d
		
	}

	fmt.Println(max_speed)
}