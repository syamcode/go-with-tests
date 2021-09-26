package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var result []int
	for _, numbers := range numbersToSum {
		sum := 0
		for _, number := range numbers {
			sum += number
		}
		result = append(result, sum)
	}
	return result
}
