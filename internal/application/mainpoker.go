package application

import (
	"encoding/base64"
	"go.uber.org/zap"
	"io/ioutil"
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

func (app *Application) mainpoker(rw http.ResponseWriter, _ *http.Request) {
	pic := app.getPicture()
	if pic == nil {
		return
	}

	dataBase64 := base64.StdEncoding.EncodeToString(pic)
	data := strings.Replace(pagePoker, "{IMAGE}", dataBase64, -1)
	rw.Write([]byte(data))
}

func (app *Application) getPicture() []byte {
	openFileName := "cardImage/144.jpg"
	fileData, err := ioutil.ReadFile(openFileName)
	if err != nil {
		app.logger.Error("error open picture", zap.Error(err))
		return nil
	}

	return fileData
}
