package routers

import "gvb_server/api"

func (router RouterGroup) MenuRouter() {
	settingsApi := api.ApiGroupApp.MenuApi
	router.POST("menu_create", settingsApi.MenuCreateView)
	router.GET("menu_list", settingsApi.MenuListView)
	router.GET("menu_names_list", settingsApi.MenuNameList)
	router.PUT("menus_update/:id", settingsApi.MenuUpdateView)
	router.DELETE("menu_delete", settingsApi.MenuRemoveView)
	router.GET("menus_detail/:id", settingsApi.MenuDetailView)
}
