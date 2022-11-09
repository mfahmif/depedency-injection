//go:build wireinject
// +build wireinject

package main

import (
	"net/http"
	"restful-api/app"
	"restful-api/controller"
	"restful-api/middleware"
	"restful-api/repository"
	"restful-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

var routerSet = wire.NewSet(
	app.NewRouter,
	wire.Bind(new(http.Handler), new(*httprouter.Router)),
)

func InitializedServer() *http.Server {
	wire.Build(app.Newdb,
		validator.New,
		categorySet,
		routerSet,
		middleware.New_Auth_Middleware,
		NewServer)
	return nil
}
