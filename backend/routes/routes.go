package routes

import (
	"net/http"
	"portal/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// API路由
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/home", handlers.HomeHandler).Methods("GET")
	api.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	api.HandleFunc("/dashboard", handlers.DashboardHandler).Methods("GET")

	// 静态文件服务
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../frontend/build")))

	// 处理前端路由
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/build/index.html")
	})

	return router
}
