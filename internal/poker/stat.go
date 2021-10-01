package poker

import (
	"fmt"
)

//--- статистика максимальных комбинаций из 7 карт: 1 карт из оставшейся колоды и + 6 карт на руках ---
func (p *Poker) statAllMaxCombHand_6_cards(deckCards, handCards []Card) []int {
	if len(handCards) != 6 {
		fmt.Println("invalid combination handCards")
		return nil
	}
	statResult := make([]int, 11, 11)
	stat := make([]int, 11, 11)

	for i := 0; i < len(deckCards); i++ {
		var tempComb = []Card{
			deckCards[i],
			handCards[0],
			handCards[1],
			handCards[2],
			handCards[3],
			handCards[4],
			handCards[5],
		}
		stat = p.statistics(tempComb)
		statResult = p.sumStat(statResult, stat)
	}
	return statResult
}

//--- статистика максимальных комбинаций из 7 карт: 2 карты из оставшейся колоды и + 5 карт на руках ---
func (p *Poker) statAllMaxCombHand_5_cards(deckCards, handCards []Card) []int {
	if len(handCards) != 5 {
		fmt.Println("invalid combination handCards")
		return nil
	}
	ch := make(chan []int)
	statResult := make([]int, 11, 11)
	stat := make([]int, 11, 11)

	for i := 0; i < len(deckCards)-1; i++ {
		go func(deckCards, handCards []Card, i int, ch chan []int) {
			statRes := make([]int, 11, 11)
			st := make([]int, 11, 11)
			for j := 1 + i; j < len(deckCards); j++ {
				var tempComb = []Card{
					deckCards[i],
					deckCards[j],
					handCards[0],
					handCards[1],
					handCards[2],
					handCards[3],
					handCards[4],
				}
				st = p.statistics(tempComb)
				statRes = p.sumStat(statRes, st)
			}
			ch <- statRes
		}(deckCards, handCards, i, ch)
	}

	for i := 0; i < len(deckCards)-1; i++ {
		stat = <-ch
		statResult = p.sumStat(statResult, stat)
	}

	return statResult
}

//--- статистика максимальных комбинаций из 7 карт: 5 карт из оставшейся колоды и + 2 карты на руках ---
func (p *Poker) statAllMaxCombHand_2_cards(deckCards, handCards []Card) []int {
	if len(handCards) != 2 {
		fmt.Println("invalid combination handCards")
		return nil
	}

	ch := make(chan []int)
	statResult := make([]int, 11, 11)
	stat := make([]int, 11, 11)

	for i := 0; i < len(deckCards)-4; i++ {
		go func(deckCards, handCards []Card, i int, ch chan []int) {
			len_deckCards := len(deckCards)
			statRes := make([]int, 11, 11)
			st := make([]int, 11, 11)
			for j := 1 + i; j < len_deckCards-3; j++ {
				for k := 1 + j; k < len_deckCards-2; k++ {
					for l := 1 + k; l < len_deckCards-1; l++ {
						for m := 1 + l; m < len_deckCards; m++ {
							var tempComb = []Card{
								deckCards[i],
								deckCards[j],
								deckCards[k],
								deckCards[l],
								deckCards[m],
								handCards[0],
								handCards[1],
							}
							st = p.statistics(tempComb)
							statRes = p.sumStat(statRes, st)
						}
					}
				}
			}
			ch <- statRes
		}(deckCards, handCards, i, ch)
	}

	for i := 0; i < len(deckCards)-4; i++ {
		stat = <-ch
		statResult = p.sumStat(statResult, stat)
	}

	return statResult
}

// --- заполнение массива максимальными комбинациями из 5-ти карт из оставшейся колоды карт (из всевозможных 7ми карточных комбинаций) ---
func (p *Poker) statAllMaxCombOther(deckCards []Card) []int {
	ch := make(chan []int)
	statResult := make([]int, 11, 11)
	stat := make([]int, 11, 11)

	for i := 0; i < len(deckCards)-6; i++ {
		for j := 1 + i; j < len(deckCards)-5; j++ {

			go func(deckCards []Card, i, j int, ch chan []int) {
				len_deckCards := len(deckCards)
				statRes := make([]int, 11, 11)
				st := make([]int, 11, 11)
				for k := 1 + j; k < len_deckCards-4; k++ {
					for l := 1 + k; l < len_deckCards-3; l++ {
						for m := 1 + l; m < len_deckCards-2; m++ {
							for n := 1 + m; n < len_deckCards-1; n++ {
								for r := 1 + n; r < len_deckCards; r++ {
									var tempComb = []Card{
										deckCards[i],
										deckCards[j],
										deckCards[k],
										deckCards[l],
										deckCards[m],
										deckCards[n],
										deckCards[r],
									}
									st = p.statistics(tempComb)
									statRes = p.sumStat(statRes, st)
								}
							}
						}
					}
				}

				ch <- statRes
			}(deckCards, i, j, ch)
		}
	}

	for i := 0; i < len(deckCards)-6; i++ {
		for j := 1 + i; j < len(deckCards)-5; j++ {
			stat = <-ch
			statResult = p.sumStat(statResult, stat)
		}
	}

	return statResult
}

// --- из 7 карт перебираем все комбиназии по 5 карт и сохраняем результат комбинаций ---
func (p *Poker) statistics(tempComb []Card) []int {
	allComb := make([][]Card, 1)
	allComb[0] = make([]Card, 5)
	allTempComb, _ := p.allCombinations_5_Cards(tempComb)
	allComb[0] = p.maxCombOfAllComb5card(allTempComb)
	stat := p.statCombination(allComb)
	return stat
}

// --- суммируем 2 массива stat к statResult ---
func (p *Poker) sumStat(statResult, stat []int) []int {
	statResult[0] += stat[0]
	statResult[1] += stat[1]
	statResult[2] += stat[2]
	statResult[3] += stat[3]
	statResult[4] += stat[4]
	statResult[5] += stat[5]
	statResult[6] += stat[6]
	statResult[7] += stat[7]
	statResult[8] += stat[8]
	statResult[9] += stat[9]
	statResult[10] += stat[10]

	return statResult
}

// --- Подсчет кол-ва комбинаций ---
func (p *Poker) statCombination(allComb_5_Cards [][]Card) []int {
	var rf, sf, k, fh, f, s, c, dp, pr, st int
	n := len(allComb_5_Cards)
	for _, comb := range allComb_5_Cards {
		_, nComb := p.pokerCombination(p.sortComb(comb))
		switch nComb {
		case 1: //"Пара":
			pr++
		case 2: // "Две пары":
			dp++
		case 3: //"Сет":
			c++
		case 4: //"Стрит":
			s++
		case 5: //"Флеш":
			f++
		case 6: //"ФуллХаус":
			fh++
		case 7: //"Каре":
			k++
		case 8: //"СтритФлеш":
			sf++
		case 9: //"РоялФлеш":
			rf++
		default:
			st++
		}
	}
	var resultComb = []int{st, pr, dp, c, s, f, fh, k, sf, rf, n}
	return resultComb
}
