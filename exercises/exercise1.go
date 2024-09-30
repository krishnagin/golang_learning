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
