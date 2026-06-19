package kata

import "math/big"

func sum(n, p *big.Int, shift int64) *big.Int {
	nn := n.Int64()
	res := big.NewInt(0)
	four := big.NewInt(4)
	for k := int64(0); k <= nn; k++ {
		term := new(big.Int).Binomial(2*nn-k+shift, k)
		term.Mul(term, p)
		term.Mul(term, new(big.Int).Exp(four, big.NewInt(nn-k), nil))
		if k%2 == 1 {
			term.Neg(term)
		}
		res.Add(res, term)
	}
	return res
}

func V1(n, p *big.Int) *big.Int { return sum(n, p, 0) }

func U1(n, p *big.Int) *big.Int { return sum(n, p, 1) }

// v(n, p) = p * (2n + 1)
func VEff(n, p *big.Int) *big.Int {
	r := new(big.Int).Lsh(n, 1)
	r.Add(r, big.NewInt(1))
	return r.Mul(r, p)
}

// u(n, p) = p * (n + 1)
func UEff(n, p *big.Int) *big.Int {
	r := new(big.Int).Add(n, big.NewInt(1))
	return r.Mul(r, p)
}