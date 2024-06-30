package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/plugins/log_stash"
)

func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	//连接数据库
	global.DB = core.InitGorm()
	log := log_stash.New("192.168.193.59", "xxxx")
	log.Error(fmt.Sprintf("%s你好啊", "aqdegtjulo"))
}
