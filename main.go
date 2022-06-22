package main

import (
	"Ramdoni007/Learn-Golang-RestFullAPI/app"
	"Ramdoni007/Learn-Golang-RestFullAPI/controller"
	"Ramdoni007/Learn-Golang-RestFullAPI/helper"
	"Ramdoni007/Learn-Golang-RestFullAPI/middleware"
	"Ramdoni007/Learn-Golang-RestFullAPI/repository"
	"Ramdoni007/Learn-Golang-RestFullAPI/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
