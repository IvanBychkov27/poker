package poker

import (
	"fmt"
)

func (p *Poker) buildResult(cardsHand, cardsTable, outCards []Card, nPlayers int) (result string) {

	comb, nameComb, victory, message := p.buildPokerWin(cardsHand, cardsTable, outCards, nPlayers)

	if len(comb) != 0 {
		result = result + p.cardsDistributed(cardsHand, cardsTable, outCards)
		result = result + "<h2>Ваша комбинация: <B>" + nameComb + "</B><br></h2>"
		result = result + "<B>" + p.cardNameFilesImage(comb) + "</B><br><br>"
		result = result + "Вероятность победы: <B>" + fmt.Sprintf("%2.2f", victory) + " %</B><BR>"
		if victory > 50 {
			result = result + "<B>Рекомендуем увеличить ставку!!!</B><br>"
		}

		result = result + "<br>-------------------------------------------<br>"
		result = result + "Возможные варианты комбинаций у соперников:<br>"
		result = result + "-------------------------------------------<br>"
		result = result + p.statCombOpponents(cardsTable, cardsHand, outCards)
	} else {
		result = result + message
	}

	return result
}

// --- Возможные варианты комбинаций у соперников: ---
func (p *Poker) statCombOpponents(cardTable_5cards, cardHand_2cards, outCards []Card) string {
	outCards = append(outCards, cardHand_2cards...)
	deckCards := p.deckCards_NoFull(outCards)
	statComb := p.statAllMaxCombHand_5_cards(deckCards, cardTable_5cards)
	return p.printStatCombination(statComb)
}
