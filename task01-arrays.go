package homework

func average(input [15]float32) (result float32) {
	var sum float32 = 0
	for key, _ := range input {
		sum += input[key]
	}
	length := float32(len(input))
	res := sum / length
	return res
}
