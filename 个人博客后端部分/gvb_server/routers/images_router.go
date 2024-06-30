package routers

import "gvb_server/api"

func (router RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi
	router.POST("images_upload", app.ImageUploadView)
	router.GET("images_namelist", app.ImageNameListView)
	router.GET("images_list", app.ImageListView)
	router.DELETE("images_remove", app.ImageRemoveView)
	router.PUT("images_update", app.ImageUpdateView)
}
