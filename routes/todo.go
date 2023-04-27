package routes

import (
	"devcode-todolist/configs/database"
	controllers "devcode-todolist/controllers/todo"

	"github.com/labstack/echo/v4"
)

func TodoRoute(e *echo.Echo) {
	todoController := controllers.NewTodoController(database.DB.GetConnection())

	e.GET("/todo-items", todoController.GetAllTodo)
	e.GET("/todo-items/:id", todoController.GetOneTodo)
	e.POST("/todo-items", todoController.CreateTodo)
	e.PATCH("/todo-items/:id", todoController.UpdateTodo)
	e.DELETE("/todo-items/:id", todoController.DeleteTodo)
}
