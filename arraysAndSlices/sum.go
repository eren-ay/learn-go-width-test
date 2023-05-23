package main

func Sum(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func sumAll(numbers ...[]int) []int {
	lengthOfNumbers := len(numbers)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbers {
		sums[i] = Sum(numbers)
	}

	return sums
}
