package poker

import (
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
}

func NewPoker(logger *zap.Logger) *Poker {
	p := &Poker{
		logger: logger,
	}

	return p
}

func (p *Poker) Poker(req *http.Request) string {
	cardHand := p.cardsGame(req, constCardsHand)
	cardTable := p.cardsGame(req, constCardsTable)
	cardOut := p.cardsGame(req, constCardsOut)

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
	return result
}
