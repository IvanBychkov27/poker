package application

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func (app *Application) handler(rw http.ResponseWriter, req *http.Request) {
	picName := req.URL.Query().Get("pic")
	if picName == "" {
		http.Error(rw, "empty query arg: pic", http.StatusBadRequest)
		return
	}

	pic, err := app.p.GetPicture(picName)
	if err != nil {
		return
	}

	dataBase64 := base64.StdEncoding.EncodeToString(pic)
	im := strings.Replace(image, "{IMAGE}", dataBase64, -1)

	data := pageStart + title + im + pageClose

	rw.Write([]byte(data))
}
