package main

import "fmt"

func main() {
	data := 3

	AddFive(&data)

	fmt.Println("data:", data)
}

func AddFive(data *int) {
	*data += 5
}
