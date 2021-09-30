package application

import (
	"encoding/base64"
	"net/http"
	"strings"
)

const pagePoker string = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Poker</title>
	</head>
<body>
	<p><b>Poker</b>
	<p><img src="data:image/jpg; base64,{IMAGE}">
</body>
</html>
`

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
	data := strings.Replace(pagePoker, "{IMAGE}", dataBase64, -1)
	rw.Write([]byte(data))
}
