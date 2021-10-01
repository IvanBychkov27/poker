package application

const (
	pageStart = `
<!DOCTYPE html>
<html lang='en'>
    <head>
        <meta charset='utf-8'>
        <title>Poker</title>
        <!-- Ссылка на CSS стили и иконку сайта -->
        <link rel='stylesheet' href='internal/static/css/main.css'>
        <link rel='shortcut icon' href='internal/static/img/favicon.ico' type='image/x-icon'>
        <!-- Подключаем новый шрифт для сайта от Google Fonts -->
        <link rel='stylesheet' href='https://fonts.googleapis.com/css?family=Ubuntu+Mono:400,700'>
    </head>
<body>
`
	pageClose = `
</body>
</htlm>
`

	title = `
	<p><b>Poker</b>
	
`
	image = `
	<p><img src="data:image/jpg; base64,{IMAGE}">
`
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
