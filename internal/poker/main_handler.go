package poker

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"strconv"
)

func (p *Poker) MainHandler(c *gin.Context) {
	fileName := "index.tmpl.html"
	data := p.ParseFormHTML(c)
	p.saveFileHTML("templates/"+fileName, data)
	c.HTML(200, fileName, nil)
	p.saveFileCountVisits()
}

func (p *Poker) saveFileHTML(fileName string, data []byte) {
	err := ioutil.WriteFile(fileName, data, 0777)
	if err != nil {
		p.logger.Error("error write file ", zap.Error(err))
	}
}

func (p *Poker) saveFileCountVisits() {
	p.countVisits++
	data := []byte(strconv.Itoa(p.countVisits))
	err := ioutil.WriteFile("internal/data/count_visits.txt", data, 0777)
	if err != nil {
		p.logger.Error("error write file ", zap.Error(err))
	}
}
