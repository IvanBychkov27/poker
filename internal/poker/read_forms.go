package poker

import (
	"net/http"
	"strconv"
)

// --- чтение введенных карт из формы---
func (p *Poker) cardsGame(req *http.Request, formName string) (cards []Card) {
	nForm := formName
	suit := "p"
	for m := 1; m < 5; m++ {
		switch m {
		case 2:
			suit = "k"
		case 3:
			suit = "b"
		case 4:
			suit = "ch"
		}
		for i := 2; i < 15; i++ {
			nForm += strconv.Itoa(i) + suit
			result := req.Form[nForm]
			nForm = formName
			if len(result) == 1 {
				cards = append(cards, Card{byte(i), byte(m)})
			}
		}
	}
	return cards
}
