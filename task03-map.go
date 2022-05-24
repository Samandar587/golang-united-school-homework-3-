package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0)
	slice := make([]string, 0)
	for key := range input {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, k := range keys {
		slice = append(slice, input[k])
	}
	return slice
}
