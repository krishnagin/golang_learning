package exercises

func EvenNumbers(list []int) (result []int) {

	for _, v := range list {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return
}

func OddNumbers(list []int) (result []int) {

	for _, v := range list {
		if v%2 != 0 {
			result = append(result, v)
		}
	}
	return
}

func PrimeNumbers(list []int) (result []int) {

	for _, v := range list {
		if v != 0 && v != 1 && isPrime(v) {
			result = append(result, v)
		}
	}

	return
}

func PrimeNumbersAndOdd(list []int) (result []int) {

	for _, v := range list {
		if v != 0 && v != 1 && isPrime(v) && v%2 != 0 {
			result = append(result, v)
		}
	}

	return
}

func EvenAndMulitplesofFive(list []int) (result []int) {

	for _, v := range list {
		if v%2 == 0 && v%5 == 0 {
			result = append(result, v)
		}
	}
	return
}

func isPrime(n int) (result bool) {
	incoming := n
	var times int
	for ; n > 0; n-- {
		if incoming%n == 0 {
			times++
		}
	}
	if times <= 2 {
		result = true
	}
	return
}

func FunctionHandler(list []int, o func(list []int) []int) []int {
	return o(list)
}
