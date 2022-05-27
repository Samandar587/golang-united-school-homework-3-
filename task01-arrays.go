package homework

func average(input [50]float32) (result float32) {
	var sum float32 = 0
	for i := 0; i < len(input); i++ {
		sum = sum + input[i]
	}
	count := 0
	for j := 1; j < len(input); j++ {
		count = count + 1
	}
	length := float32(count)
	result = sum / length
	return result
}
