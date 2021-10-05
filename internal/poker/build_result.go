package poker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func (p *Poker) buildResult(cardsHand, cardsTable, outCards []Card, nPlayers int) string {
	comb, nameComb, victory, message := p.buildPokerWin(cardsHand, cardsTable, outCards, nPlayers)

	if len(comb) == 0 {
		return message
	}

	resultForm := p.cardsDistributed(cardsHand, cardsTable, outCards)
	myComb := "<h2><B>" + nameComb + "</B><br></h2>"
	myComb += "<B>" + p.cardNameFilesImage(comb) + "</B><br><br>"
	myComb += "Вероятность победы: <B>" + fmt.Sprintf("%2.2f", victory) + " %</B><BR>"

	if victory > 50 {
		myComb += "<B>Рекомендуем увеличить ставку!!!</B><br>"
	}

	resultForm = strings.Replace(resultForm, "{head_victory}", myComb, 1)

	allApp, data := p.statCombOpponents(cardsTable, cardsHand, outCards)
	resultForm = strings.Replace(resultForm, "{table_victory}", allApp, 1)

	dataChart := p.setDataChart(data)
	resultPageTop := strings.Replace(p.PageTop, "{chart_comb}", dataChart, 1)

	return resultPageTop + resultForm
}

// --- Возможные варианты комбинаций у соперников: ---
func (p *Poker) statCombOpponents(cardTable_5cards, cardHand_2cards, outCards []Card) (string, map[string]float64) {
	outCards = append(outCards, cardHand_2cards...)
	outCards = append(outCards, cardTable_5cards...)
	deckCards := p.deckCards_NoFull(outCards)
	statComb := p.statAllMaxCombHand_5_cards(deckCards, cardTable_5cards)
	res, data := p.printStatCombination(statComb)
	return res, data
}

func (p *Poker) setDataChart(data map[string]float64) string {
	type temp struct {
		k string
		v float64
	}
	td := make([]temp, 0, len(data))
	for key, val := range data {
		t := temp{key, val}
		td = append(td, t)
	}

	// сортировка по убыванию
	sort.SliceStable(td, func(i, j int) bool {
		return td[i].v*100 > td[j].v*100
	})

	mas := make([]string, 0, len(data))
	for _, t := range td {
		d := fmt.Sprintf("['%s',%2.2f]", t.k, t.v)
		mas = append(mas, d)
	}

	js, _ := json.Marshal(mas)
	js = bytes.ReplaceAll(js, []byte(`"`), []byte(``))

	return "[['Комбинация', 'Вероятность']," + string(js)[1:]
}
