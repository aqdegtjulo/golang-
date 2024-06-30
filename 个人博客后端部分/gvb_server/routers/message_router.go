package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) MessageRouter() {
	app := api.ApiGroupApp.MessageApi
	router.POST("message_create", app.MessageCreateView)
	router.GET("message_list_all_view", app.MessageListAllView)
	router.GET("message_list_view", middleware.JwtAuth(), app.MessageListView)
	router.GET("message_record", middleware.JwtAuth(), app.MessageRecordView)
}
