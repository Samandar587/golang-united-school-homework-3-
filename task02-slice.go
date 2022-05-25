package homework

func reverse(input []int64) (result []int64) {
	input1 := make([]int64, 0)
	k := len(input) - 1
	for i := 0; i < len(input); i++ {
		input1 = append(input1, input[k])
		k--
	}
	return input1
}
