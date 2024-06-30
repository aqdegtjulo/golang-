package flag

import "gvb_server/models"

func EsCreateIndex() {
	//models.ArticleModel{}.CreateIndex()
	err := models.FullTextModel{}.CreateIndex()
	if err != nil {
		return
	}
}
