package poker

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

const (
	constCardsHand  = "hand"
	constCardsTable = "table"
	constCardsOut   = "out"
)

type Poker struct {
	logger *zap.Logger
	Form   string
}

func NewPoker(logger *zap.Logger) *Poker {
	p := &Poker{
		logger: logger,
	}

	return p
}

func (p *Poker) Poker(req *http.Request) (string, error) {
	cardHand := p.cardsGame(req, constCardsHand)
	cardTable := p.cardsGame(req, constCardsTable)
	cardOut := p.cardsGame(req, constCardsOut)

	if len(cardHand) != 2 {
		return "", fmt.Errorf("error: there are no two cards on hand ")
	}

	nPlayers := req.Form["nPlayers"]
	nPlay, _ := strconv.Atoi(nPlayers[0])
	if nPlay == 0 {
		nPlay = 1
	}

	timeStart := time.Now()
	result := p.buildResult(cardHand, cardTable, cardOut, nPlay)
	timeEnd := time.Now()

	dif := timeEnd.Sub(timeStart)
	if dif.Milliseconds() > 100 {
		p.logger.Debug("poker", zap.String("time", dif.String()))
	}
	//p.saveResultFileHTML(result)
	return result, nil
}
