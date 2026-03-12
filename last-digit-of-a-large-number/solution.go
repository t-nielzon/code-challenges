package kata

import "math/big"

func LastDigit(n1, n2 *big.Int) int {
	// 0^0 = 1 by convention
	if n2.Sign() == 0 {
		return 1
	}

	base := new(big.Int).Mod(n1, big.NewInt(10)).Int64()

	// any cycle of last digits repeats with period dividing 4
	mod4 := new(big.Int).Mod(n2, big.NewInt(4)).Int64()

	// map mod4==0 to exponent 4 so we never compute x^0 incorrectly
	if mod4 == 0 {
		mod4 = 4
	}

	result := int64(1)
	for i := int64(0); i < mod4; i++ {
		result = (result * base) % 10
	}
	return int(result)
}