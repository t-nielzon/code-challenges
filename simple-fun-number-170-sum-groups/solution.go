package main

func SumGroups(arr []int) int {
	for {
		newArr := []int{}
		i := 0

		for i < len(arr) {
			sum := arr[i]
			j := i + 1

			for j < len(arr) && (arr[i]%2 == arr[j]%2) {
				sum += arr[j]
				j++
			}

			newArr = append(newArr, sum)
			i = j
		}

		if len(newArr) == len(arr) {
			return len(newArr)
		}

		arr = newArr
	}
}