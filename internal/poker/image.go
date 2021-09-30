package poker

import (
	"io/ioutil"
	"strconv"
)

func (p *Poker) GetPicture(pic string) ([]byte, error) {
	openFileName := "cardImage/" + pic + ".jpg"
	fileData, err := ioutil.ReadFile(openFileName)
	if err != nil {
		return nil, err
	}

	return fileData, nil
}

//--- имена файлов по картам ---
func (p *Poker) cardNameFilesImage(cards []Card) string {
	result := ""
	if len(cards) > 0 {
		for _, c := range cards {
			nameCard := p.parCard(c)
			nameFile := strconv.Itoa(int(c.value)) + strconv.Itoa(int(c.suil))
			result = result + `<img src="cardImage/` + nameFile + `.jpg" alt=" ` + nameCard + `" />`
		}
	}
	return result
}
