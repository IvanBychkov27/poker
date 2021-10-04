package application

import (
	"fmt"
	"net/http"
)

func (app *Application) poker(wr http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		_, _ = fmt.Fprint(wr, pageTop, form, err, pageBottom)
		return
	}

	nPlayers := req.Form["nPlayers"]
	if len(nPlayers) > 0 {

		resultForm, errPoker := app.p.Poker(req)

		if errPoker != nil {
			_, _ = fmt.Fprint(wr, pageTop, form, pageBottom)
		}
		_, _ = fmt.Fprint(wr, pageTop, resultForm, pageBottom)
	} else {
		_, _ = fmt.Fprint(wr, pageTop, form, pageBottom)
	}

}
