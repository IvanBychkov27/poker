package poker

//--- максимальная покерная комбинация (5 карт) из всех вариантов из 7 карт  ---
func (p *Poker) maxPokerCombinationOf7Cards(pokerComb7cards []Card) []Card {
	maxComb := make([][]Card, 1)
	maxComb[0] = make([]Card, 5)
	allTempComb, _ := p.allCombinations_5_Cards(pokerComb7cards)
	maxComb[0] = p.maxCombOfAllComb5card(allTempComb)
	return maxComb[0]
}

// --- выбор максимальной комбинации из всех наборов комбинаций по 5 карт --- !!!
func (p *Poker) maxCombOfAllComb5card(allComb [][]Card) []Card {
	maxComb := make([]Card, 1)
	maxComb = allComb[0]
	for i := 1; i < len(allComb); i++ {
		max := p.maxCombination(maxComb, allComb[i])
		if max > 1 {
			maxComb = allComb[i]
		}
	}
	return maxComb
}

// --- if comb1 > comb2 - 1;  if comb1 < comb2 - 2 - if comb1 = comb2 - 0 --- комбинации должны содержать 5 карт---!!!
func (p *Poker) maxCombination(comb1, comb2 []Card) byte {
	funcMax := byte(0)
	_, nComb1 := p.pokerCombination(comb1)
	_, nComb2 := p.pokerCombination(comb2)
	if nComb1 > nComb2 {
		funcMax = 1
	}
	if nComb1 < nComb2 {
		funcMax = 2
	}
	if nComb1 == nComb2 {
		comb1 = p.sortComb(comb1)
		comb2 = p.sortComb(comb2)
		for i := 0; i < 5; i++ {
			if comb1[i].value > comb2[i].value {
				funcMax = 1
				break
			}
			if comb1[i].value < comb2[i].value {
				funcMax = 2
				break
			}
			if comb1[i].value == comb2[i].value {
				funcMax = 0
			}
		}
	}
	return funcMax
}
