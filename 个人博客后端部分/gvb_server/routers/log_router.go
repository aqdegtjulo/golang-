package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) LogRouter() {
	app := api.ApiGroupApp.LogApi
	router.GET("log_list", app.LogListView)
	router.DELETE("log_remove", middleware.JwtAuth(), app.LogRemoveListView)
}
