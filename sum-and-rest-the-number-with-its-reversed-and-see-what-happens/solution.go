package kata

var sumRestCache []int
var sumRestNext int = 10

func sumRestReverse(n int) int {
	r := 0
	for n > 0 {
		r = r*10 + n%10
		n /= 10
	}
	return r
}

func sumRestGenerate(upTo int) {
	for len(sumRestCache) < upTo {
		i := sumRestNext
		if i%10 != 0 {
			r := sumRestReverse(i)
			if r != i {
				d := i - r
				if d < 0 {
					d = -d
				}
				if (i+r)%d == 0 {
					sumRestCache = append(sumRestCache, i)
				}
			}
		}
		sumRestNext++
	}
}

func SumRest(n int) int {
	sumRestGenerate(n)
	return sumRestCache[n-1]
}