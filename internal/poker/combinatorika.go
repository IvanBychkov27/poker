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

// --- расчет кол-ва комбинаций по формуле комбинаторики:  k из n  =  n! / (k! * (n - k)! )  - где k < n !!! ---
func (p *Poker) combinatorika(k, n uint64) uint64 {
	var comb, chislitel uint64
	chislitel = n
	for i := uint64(1); i < k; i++ {
		chislitel *= n - i
	}
	comb = chislitel / p.factorial(k)
	return comb
}

// --- расчет факториала ---
func (p *Poker) factorial(x uint64) uint64 {
	if x == 0 {
		return 1
	}
	return x * p.factorial(x-1)
}
