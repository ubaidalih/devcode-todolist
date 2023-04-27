package controllers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type activityController struct {
	db *gorm.DB
}

type ActivityController interface {
	GetAllActivity(c echo.Context) error
	GetOneActivity(c echo.Context) error
	CreateActivity(c echo.Context) error
	UpdateActivity(c echo.Context) error
	DeleteActivity(c echo.Context) error
}

func NewActivityController(db *gorm.DB) ActivityController {
	return &activityController{db: db}
}
