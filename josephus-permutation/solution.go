package main

func Josephus(items []interface{}, k int) []interface{} {
	if len(items) == 0 {
		return []interface{}{}
	}

	result := []interface{}{}
	arr := make([]interface{}, len(items))
	copy(arr, items)

	index := 0
	for len(arr) > 0 {
		index = (index + k - 1) % len(arr)
		result = append(result, arr[index])
		arr = append(arr[:index], arr[index+1:]...)
		if len(arr) > 0 {
			index = index % len(arr)
		}
	}

	return result
}