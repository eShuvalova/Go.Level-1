package main

import "fmt"

func main() {
	arr := []int{14, 10, 8, 4, 5, 6, 2}

	for i := 1; i < len(arr); i++ {
		number := arr[i]

		for j := i - 1; j >= 0; j-- {
			if number < arr[j] {
				arr[j+1] = arr[j]
			} else {
				arr[j+1] = number
				break
			}
			arr[j] = number
		}
	}
	fmt.Println(arr)
}
