package poker

// --- покерная комбинация из 5 карт --- !!! pComb - д.б. отсортированным !!!
func (p *Poker) pokerCombination(pComb []Card) (string, byte) {
	strComb := ""
	byteComb := byte(0)
	if len(pComb) != 5 {
		return strComb, byteComb
	}

	// --- "Флеш" ---
	if pComb[0].suil == pComb[1].suil && pComb[0].suil == pComb[2].suil {
		if pComb[0].suil == pComb[3].suil && pComb[0].suil == pComb[4].suil {
			strComb = "Флеш"
			byteComb = 5
		}
	}

	// --- "Стрит" - "СтритФлеш" - "РоялФлеш" ---
	if (pComb[0].value == pComb[1].value+1 || pComb[0].value == pComb[1].value+9) && pComb[1].value == pComb[2].value+1 {
		if pComb[2].value == pComb[3].value+1 && pComb[3].value == pComb[4].value+1 {
			if strComb == "Флеш" {
				if pComb[0].value == 14 && pComb[1].value == 13 {
					strComb = "РоялФлеш"
					byteComb = 9
				} else {
					strComb = "СтритФлеш"
					byteComb = 8
				}
			} else {
				strComb = "Стрит"
				byteComb = 4
			}
		}
	}

	if strComb != "" {
		return strComb, byteComb
	}

	// --- Каре ---
	if pComb[0].value == pComb[1].value && pComb[0].value == pComb[2].value && pComb[0].value == pComb[3].value {
		strComb = "Каре"
		byteComb = 7
	}
	if pComb[1].value == pComb[2].value && pComb[1].value == pComb[3].value && pComb[1].value == pComb[4].value {
		strComb = "Каре"
		byteComb = 7
	}
	if strComb != "" {
		return strComb, byteComb
	}

	// --- ФуллХаус ---
	if pComb[0].value == pComb[1].value && pComb[0].value == pComb[2].value && pComb[3].value == pComb[4].value {
		strComb = "ФуллХаус"
		byteComb = 6
	}
	if pComb[0].value == pComb[1].value && pComb[2].value == pComb[3].value && pComb[2].value == pComb[4].value {
		strComb = "ФуллХаус"
		byteComb = 6
	}
	if strComb != "" {
		return strComb, byteComb
	}

	// --- Сет ---
	if pComb[0].value == pComb[1].value && pComb[0].value == pComb[2].value {
		strComb = "Сет"
		byteComb = 3
	}
	if pComb[1].value == pComb[2].value && pComb[1].value == pComb[3].value {
		strComb = "Сет"
		byteComb = 3
	}
	if pComb[2].value == pComb[3].value && pComb[2].value == pComb[4].value {
		strComb = "Сет"
		byteComb = 3
	}
	if strComb != "" {
		return strComb, byteComb
	}

	// --- Пара - Две пары ---
	var para, doublPara byte
	for i := 0; i < 4; i++ {
		if pComb[i].value == pComb[i+1].value {
			if para > 0 {
				doublPara++
			} else {
				para++
			}
		}
	}

	if doublPara > 0 {
		strComb = "Две пары"
		byteComb = 2
	} else {
		if para > 0 {
			strComb = "Пара"
			byteComb = 1
		}
	}

	if strComb != "" {
		return strComb, byteComb
	}

	strComb = "Старшая карта " + p.parCard(pComb[0])
	return strComb, byteComb
}
