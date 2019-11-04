package main

import "fmt"

func main() {
	x := []int{1,2,2,3,4,4,4,5,66,7,1}
	//x = remove(x)
	fmt.Printf("%d", x)

	s := []int{1,2,3,10}
	fmt.Println(s)
	copy(s, []int{4,5,6,7,8})
	fmt.Println(s)
}

func remove(slice []int) []int {
	for i := range slice {
		if i > len(slice) -1 {
			return slice
		}

		fmt.Printf("%d\n", slice)

		if i < len(slice) - 1 && slice[i] == slice[i+1] {
			copy(slice[i:], slice[i+1:])
			slice = slice[:len(slice) - 1]
			fmt.Printf("%d \n", slice)
			return remove(slice)
		}
	}

	return slice
}

