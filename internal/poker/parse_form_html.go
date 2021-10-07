package poker

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func (p *Poker) ParseFormHTML(c *gin.Context) []byte {
	nPlayers := c.Request.URL.Query().Get("nPlayers")
	if nPlayers == "" {
		return []byte(pageTop + clearForm(form) + pageBottom)
	}

	if len(nPlayers) > 0 {
		resultForm, errPoker := p.Poker(c, nPlayers)
		if errPoker != nil {
			data := p.SetCheckbox(form)
			return []byte(pageTop + clearForm(data) + errPoker.Error() + pageBottom)
		}

		data := p.SetCheckbox(resultForm)
		return []byte(clearForm(data) + pageBottom)
	} else {
		data := p.SetCheckbox(form)
		return []byte(pageTop + clearForm(data) + pageBottom)
	}
}

func clearForm(form string) string {
	res := strings.Replace(form, "{head_cards}", " ", 1)
	res = strings.Replace(res, "{table_cards}", " ", 1)
	res = strings.Replace(res, "{released_cards}", " ", 1)
	res = strings.Replace(res, "{head_victory}", " ", 1)
	res = strings.Replace(res, "{table_victory}", " ", 1)
	return res
}
