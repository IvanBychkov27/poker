package poker

// --- сортировка комбинации карт от большего к меньшему ---
func (p *Poker) sortComb(pComb []Card) []Card {
	for m := 0; m < len(pComb)-1; m++ {
		for i := 1; i < len(pComb); i++ {
			if pComb[i].value > pComb[i-1].value {
				pComb[i-1], pComb[i] = pComb[i], pComb[i-1]
			}
		}
	}
	return pComb
}
