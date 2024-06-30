package routers

import "gvb_server/api"

func (router RouterGroup) TagRouter() {
	app := api.ApiGroupApp.TagApi
	router.POST("tag_create", app.TagCreateView)
	router.GET("tag_list", app.TagListView)
	router.PUT("tag_update/:id", app.TagUpdateView)
	router.DELETE("tag_remove", app.TagRemoveView)
}
