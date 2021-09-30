package poker

type Card struct {
	value byte
	suil  byte
}

// --- заполнение массива всеми комбинациями из 5-ти карт из колоды карт ---
func (p *Poker) allCombinations_5_Cards(deckCards []Card) ([][]Card, int) {
	len_deckCards := len(deckCards)
	lenAllComb := p.combinatorika(5, uint64(len_deckCards))
	allComb := make([][]Card, lenAllComb)
	for i := range allComb {
		allComb[i] = make([]Card, 5)
	}
	comb := 0
	for i := 0; i < len_deckCards-4; i++ {
		for j := 1 + i; j < len_deckCards-3; j++ {
			for k := 1 + j; k < len_deckCards-2; k++ {
				for l := 1 + k; l < len_deckCards-1; l++ {
					for m := 1 + l; m < len_deckCards; m++ {
						allComb[comb][0] = deckCards[i]
						allComb[comb][1] = deckCards[j]
						allComb[comb][2] = deckCards[k]
						allComb[comb][3] = deckCards[l]
						allComb[comb][4] = deckCards[m]
						comb++
					}
				}
			}
		}
	}
	return allComb, comb
}

// --- колода карт без вышедших карт ---
func (p *Poker) deckCards_NoFull(cardDel []Card) []Card {
	var c Card
	deckNoFull := make([]Card, 0, 52)
	for s := byte(1); s < 5; s++ {
		for v := byte(2); v < 15; v++ {
			f := true
			c.value = v
			c.suil = s
			for _, c_del := range cardDel {
				if c_del == c {
					f = false
					break
				}
			}
			if f {
				deckNoFull = append(deckNoFull, c)
			}
		}
	}
	return deckNoFull
}

//--- полная калода карт из 52 шт ---
func (p *Poker) deckCardsFull() []Card {
	var c Card
	deckFull := make([]Card, 0, 52)
	for s := byte(1); s < 5; s++ {
		for v := byte(2); v < 15; v++ {
			c.value = v
			c.suil = s
			deckFull = append(deckFull, c)
		}
	}
	return deckFull
}
