package poker

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func (p *Poker) MainHandler(c *gin.Context) {
	fileName := "index.tmpl.html"
	data := p.ParseFormHTML(c)
	p.saveFileHTML("templates/"+fileName, data)
	c.HTML(200, fileName, nil)
}

func (p *Poker) saveFileHTML(fileName string, data []byte) {
	err := ioutil.WriteFile(fileName, data, 0777)
	if err != nil {
		fmt.Println("error write file ", fileName)
	}
}
