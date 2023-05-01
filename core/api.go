package core

import (
	"devcode-todolist/routes"

	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()

	/* Middlewares */
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	// e.Use(middlewares.Cors())

	/* Routes */
	routes.ActivityRoute(e)
	routes.TodoRoute(e)

	port := "3030"
	go e.Logger.Fatal(e.Start(":" + port))
}
