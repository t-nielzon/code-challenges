package kata

import "math/big"

func LastDigit(a, b *big.Int) int {
	zero := big.NewInt(0)
	if b.Cmp(zero) == 0 {
		return 1
	}
	ten := big.NewInt(10)
	four := big.NewInt(4)

	aMod := new(big.Int).Mod(a, ten).Int64()
	bMod := new(big.Int).Mod(b, four).Int64()

	exp := bMod
	if b.Cmp(four) >= 0 {
		exp += 4
	}

	result := int64(1)
	base := aMod
	for i := int64(0); i < exp; i++ {
		result = (result * base) % 10
	}
	return int(result)
}