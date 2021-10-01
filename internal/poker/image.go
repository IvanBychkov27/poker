package poker

import (
	"encoding/base64"
	"go.uber.org/zap"
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

//--- имена файлов по картам --- вставка картинок
func (p *Poker) cardNameFilesImage(cards []Card) string {
	result := ""
	if len(cards) > 0 {
		cards = p.sortComb(cards)
		for _, c := range cards {
			nameFile := strconv.Itoa(int(c.value)) + strconv.Itoa(int(c.suil))
			result = result + `<img src="data:image/jpg; base64,` + p.imageData(nameFile) + `">`
		}
	}
	return result
}

func (p *Poker) imageData(nameCard string) string {
	openFileName := "cardImage/" + nameCard + ".jpg"
	fileData, err := ioutil.ReadFile(openFileName)
	if err != nil {
		p.logger.Error("error open file", zap.Error(err))
		return openFileName
	}
	return base64.StdEncoding.EncodeToString(fileData)
}

////--- имена файлов по картам ---
//func (p *Poker) cardNameFilesImage_01(cards []Card) string {
//	result := ""
//	if len(cards) > 0 {
//		for _, c := range cards {
//			nameCard := p.parCard(c)
//			nameFile := strconv.Itoa(int(c.value)) + strconv.Itoa(int(c.suil))
//			result = result + `<img src="cardImage/` + nameFile + `.jpg" alt=" ` + nameCard + `" />`
//		}
//	}
//	return result
//}
