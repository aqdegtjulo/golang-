package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	router.POST("article_create", middleware.JwtAuth(), app.ArticleCreateView)
	router.GET("article_list", middleware.JwtAuth(), app.ArticleListView)
	router.GET("article_detail/:id", app.ArticleDetailView)
	router.GET("article_detail/detail", app.ArticleDetailByTitleView)
	router.GET("article_calendar", app.ArticleCalendarView)
	router.GET("article_tag_view", app.ArticleTagListView)
	router.PUT("article_update", app.ArticleUpdateView)
	router.DELETE("article_remove", app.ArticleRemoveView)
	router.POST("article/collect", middleware.JwtAuth(), app.ArticleCollCreateView)
	router.GET("article_collList_view", middleware.JwtAuth(), app.ArticleCollListView)
	router.DELETE("article_coll_remove", middleware.JwtAuth(), app.ArticleCollBatchRemoveView)
	router.GET("full_text_search", app.FullTextSearchView)
}
