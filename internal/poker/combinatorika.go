package poker

import (
	"math"
)

// --- Процент побед при разных количествах соперников ---
func (p *Poker) percentWinsDiffNumbersOpponents(x float64, n int) float64 {
	var result float64
	if n > 15 {
		n = 15
	}
	if n < 0 {
		n = 0
	}
	x = 1 - x
	result = float64(n) * x
	for m := 2; m < n+1; m++ {
		result += math.Pow(float64(-1), float64(m-1)) * float64(p.combinatorika(uint64(m), uint64(n))) * math.Pow(x, float64(m))
	}
	return 1 - result
}

// --- расчет кол-ва комбинаций по формуле комбинаторики:  k из n  =  n! / (k! * (n - k)! )  - где k < n !!! --- сочетания без повторений
func (p *Poker) combinatorika(k, n uint64) uint64 {
	num := n
	for i := uint64(1); i < k; i++ {
		num *= n - i
	}
	return num / p.factorial(k)
}

// --- расчет факториала ---
func (p *Poker) factorial(x uint64) uint64 { // алгоритм "в лоб"
	res := uint64(1)
	switch x {
	case 2:
		res = uint64(2)
	case 3:
		res = uint64(6)
	case 4:
		res = uint64(24)
	case 5:
		res = uint64(120)
	case 6:
		res = uint64(720)
	case 7:
		res = uint64(5040)
	case 8:
		res = uint64(40320)
	case 9:
		res = uint64(362880)
	case 10:
		res = uint64(3628800)
	case 11:
		res = uint64(39916800)
	case 12:
		res = uint64(479001600)
	case 13:
		res = uint64(6227020800)
	case 14:
		res = uint64(87178291200)
	case 15:
		res = uint64(1307674368000)
	case 16:
		res = uint64(20922789888000)
	case 17:
		res = uint64(355687428096000)
	case 18:
		res = uint64(6402373705728000)
	case 19:
		res = uint64(121645100408832000)
	case 20:
		res = uint64(2432902008176640000)
	case 21:
		res = uint64(14197454024290336768)
	default:
		p.logger.Debug("error factorial x > 21")
	}

	return res
}

//------------------------
func (p *Poker) factorial_02(n uint64) uint64 { // наивный алгоритм
	res := uint64(1)
	for i := uint64(2); i < n+1; i++ {
		res *= i
	}
	return res
}

func (p *Poker) factorial_01(x uint64) uint64 { // рекурсивный алгоритм
	if x == 0 {
		return 1
	}
	return x * p.factorial_01(x-1)
}
