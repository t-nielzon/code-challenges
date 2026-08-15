package kata

import (
	"regexp"
	"strconv"
)

func SumOfCubic(s string) string {
	isCubic := func(n int) bool {
		sum := 0
		temp := n
		for temp > 0 {
			digit := temp % 10
			sum += digit * digit * digit
			temp /= 10
		}
		return sum == n
	}

	re := regexp.MustCompile(`\d+`)
	digitGroups := re.FindAllString(s, -1)

	var cubicNumbers []int
	total := 0

	for _, group := range digitGroups {
		for len(group) > 3 {
			chunk := group[:3]
			num, _ := strconv.Atoi(chunk)
			if isCubic(num) {
				cubicNumbers = append(cubicNumbers, num)
				total += num
			}
			group = group[3:]
		}

		if len(group) > 0 {
			num, _ := strconv.Atoi(group)
			if isCubic(num) {
				cubicNumbers = append(cubicNumbers, num)
				total += num
			}
		}
	}

	if len(cubicNumbers) == 0 {
		return "Unlucky"
	}

	result := ""
	for _, num := range cubicNumbers {
		result += strconv.Itoa(num) + " "
	}
	result += strconv.Itoa(total) + " Lucky"
	return result
}