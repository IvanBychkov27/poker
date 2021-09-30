package poker

import (
	"sort"
)

// --- сортировка комбинации карт от большего к меньшему ---
func (p *Poker) sortComb(pComb []Card) []Card { // 2 sec
	for m := 0; m < len(pComb)-1; m++ {
		for i := 1; i < len(pComb); i++ {
			if pComb[i].value > pComb[i-1].value {
				pComb[i-1], pComb[i] = pComb[i], pComb[i-1]
			}
		}
	}
	return pComb
}

func (p *Poker) sortComb2(pComb []Card) []Card { // 7sec
	sort.SliceStable(pComb, func(i, j int) bool {
		return pComb[i].value > pComb[j].value // сортировка по убыванию рейтинга
		//return pComb[i].value < pComb[j].value // сортировка по возрастанию рейтинга
	})
	return pComb
}
