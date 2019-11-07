package integers

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// 获取传递的参数的长度
	// lengthOfNumbers := len(numbersToSum)

	// sums = make([]int, lengthOfNumbers)
	// sums = make([]int, lengthOfNumbers)
	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }

	var sums []int
	for _, numbers := range numbersToSum {
		s := Sum(numbers)
		fmt.Println(s)
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

func main() {
	len := 4
	a := []int{1, 2}
	b := [3]int{1, 2}
	c := make([]int, len)

	fmt.Printf("%T %#v,%T %#v,%T %#v", a, a, b, b, c, c)
}
