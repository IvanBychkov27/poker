package poker

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// --- чтение введенных карт из формы---
func (p *Poker) cardsGame(c *gin.Context, formName string) (cards []Card) {
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
			result := c.Request.URL.Query().Get(nForm)
			nForm = formName
			if result == "on" {
				cards = append(cards, Card{byte(i), byte(m)})
			}
		}
	}
	return cards
}
