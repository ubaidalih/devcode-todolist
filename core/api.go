package core

import (
	"devcode-todolist/configs/middlewares"
	"devcode-todolist/routes"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() {
	e := echo.New()

	/* Middlewares */
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.Cors())

	/* Routes */
	routes.ActivityRoute(e)
	routes.TodoRoute(e)

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
