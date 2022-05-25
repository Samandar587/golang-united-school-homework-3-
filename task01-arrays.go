package homework

func average(input []float32) (result float32) {
	var sum float32 = 0
	for i := 0; i < len(input); i++ {
		sum = sum + input[i]
	}
	length := float32(len(input))
	result = sum / length
	return result
}
