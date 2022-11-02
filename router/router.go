package router

import (
	"github.com/ditrit/badaas/controllers"
	"github.com/ditrit/badaas/router/middlewares"
	"github.com/gorilla/mux"
)

// Default router of badaas, initialize all routes.
func SetupRouter(
	//middlewares
	jsonController middlewares.JSONController,
	middlewareLogger middlewares.MiddlewareLogger,

	// controllers
	informationController controllers.InformationController,
) *mux.Router {
	router := mux.NewRouter()
	router.Use(middlewareLogger.Handle)

	router.HandleFunc(
		"/info",
		jsonController.Wrap(informationController.Info),
	).Methods("GET")

	return router
}
