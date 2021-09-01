package application

import (
	"net/http"
	"time"
)

const pagePoker string = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<meta http-equiv="refresh" content="10">
		<title>Poker</title>
	</head>
<body>
	<p><b>The Poker project is working </b>
	<p>
`

func (app *Application) mainpoker(rw http.ResponseWriter, _ *http.Request) {
	rw.Write([]byte(pagePoker + time.Now().Format("02-01-2006 15:04:05") + "</body></html>"))
}
