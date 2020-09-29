package main

import "fmt"

func main() {

	var limite float64
	var producto float64 = 1
	var suma float64 = 1
	var i float64 = 1

	fmt.Scanln(&limite)
	for i <= limite {
		producto = producto + i
		suma = suma + (1 / producto)
		i = i + 1
	}
	fmt.Print(suma)

}
