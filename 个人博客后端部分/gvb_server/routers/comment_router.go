package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) CommentRouter() {
	app := api.ApiGroupApp.CommentApi
	router.POST("comment_create", middleware.JwtAuth(), app.CommentCreateView)
	router.GET("comment_list", app.CommentListView)
	router.POST("comment_digg/:id", middleware.JwtAuth(), app.CommentDigg)
	router.DELETE("comment_remove/:id", middleware.JwtAuth(), app.CommentRemoveView)
}
