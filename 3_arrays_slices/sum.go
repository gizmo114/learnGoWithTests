package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var return_slice []int
	for _, slice := range numbers {
		return_slice = append(return_slice, Sum(slice))
	}
	return return_slice
}
