package routers

import "gvb_server/api"

func (router RouterGroup) AdvertRouter() {
	app := api.ApiGroupApp.AdvertApi
	router.POST("adverts_create", app.AdvertCreateView)
	router.GET("adverts_list", app.AdvertListView)
	router.PUT("adverts_update/:id", app.AdvertUpdateView)
	router.DELETE("adverts_remove", app.AdvertRemoveView)
}
