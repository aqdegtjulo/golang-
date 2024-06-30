package main

import (
	_ "github.com/gin-gonic/gin"
	"gvb_server/core"
	_ "gvb_server/docs"
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
	"gvb_server/utils"
)

// @title gvb_sever API文档
// @version 1.0
// @description gvb_sever API文档
// @host 127.0.0.01:8080
// @BasePath /
func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	//连接数据库
	global.DB = core.InitGorm()

	core.InitAddrDB()
	defer global.AddrDB.Close()
	//连接redis
	global.Redis = core.ConnectRedis()
	//连接es
	global.ESClient = core.ESConnect()
	//表结构转移数据库
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	//启动gin
	router := routers.InitRouter()
	utils.PrintSystem()
	router.Run(global.Config.System.Addr())
}
