package main

import "fmt"

func split(sum int) (z, x, y int) {
	x = sum * 4 / 9
	y = sum - x
	
	z = y - x
	return z, z, z
}

func main() {
	fmt.Println(split(1000))
}
