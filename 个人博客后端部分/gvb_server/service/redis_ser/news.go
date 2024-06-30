package redis_ser

import (
	"encoding/json"
	"fmt"
	"gvb_server/global"
	"time"
)

const newsIndex = "news_index"

type NewData struct {
	Index    interface{} `json:"index"`
	Title    string      `json:"title"`
	HotValue string      `json:"hotValue"`
	Link     string      `json:"link"`
}

// SetNews 设置某一个数据，重复执行，重复累加
func SetNews(key string, newData []NewData) error {
	byteData, _ := json.Marshal(newData)
	//设置到期时间，每过一小时缓存就到期清理
	err := global.Redis.Set(fmt.Sprintf("%s_%s", newsIndex, key), byteData, 1*time.Hour).Err()
	return err
}

func GetNews(key string) (newData []NewData, err error) {
	res := global.Redis.Get(fmt.Sprintf("%s_%s", newsIndex, key)).Val()
	err = json.Unmarshal([]byte(res), &newData)
	return
}
