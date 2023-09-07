package main

import (
	"github.com/FianGumilar/restful-api-echo/config"
	"github.com/FianGumilar/restful-api-echo/handler"
	"github.com/FianGumilar/restful-api-echo/infrastucture/database"
	"github.com/FianGumilar/restful-api-echo/repository"
	"github.com/FianGumilar/restful-api-echo/service"
	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.NewAppConfig()

	//dbPgConnection := database.GetDbPostgres(conf)

	dbSqlConnection := database.GetSqlConnection(conf)

	// repository
	categoryRepository := repository.NewCategoryRepository(dbSqlConnection)

	// service
	categoryService := service.NewCategoryService(categoryRepository)

	//handler
	categoryHandler := handler.NewCategoryHandler(categoryService)

	app := echo.New()
	api := app.Group("/api")

	api.GET("/categories", categoryHandler.FindAll)
	api.POST("/categories", categoryHandler.Create)
	api.PUT("/categories", categoryHandler.Update)
	api.DELETE("/categories", categoryHandler.Delete)
	api.GET("/categories/:id", categoryHandler.FindByID)

	app.Logger.Fatal(app.Start(":" + conf.Server.Port))
}
