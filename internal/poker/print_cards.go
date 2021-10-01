package poker

import (
	"fmt"
	"strconv"
)

// --- Печать инфо о кол-ве комбинаций ---
func (p *Poker) printStatCombination(statComb []int) string {
	res := ""
	if len(statComb) == 11 {
		n, rf := statComb[10], statComb[9]
		sf, k := statComb[8], statComb[7]
		fh, f := statComb[6], statComb[5]
		s, c := statComb[4], statComb[3]
		dp, pr := statComb[2], statComb[1]
		st := statComb[0]

		res = "Кол-во всех комбинаций = " + strconv.Itoa(n) + " шт.<br>"
		if rf != 0 {
			ver := float64(rf) / float64(n)
			if ver < 0.0001 {
				res += "РоялФлеш  = " + strconv.Itoa(rf) + " шт. (" + fmt.Sprintf("%2.5f", (float64(rf)*100)/float64(n)) + " %)<br>"
			} else if ver < 0.01 {
				res += "РоялФлеш  = " + strconv.Itoa(rf) + " шт. (" + fmt.Sprintf("%2.3f", (float64(rf)*100)/float64(n)) + " %)<br>"
			} else {
				res += "РоялФлеш  = " + strconv.Itoa(rf) + " шт. (" + fmt.Sprintf("%2.1f", (float64(rf)*100)/float64(n)) + " %)<br>"
			}
		}
		if sf != 0 {
			if float64(sf)/float64(n) < 0.01 {
				res += "СтритФлеш  = " + strconv.Itoa(sf) + " шт. (" + fmt.Sprintf("%2.3f", (float64(sf)*100)/float64(n)) + " %)<br>"
			} else {
				res += "СтритФлеш  = " + strconv.Itoa(sf) + " шт. (" + fmt.Sprintf("%2.1f", (float64(sf)*100)/float64(n)) + " %)<br>"
			}
		}
		if k != 0 {
			if float64(k)/float64(n) < 0.01 {
				res += "Каре&emsp;&emsp;&emsp; = " + strconv.Itoa(k) + " шт. (" + fmt.Sprintf("%2.3f", (float64(k)*100)/float64(n)) + " %)<br>"
			} else {
				res += "Каре&emsp;&emsp;&emsp; = " + strconv.Itoa(k) + " шт. (" + fmt.Sprintf("%2.1f", (float64(k)*100)/float64(n)) + " %)<br>"
			}
		}
		if fh != 0 {
			if float64(fh)/float64(n) < 0.01 {
				res += "ФуллХаус&ensp; = " + strconv.Itoa(fh) + " шт. (" + fmt.Sprintf("%2.3f", (float64(fh)*100)/float64(n)) + " %)<br>"
			} else {
				res += "ФуллХаус&ensp; = " + strconv.Itoa(fh) + " шт. (" + fmt.Sprintf("%2.1f", (float64(fh)*100)/float64(n)) + " %)<br>"
			}
		}
		if f != 0 {
			if float64(f)/float64(n) < 0.01 {
				res += "Флеш&emsp;&emsp;&ensp; = " + strconv.Itoa(f) + " шт. (" + fmt.Sprintf("%2.3f", (float64(f)*100)/float64(n)) + " %)<br>"
			} else {
				res += "Флеш&emsp;&emsp;&ensp; = " + strconv.Itoa(f) + " шт. (" + fmt.Sprintf("%2.1f", (float64(f)*100)/float64(n)) + " %)<br>"
			}
		}
		if s != 0 {
			if float64(s)/float64(n) < 0.01 {
				res += "Стрит&emsp;&emsp;&nbsp; = " + strconv.Itoa(s) + " шт. (" + fmt.Sprintf("%2.3f", (float64(s)*100)/float64(n)) + " %)<br>"
			} else {
				res += "Стрит&emsp;&emsp;&nbsp; = " + strconv.Itoa(s) + " шт. (" + fmt.Sprintf("%2.1f", (float64(s)*100)/float64(n)) + " %)<br>"
			}
		}
		if c != 0 {
			if float64(c)/float64(n) < 0.01 {
				res += "Сет&emsp;&emsp;&emsp;&ensp; = " + strconv.Itoa(c) + " шт. (" + fmt.Sprintf("%2.3f", (float64(c)*100)/float64(n)) + " %)<br>"
			} else {
				res += "Сет&emsp;&emsp;&emsp;&ensp; = " + strconv.Itoa(c) + " шт. (" + fmt.Sprintf("%2.1f", (float64(c)*100)/float64(n)) + " %)<br>"
			}
		}
		if dp != 0 {
			if float64(dp)/float64(n) < 0.01 {
				res += "Две пары&nbsp;&nbsp; = " + strconv.Itoa(dp) + " шт. (" + fmt.Sprintf("%2.3f", (float64(dp)*100)/float64(n)) + " %)<br>"
			} else {
				res += "Две пары&nbsp;&nbsp; = " + strconv.Itoa(dp) + " шт. (" + fmt.Sprintf("%2.1f", (float64(dp)*100)/float64(n)) + " %)<br>"
			}
		}
		if pr != 0 {
			if float64(pr)/float64(n) < 0.01 {
				res += "Пара&emsp;&emsp;&emsp; = " + strconv.Itoa(pr) + " шт. (" + fmt.Sprintf("%2.3f", (float64(pr)*100)/float64(n)) + " %)<br>"
			} else {
				res += "Пара&emsp;&emsp;&emsp; = " + strconv.Itoa(pr) + " шт. (" + fmt.Sprintf("%2.1f", (float64(pr)*100)/float64(n)) + " %)<br>"
			}
		}
		if st != 0 {
			if float64(pr)/float64(n) < 0.01 {
				res += "СтаршКарта = " + strconv.Itoa(st) + " шт. (" + fmt.Sprintf("%2.3f", (float64(st)*100)/float64(n)) + " %)<br>"
			} else {
				res += "СтаршКарта = " + strconv.Itoa(st) + " шт. (" + fmt.Sprintf("%2.1f", (float64(st)*100)/float64(n)) + " %)<br>"
			}
		}
	}
	return res
}

