package main

import (
	"net/http"
	"restful-api/helper"
	"restful-api/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.Auth_Middleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {

	// db := app.Newdb()
	// validator := validator.New()

	// categoryRepository := repository.NewCategoryRepository()
	// categoryService := service.NewCategoryService(categoryRepository, db, validator)
	// categoryController := controller.NewCategoryController(categoryService)

	// router := app.NewRouter(categoryController)
	// authMiddleware := middleware.New_Auth_Middleware(router)
	// server := NewServer(authMiddleware)

	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
