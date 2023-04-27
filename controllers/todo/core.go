package controllers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type todoController struct {
	db *gorm.DB
}

type TodoController interface {
	GetAllTodo(c echo.Context) error
	GetOneTodo(c echo.Context) error
	CreateTodo(c echo.Context) error
	UpdateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}

func NewTodoController(db *gorm.DB) TodoController {
	return &todoController{db: db}
}
