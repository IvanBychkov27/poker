// https://git.heroku.com/poker-iv.git
package poker

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

const (
	cardsHand  = "hand"
	cardsTable = "table"
	cardsOut   = "out"
)

type Poker struct {
	logger    *zap.Logger
	PageTop   string
	Form      string
	cardHand  []Card
	cardTable []Card
	cardOut   []Card
	nPlay     int
}

func NewPoker(logger *zap.Logger) *Poker {
	p := &Poker{
		logger:  logger,
		PageTop: pageTop,
		Form:    form,
	}
	return p
}

func (p *Poker) Poker(c *gin.Context, nPlayers string) (string, error) {
	cardHand := p.cardsGame(c, cardsHand)
	cardTable := p.cardsGame(c, cardsTable)
	cardOut := p.cardsGame(c, cardsOut)

	nPlay, _ := strconv.Atoi(nPlayers)
	if nPlay == 0 {
		nPlay = 1
	}

	p.cardHand = cardHand
	p.cardTable = cardTable
	p.cardOut = cardOut
	p.nPlay = nPlay

	if len(cardHand) != 2 {
		return "", fmt.Errorf("Введите две Ваши карты!")
	}

	//timeStart := time.Now()
	result := p.buildResult(cardHand, cardTable, cardOut, nPlay)
	//timeEnd := time.Now()

	//dif := timeEnd.Sub(timeStart)
	//if dif.Milliseconds() > 100 {
	//	p.logger.Debug("poker", zap.String("time", dif.String()))
	//}
	return result, nil
}

// setCheckbox включает флажки на форме в соответствии с выбранными картами
func (p *Poker) SetCheckbox(form string) string {
	res := p.setCheck(form, cardsHand, p.cardHand)
	res = p.setCheck(res, cardsTable, p.cardTable)
	res = p.setCheck(res, cardsOut, p.cardOut)

	players := "nPlayers"
	if strings.Contains(res, players) {
		idx := strings.Index(res, players) + len(players) + 9
		res = res[:idx] + strconv.Itoa(p.nPlay) + `"` + res[idx+2:]
	}
	return res
}

func (p *Poker) setCheck(form, typeCheckbox string, cards []Card) string {
	for _, card := range cards {
		suit := "p"
		switch card.suil {
		case 2:
			suit = "k"
		case 3:
			suit = "b"
		case 4:
			suit = "ch"
		}

		cardForm := typeCheckbox + strconv.Itoa(int(card.value)) + suit

		if !strings.Contains(form, cardForm) {
			continue
		}
		idx := strings.Index(form, cardForm) + len(cardForm) + 1
		form = form[:idx] + " checked" + form[idx:]
	}
	return form
}
