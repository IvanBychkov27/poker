package poker

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func (p *Poker) ParseFormHTML(c *gin.Context) []byte {
	nPlayers := c.Request.URL.Query().Get("nPlayers")
	if nPlayers == "" {
		return []byte(pageTop + p.clearForm(form+pageBottom))
	}

	if len(nPlayers) > 0 {
		resultForm, errPoker := p.Poker(c, nPlayers)
		if errPoker != nil {
			data := p.SetCheckbox(form)
			return []byte(pageTop + p.clearForm(data) + errPoker.Error() + p.clearForm(pageBottom))
		}

		data := p.SetCheckbox(resultForm)
		return []byte(p.clearForm(data + pageBottom))
	} else {
		data := p.SetCheckbox(form)
		return []byte(pageTop + p.clearForm(data+pageBottom))
	}
}

func (p *Poker) clearForm(form string) string {
	res := strings.Replace(form, "{head_cards}", " ", 1)
	res = strings.Replace(res, "{table_cards}", " ", 1)
	res = strings.Replace(res, "{released_cards}", " ", 1)
	res = strings.Replace(res, "{head_victory}", " ", 1)
	res = strings.Replace(res, "{table_victory}", " ", 1)
	res = strings.Replace(res, "{count_visits}", strconv.Itoa(p.countVisits), 1)
	return res
}
