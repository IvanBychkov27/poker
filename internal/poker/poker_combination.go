package poker

// --- покерная комбинация из 5 карт --- !!! pComb - д.б. отсортированным !!!
func (p *Poker) pokerCombination(pComb []Card) (strComb string, byteComb byte) {
	if len(pComb) != 5 {
		return
	}

	// --- "Флеш" ---
	if pComb[0].suil == pComb[1].suil && pComb[0].suil == pComb[2].suil {
		if pComb[0].suil == pComb[3].suil && pComb[0].suil == pComb[4].suil {
			strComb, byteComb = "Флеш", 5
		}
	}

	// --- "Стрит" - "СтритФлеш" - "РоялФлеш" ---
	if (pComb[0].value == pComb[1].value+1 || pComb[0].value == pComb[1].value+9) && pComb[1].value == pComb[2].value+1 {
		if pComb[2].value == pComb[3].value+1 && pComb[3].value == pComb[4].value+1 {
			if strComb != "Флеш" {
				return "Стрит", 4
			}
			if pComb[0].value == 14 && pComb[1].value == 13 {
				return "РоялФлеш", 9
			}
			return "СтритФлеш", 8
		}
	}
	if strComb != "" {
		return
	}

	// --- Каре ---
	if pComb[0].value == pComb[1].value && pComb[0].value == pComb[2].value && pComb[0].value == pComb[3].value ||
		pComb[1].value == pComb[2].value && pComb[1].value == pComb[3].value && pComb[1].value == pComb[4].value {
		return "Каре", 7
	}

	// --- ФуллХаус ---
	if pComb[0].value == pComb[1].value && pComb[0].value == pComb[2].value && pComb[3].value == pComb[4].value ||
		pComb[0].value == pComb[1].value && pComb[2].value == pComb[3].value && pComb[2].value == pComb[4].value {
		return "ФуллХаус", 6
	}

	// --- Сет ---
	if pComb[0].value == pComb[1].value && pComb[0].value == pComb[2].value ||
		pComb[1].value == pComb[2].value && pComb[1].value == pComb[3].value ||
		pComb[2].value == pComb[3].value && pComb[2].value == pComb[4].value {
		return "Сет", 3
	}

	// --- Пара - Две пары ---
	var para byte
	for i := 0; i < 4; i++ {
		if pComb[i].value == pComb[i+1].value {
			if para > 0 {
				return "Две пары", 2
			}
			para++
		}
	}

	if para > 0 {
		return "Пара", 1
	}

	strComb = "Старшая карта " + p.parCard(pComb[0])
	return
}
