package main

import "sort"

func MaxPartitionProduct(n int) ([][]int, int) {
	var partitions [][]int
	var maxProduct int

	if n == 1 {
		partitions = append(partitions, []int{1})
		return partitions, 1
	}

	remainder := n % 3

	switch remainder {
	case 0:
		partition := make([]int, n/3)
		for i := range partition {
			partition[i] = 3
		}
		partitions = append(partitions, partition)
		maxProduct = 1
		for i := 0; i < n/3; i++ {
			maxProduct *= 3
		}

	case 1:
		threeCount := (n - 4) / 3

		p1 := make([]int, threeCount+1)
		p1[0] = 4
		for i := 1; i < len(p1); i++ {
			p1[i] = 3
		}
		sort.Sort(sort.Reverse(sort.IntSlice(p1)))
		partitions = append(partitions, p1)

		p2 := make([]int, threeCount+2)
		p2[0] = 2
		p2[1] = 2
		for i := 2; i < len(p2); i++ {
			p2[i] = 3
		}
		sort.Sort(sort.Reverse(sort.IntSlice(p2)))
		partitions = append(partitions, p2)

		maxProduct = 1
		for i := 0; i < threeCount; i++ {
			maxProduct *= 3
		}
		maxProduct *= 4

	case 2:
		partition := make([]int, (n-2)/3+1)
		partition[0] = 2
		for i := 1; i < len(partition); i++ {
			partition[i] = 3
		}
		sort.Sort(sort.Reverse(sort.IntSlice(partition)))
		partitions = append(partitions, partition)
		maxProduct = 1
		for i := 0; i < (n-2)/3; i++ {
			maxProduct *= 3
		}
		maxProduct *= 2
	}

	sort.Slice(partitions, func(i, j int) bool {
		return len(partitions[i]) < len(partitions[j])
	})

	return partitions, maxProduct
}