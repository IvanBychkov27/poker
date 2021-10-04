package poker

import (
	"strings"
)

// --- Расчет покерной вероятности ---
func (p *Poker) buildPokerWin(cardsHand, cardsTable, outCards []Card, nomberOpponents int) (yourCombCards []Card, nameComb string, victory float64, message string) {
	cardsHandAndTable := cardsHand
	cardsHandAndTable = append(cardsHandAndTable, cardsTable...)
	allOutCardsFromDeck := cardsHandAndTable
	allOutCardsFromDeck = append(allOutCardsFromDeck, outCards...)

	if len(cardsHand) != 2 {
		message = "Введите две Ваши карты!<br>"
		return yourCombCards, nameComb, victory, message
	}
	if message = p.cardsControl(cardsHand, cardsTable, outCards); message != "" {
		return yourCombCards, nameComb, victory, message
	}

	switch len(cardsTable) {
	case 0:
		message += p.cardsDistributed(cardsHand, cardsTable, outCards)
		message += "Ваши возможные варианты комбинаций: <br>"
		dCardNoFull := p.deckCards_NoFull(allOutCardsFromDeck)
		statAllMaxCombHand := p.statAllMaxCombHand_2_cards(dCardNoFull, cardsHandAndTable)
		message += p.printStatCombination(statAllMaxCombHand)
		message += p.recommendations(statAllMaxCombHand)
	case 3:
		message += p.cardsDistributed(cardsHand, cardsTable, outCards)
		message += "Ваши возможные варианты комбинаций: <br>"
		dCardNoFull := p.deckCards_NoFull(allOutCardsFromDeck)
		statAllMaxCombHand := p.statAllMaxCombHand_5_cards(dCardNoFull, cardsHandAndTable)
		message += p.printStatCombination(statAllMaxCombHand)
		message += p.recommendations(statAllMaxCombHand)
	case 4:
		message += p.cardsDistributed(cardsHand, cardsTable, outCards)
		message += "Ваши возможные варианты комбинаций: <br>"
		dCardNoFull := p.deckCards_NoFull(allOutCardsFromDeck)
		statAllMaxCombHand := p.statAllMaxCombHand_6_cards(dCardNoFull, cardsHandAndTable)
		message += p.printStatCombination(statAllMaxCombHand)
		message += p.recommendations(statAllMaxCombHand)
	case 5:
		yourCombCards = p.maxPokerCombinationOf7Cards(cardsHandAndTable)
		yourCombCards = p.sortComb(yourCombCards)
		nameComb, _ = p.pokerCombination(yourCombCards)
		victory = p.percentVictory(cardsTable, cardsHand, outCards)
		victory = p.percentWinsDiffNumbersOpponents(victory, nomberOpponents)
	default:
		message = "Введите все карты выложенные на столе!<br>"
	}
	return yourCombCards, nameComb, victory * 100, message
}

// --- проверка введенных карт на корректность ввода ---
func (p *Poker) cardsControl(cardsHand, cardsTable, outCards []Card) (massage string) {
	for _, v := range cardsTable {
		if cardsHand[0] == v || cardsHand[1] == v {
			massage = "Не корректный ввод карт стола! "
			return massage
		}
	}
	for _, v := range outCards {
		if cardsHand[0] == v || cardsHand[1] == v {
			massage = "Не корректный ввод вышедших карт! "
			return massage
		}
		for _, vt := range cardsTable {
			if vt == v {
				massage = "Не корректный ввод вышедших карт! "
				return massage
			}
		}
	}
	return ""
}

func (p *Poker) recommendations(statInfo []int) string {
	if statInfo[0] < statInfo[10]/2 {
		return "<br><B>Рекомендуем играть!</B><br>"
	}
	return "<br>Далее рекомендуем не играть, но решать тебе!<br>"
}

func (p *Poker) cardsDistributed(cardsHand, cardsTable, outCards []Card) string {
	result := strings.Replace(p.Form, "{head_cards}", p.cardNameFilesImage(cardsHand), 1)
	if len(cardsTable) != 0 {
		result = strings.Replace(result, "{table_cards}", p.cardNameFilesImage(cardsTable), 1)
	}
	if len(outCards) != 0 {
		result = strings.Replace(result, "{released_cards}", p.cardNameFilesImage(outCards), 1)
	}
	return result
}

func (p *Poker) cardsDistributed_01(cardsHand, cardsTable, outCards []Card) string {
	result := "<B>Карты</B><br><br>на руках: "
	result += p.cardNameFilesImage(cardsHand) + "<br><br><br>"
	if len(cardsTable) != 0 {
		result += "на столе: " + p.cardNameFilesImage(cardsTable) + "<br><br><br>"
	}
	if len(outCards) != 0 {
		result += "вышедшие: " + p.cardNameFilesImage(outCards) + "<br><br><br>"
	}
	return result
}

// ---  Расчет процента побед с известной комбинацией карт на руках ---
func (p *Poker) percentVictory(cardTable_5cards, cardHand_2cards, outCards []Card) float64 {
	var victory int
	cards := cardTable_5cards
	cards = append(cards, cardHand_2cards...)
	maxCombHand := p.maxPokerCombinationOf7Cards(cards)
	cards = append(cards, outCards...)
	deckCards := p.deckCards_NoFull(cards)
	nComb := 0
	for i := 0; i < len(deckCards)-1; i++ {
		for j := 1 + i; j < len(deckCards); j++ {
			var tempComb = []Card{
				deckCards[i],
				deckCards[j],
				cardTable_5cards[0],
				cardTable_5cards[1],
				cardTable_5cards[2],
				cardTable_5cards[3],
				cardTable_5cards[4],
			}
			maxTempComb := p.maxPokerCombinationOf7Cards(tempComb)
			rez := p.maxCombination(maxTempComb, maxCombHand)
			if rez != 1 {
				victory++
			}
			nComb++
		}
	}

	return float64(victory) / float64(nComb)
}
