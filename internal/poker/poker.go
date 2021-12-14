// https://git.heroku.com/poker-iv.git
package poker

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	cardsHand  = "hand"
	cardsTable = "table"
	cardsOut   = "out"
)

type Poker struct {
	logger        *zap.Logger
	PageTop       string
	Form          string
	cardHand      []Card
	cardTable     []Card
	cardOut       []Card
	nPlay         int
	statHandMx    sync.RWMutex
	statHand2Card map[string][]int
}

func NewPoker(logger *zap.Logger) *Poker {
	p := &Poker{
		logger:        logger,
		PageTop:       pageTop,
		Form:          form,
		statHand2Card: make(map[string][]int),
	}

	p.getStatDataHand2Card()
	//go p.setHand2Card()

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

	result := p.buildResult(cardHand, cardTable, cardOut, nPlay)

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

func (p *Poker) getStatDataHand2Card() {
	var err error
	var data []byte

	data, err = os.ReadFile("internal/data/stat_all_comb_hand_2_cards.txt")
	if err != nil {
		p.logger.Error("error read file:", zap.Error(err))
	}

	err = json.Unmarshal(data, &p.statHand2Card)
	if err != nil {
		p.logger.Error("error json unmarshal:", zap.Error(err))
	}
	p.logger.Debug("get stat data hand 2 card", zap.Int("count", len(p.statHand2Card)))
}

// заполняем в кеш статистику всех комбинаций из 2-х карт на руках
func (p *Poker) setHand2Card() {
	deckCards := p.deckCardsFull()
	combs, count := p.allCombinations_2_Cards(deckCards)
	p.logger.Debug("set hand 2 card", zap.Int("count", count))

	wg := &sync.WaitGroup{}

	for i, handCards := range combs {
		if i%3 == 0 {
			time.Sleep(time.Second * 10)
		}
		wg.Add(1)
		go func(wg *sync.WaitGroup, cards []Card, id int) {
			defer wg.Done()
			_ = p.statAllMaxCombHand_2_cards(p.deckCards_NoFull(cards), cards)
			p.logger.Debug("go set hand 2 card", zap.Int("id", id))
		}(wg, handCards, i+1)
	}

	wg.Wait()
	p.logger.Debug("set hand 2 card: init all")

	data, err := json.Marshal(p.statHand2Card)
	if err != nil {
		p.logger.Error("error json marshal:", zap.Error(err))
		return
	}
	fileName := "internal/data/stat_all_comb_hand_2_cards.txt"
	err = saveFile(fileName, data)
	if err != nil {
		p.logger.Error("error save file:", zap.Error(err))
	}
}

func saveFile(fileName string, data []byte) error {
	df, errCreateFile := os.Create(fileName)
	if errCreateFile != nil {
		return errCreateFile
	}
	defer df.Close()

	_, errWrite := df.Write(data)
	if errWrite != nil {
		return errWrite
	}
	return nil
}
