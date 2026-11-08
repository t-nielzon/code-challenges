package main

func SumGroups(arr []int) int {
	for {
		newArr := []int{}
		i := 0

		for i < len(arr) {
			sum := arr[i]
			isEven := arr[i]%2 == 0
			j := i + 1

			for j < len(arr) && (arr[j]%2 == 0) == isEven {
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