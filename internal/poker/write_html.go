package poker

import (
	"fmt"
	"os"
)

// --- Сохранение результата в файл index.html ---
func (p *Poker) saveResultFileHTML(massage string) {
	file, err := os.Create("index.html")
	if err != nil {
		fmt.Println("Ошибка создания файла: ", err)
	}
	defer file.Close()

	text := `<!DOCTYPE HTML>
<html>
    <title>Poker combinations</title>
   <body>
<h2>Покерная комбинация:</h2>`
	text = text + massage
	text = text + `
   </body>
</html>`

	_, _ = file.WriteString(text)
}
