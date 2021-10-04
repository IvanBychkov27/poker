package poker

import (
	"fmt"
	"strings"
)

func (p *Poker) buildResult(cardsHand, cardsTable, outCards []Card, nPlayers int) string {
	comb, nameComb, victory, message := p.buildPokerWin(cardsHand, cardsTable, outCards, nPlayers)

	if len(comb) == 0 {
		return message
	}

	result := p.cardsDistributed(cardsHand, cardsTable, outCards)
	myComb := "<h2><B>" + nameComb + "</B><br></h2>"
	myComb += "<B>" + p.cardNameFilesImage(comb) + "</B><br><br>"
	myComb += "Вероятность победы: <B>" + fmt.Sprintf("%2.2f", victory) + " %</B><BR>"

	if victory > 50 {
		myComb += "<B>Рекомендуем увеличить ставку!!!</B><br>"
	}

	result = strings.Replace(result, "{head_victory}", myComb, 1)

	allApp := p.statCombOpponents(cardsTable, cardsHand, outCards)
	result = strings.Replace(result, "{table_victory}", allApp, 1)

	return result
}

// --- Возможные варианты комбинаций у соперников: ---
func (p *Poker) statCombOpponents(cardTable_5cards, cardHand_2cards, outCards []Card) string {
	outCards = append(outCards, cardHand_2cards...)
	outCards = append(outCards, cardTable_5cards...)
	deckCards := p.deckCards_NoFull(outCards)
	statComb := p.statAllMaxCombHand_5_cards(deckCards, cardTable_5cards)
	return p.printStatCombination(statComb)
}
