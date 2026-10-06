package main

func sumGroups(arr []int) int {
	for {
		newArr := make([]int, 0)
		i := 0
		
		for i < len(arr) {
			sum := arr[i]
			parity := arr[i] % 2
			j := i + 1
			
			for j < len(arr) && arr[j]%2 == parity {
				sum += arr[j]
				j++
			}
			
			newArr = append(newArr, sum)
			i = j
		}
		
		if len(newArr) == len(arr) {
			return len(arr)
		}
		
		arr = newArr
	}
}