package application

import (
	"net/http"
	"strings"
)

func (app *Application) poker(wr http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		wr.Write([]byte(pageTop + form + err.Error() + pageBottom))
		return
	}

	if req.Form["resetButton"] != nil {
		wr.Write([]byte(pageTop + app.clearForm(form) + pageBottom))
		return
	}

	nPlayers := req.Form["nPlayers"]
	if len(nPlayers) > 0 {

		resultForm, errPoker := app.p.Poker(req)
		if errPoker != nil {
			data := app.p.SetCheckbox(form)
			wr.Write([]byte(pageTop + app.clearForm(data) + errPoker.Error() + pageBottom))
			return
		}

		data := app.p.SetCheckbox(resultForm)
		wr.Write([]byte(app.clearForm(data) + pageBottom))
	} else {
		data := app.p.SetCheckbox(form)
		wr.Write([]byte(pageTop + app.clearForm(data) + pageBottom))
	}
}

func (app *Application) clearForm(form string) string {
	res := strings.Replace(form, "{head_cards}", " ", 1)
	res = strings.Replace(res, "{table_cards}", " ", 1)
	res = strings.Replace(res, "{released_cards}", " ", 1)
	res = strings.Replace(res, "{head_victory}", " ", 1)
	res = strings.Replace(res, "{table_victory}", " ", 1)
	return res
}
