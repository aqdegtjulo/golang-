package common

import (
	"fmt"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
)

type Option struct {
	models.PageInfo
	Debug bool
	Likes []string //模糊匹配的字段
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {

	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "created_at desc" // 默认按照时间往前排
	}

	query := DB.Model(&model)
	for index, column := range option.Likes {
		if index == 0 {
			query.Where(fmt.Sprintf("%s like = ?", column), option.Key, fmt.Sprintf("%%%s%%", option.Key))
			continue
		}
		query.Or(fmt.Sprintf("%s like = ?", column), option.Key, fmt.Sprintf("%%%s%%", option.Key))
	}

	count = query.Find(&list).RowsAffected

	//这里的query会受到上面查询的影响，需要手动复位
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	err = DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
}
