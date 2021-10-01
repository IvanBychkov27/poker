package poker

type Card struct {
	value byte
	suil  byte
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
