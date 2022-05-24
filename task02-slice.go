package homework

func reverse(input []int64) (result []int64) {
	input1 := make([]int64, 4)
	k := len(input) - 1
	for index := range input {
		input1[index] = input[k]
		k--
		index++
	}
	return input1
}
