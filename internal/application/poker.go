package application

import (
	"fmt"
	"net/http"
)

func (app *Application) poker(wr http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	_, _ = fmt.Fprint(wr, pageTop, form, err)
	if err != nil {
		_, _ = fmt.Fprintf(wr, anError, err)
	} else {
		nPlayers := req.Form["nPlayers"]
		if len(nPlayers) > 0 {
			result := app.p.Poker(req)
			_, _ = fmt.Fprint(wr, result)
		}
	}
	_, _ = fmt.Fprint(wr, pageBottom)
}