func (p *Poker) nameCombination(n byte) string {
	nComb := ""
	switch n {
	case 1: //"Пара":
		nComb = "Пара"
	case 2: // "Две пары":
		nComb = "Две пары"
	case 3: //"Сет":
		nComb = "Сет"
	case 4: //"Стрит":
		nComb = "Стрит"
	case 5: //"Флеш":
		nComb = "Флеш"
	case 6: //"ФуллХаус":
		nComb = "ФуллХаус"
	case 7: //"Каре":
		nComb = "Каре"
	case 8: //"СтритФлеш":
		nComb = "СтритФлеш"
	case 9: //"РоялФлеш":
		nComb = "РоялФлеш"
	default:
		nComb = "СтаршКарта"
	}
	return nComb
}

// --- печать колоды карт ---
func (p *Poker) printDeckCards(deckCards []Card) {
	var predCard byte
	for _, c := range deckCards {
		strCard := p.parCard(c)
		if c.value < predCard {
			fmt.Println()
		}
		fmt.Print(strCard, " ")
		predCard = c.value
	}
	fmt.Println()
}

// --- печать комбинации карт ---
func (p *Poker) printComb(cardsComb []Card) {
	for _, c := range cardsComb {
		strCard := p.parCard(c)
		fmt.Print(strCard, " ")
	}
	fmt.Println()
}

// --- комбинации карт ---
func (p *Poker) strComb(cardsComb []Card) string {
	result := ""
	for _, c := range cardsComb {
		strCard := p.parCard(c)
		result = result + " " + strCard
	}
	return result
}

// --- параметр карты - перевод карты в строку ---
func (p *Poker) parCard(c Card) string {
	parCardString := ""
	switch c.value {
	case 2:
		parCardString = "2"
	case 3:
		parCardString = "3"
	case 4:
		parCardString = "4"
	case 5:
		parCardString = "5"
	case 6:
		parCardString = "6"
	case 7:
		parCardString = "7"
	case 8:
		parCardString = "8"
	case 9:
		parCardString = "9"
	case 10:
		parCardString = "10"
	case 11:
		parCardString = "В"
	case 12:
		parCardString = "Д"
	case 13:
		parCardString = "К"
	case 14:
		parCardString = "Т"
	}

	switch c.suil {
	case 1:
		parCardString += "п"
	case 2:
		parCardString += "к"
	case 3:
		parCardString += "б"
	case 4:
		parCardString += "ч"
	}

	return parCardString
}
