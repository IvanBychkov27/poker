package poker

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	constCardsHand  = "hand"
	constCardsTable = "table"
	constCardsOut   = "out"
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
		logger: logger,
	}
	return p
}

func (p *Poker) Poker(req *http.Request) (string, error) {
	cardHand := p.cardsGame(req, constCardsHand)
	cardTable := p.cardsGame(req, constCardsTable)
	cardOut := p.cardsGame(req, constCardsOut)

	nPlayers := req.Form["nPlayers"]
	nPlay, _ := strconv.Atoi(nPlayers[0])
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

// setCheckbox включает флажки на форме в соответствии с выбранными картами
func (p *Poker) SetCheckbox(form string) string {
	for _, card := range p.cardHand {
		suit := "p"
		switch card.suil {
		case 2:
			suit = "k"
		case 3:
			suit = "b"
		case 4:
			suit = "ch"
		}

		cardForm := constCardsHand + strconv.Itoa(int(card.value)) + suit

		if !strings.Contains(form, cardForm) {
			continue
		}
		idx := strings.Index(form, cardForm) + len(cardForm) + 1
		form = form[:idx] + " checked" + form[idx:]
	}

	for _, card := range p.cardTable {
		suit := "p"
		switch card.suil {
		case 2:
			suit = "k"
		case 3:
			suit = "b"
		case 4:
			suit = "ch"
		}

		cardForm := constCardsTable + strconv.Itoa(int(card.value)) + suit

		if !strings.Contains(form, cardForm) {
			continue
		}
		idx := strings.Index(form, cardForm) + len(cardForm) + 1
		form = form[:idx] + " checked" + form[idx:]
	}

	for _, card := range p.cardOut {
		suit := "p"
		switch card.suil {
		case 2:
			suit = "k"
		case 3:
			suit = "b"
		case 4:
			suit = "ch"
		}

		cardForm := constCardsOut + strconv.Itoa(int(card.value)) + suit

		if !strings.Contains(form, cardForm) {
			continue
		}
		idx := strings.Index(form, cardForm) + len(cardForm) + 1
		form = form[:idx] + " checked" + form[idx:]
	}

	players := "nPlayers"
	if strings.Contains(form, players) {
		idx := strings.Index(form, players) + len(players) + 9
		form = form[:idx] + strconv.Itoa(p.nPlay) + `"` + form[idx+2:]
	}

	return form
}
