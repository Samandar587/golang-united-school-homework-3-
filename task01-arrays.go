package homework

func average(input [15]float32) (result float32) {
	var sum float32 = 0
	for _, value := range input {
		sum = sum + value
	}
	return sum / 6
}
