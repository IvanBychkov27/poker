package poker

import (
	"fmt"
	"strconv"
)

const (
	royalFlush    = "РоялФлеш"
	straightFlush = "СтритФлеш"
	care          = "Каре"
	fullHouse     = "ФуллХаус"
	flush         = "Флеш"
	straight      = "Стрит"
	set           = "Сет"
	twoPair       = "Две пары"
	pair          = "Пара"
	highCard      = "СтаршКарта"
)

// --- Печать инфо о кол-ве комбинаций ---
func (p *Poker) printStatCombination(statComb []int) (string, map[string]float64) {
	if len(statComb) != 11 {
		return "", nil
	}
	res := ""
	data := make(map[string]float64)

	n, rf := statComb[10], statComb[9]
	sf, k := statComb[8], statComb[7]
	fh, f := statComb[6], statComb[5]
	s, c := statComb[4], statComb[3]
	dp, pr := statComb[2], statComb[1]
	st := statComb[0]

	res = "Кол-во всех комбинаций = " + strconv.Itoa(n) + " шт.<br>"
	if rf != 0 {
		ver := float64(rf) * 100 / float64(n)
		data[royalFlush] = ver
		if ver < 0.01 {
			res += royalFlush + "  = " + strconv.Itoa(rf) + " шт. (" + fmt.Sprintf("%2.5f", ver) + " %)<br>"
		} else if ver < float64(1) {
			res += royalFlush + "  = " + strconv.Itoa(rf) + " шт. (" + fmt.Sprintf("%2.3f", ver) + " %)<br>"
		} else {
			res += royalFlush + "  = " + strconv.Itoa(rf) + " шт. (" + fmt.Sprintf("%2.1f", ver) + " %)<br>"
		}
	}
	if sf != 0 {
		ver := float64(sf) * 100 / float64(n)
		data[straightFlush] = ver
		if ver < float64(1) {
			res += straightFlush + "  = " + strconv.Itoa(sf) + " шт. (" + fmt.Sprintf("%2.3f", ver) + " %)<br>"
		} else {
			res += straightFlush + "  = " + strconv.Itoa(sf) + " шт. (" + fmt.Sprintf("%2.1f", ver) + " %)<br>"
		}
	}
	if k != 0 {
		ver := float64(k) * 100 / float64(n)
		data[care] = ver
		if ver < float64(1) {
			res += care + "&emsp;&emsp;&emsp; = " + strconv.Itoa(k) + " шт. (" + fmt.Sprintf("%2.3f", ver) + " %)<br>"
		} else {
			res += care + "&emsp;&emsp;&emsp; = " + strconv.Itoa(k) + " шт. (" + fmt.Sprintf("%2.1f", ver) + " %)<br>"
		}
	}
	if fh != 0 {
		ver := float64(fh) * 100 / float64(n)
		data[fullHouse] = ver
		if ver < float64(1) {
			res += fullHouse + "&ensp; = " + strconv.Itoa(fh) + " шт. (" + fmt.Sprintf("%2.3f", ver) + " %)<br>"
		} else {
			res += fullHouse + "&ensp; = " + strconv.Itoa(fh) + " шт. (" + fmt.Sprintf("%2.1f", ver) + " %)<br>"
		}
	}
	if f != 0 {
		ver := float64(f) * 100 / float64(n)
		data[flush] = ver
		if ver < float64(1) {
			res += flush + "&emsp;&emsp;&ensp; = " + strconv.Itoa(f) + " шт. (" + fmt.Sprintf("%2.3f", ver) + " %)<br>"
		} else {
			res += flush + "&emsp;&emsp;&ensp; = " + strconv.Itoa(f) + " шт. (" + fmt.Sprintf("%2.1f", ver) + " %)<br>"
		}
	}
	if s != 0 {
		ver := float64(s) * 100 / float64(n)
		data[straight] = ver
		if ver < float64(1) {
			res += straight + "&emsp;&emsp;&nbsp; = " + strconv.Itoa(s) + " шт. (" + fmt.Sprintf("%2.3f", ver) + " %)<br>"
		} else {
			res += straight + "&emsp;&emsp;&nbsp; = " + strconv.Itoa(s) + " шт. (" + fmt.Sprintf("%2.1f", ver) + " %)<br>"
		}
	}
	if c != 0 {
		ver := float64(c) * 100 / float64(n)
		data[set] = ver
		if ver < float64(1) {
			res += set + "&emsp;&emsp;&emsp;&ensp; = " + strconv.Itoa(c) + " шт. (" + fmt.Sprintf("%2.3f", ver) + " %)<br>"
		} else {
			res += set + "&emsp;&emsp;&emsp;&ensp; = " + strconv.Itoa(c) + " шт. (" + fmt.Sprintf("%2.1f", ver) + " %)<br>"
		}
	}
	if dp != 0 {
		ver := float64(dp) * 100 / float64(n)
		data[twoPair] = ver
		if ver < float64(1) {
			res += twoPair + "&nbsp;&nbsp; = " + strconv.Itoa(dp) + " шт. (" + fmt.Sprintf("%2.3f", ver) + " %)<br>"
		} else {
			res += twoPair + "&nbsp;&nbsp; = " + strconv.Itoa(dp) + " шт. (" + fmt.Sprintf("%2.1f", ver) + " %)<br>"
		}
	}
	if pr != 0 {
		ver := float64(pr) * 100 / float64(n)
		data[pair] = ver
		if ver < float64(1) {
			res += pair + "&emsp;&emsp;&emsp; = " + strconv.Itoa(pr) + " шт. (" + fmt.Sprintf("%2.3f", ver) + " %)<br>"
		} else {
			res += pair + "&emsp;&emsp;&emsp; = " + strconv.Itoa(pr) + " шт. (" + fmt.Sprintf("%2.1f", ver) + " %)<br>"
		}
	}
	if st != 0 {
		ver := float64(st) * 100 / float64(n)
		data[highCard] = ver
		if ver < float64(1) {
			res += highCard + " = " + strconv.Itoa(st) + " шт. (" + fmt.Sprintf("%2.3f", ver) + " %)<br>"
		} else {
			res += highCard + " = " + strconv.Itoa(st) + " шт. (" + fmt.Sprintf("%2.1f", ver) + " %)<br>"
		}
	}
	return res, data
}

// ---- перевод карты в строку ---
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
