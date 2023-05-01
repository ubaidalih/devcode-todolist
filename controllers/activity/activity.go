package controllers

import (
	"devcode-todolist/entities"
	"devcode-todolist/entities/requests"
	"devcode-todolist/entities/responses"
	"devcode-todolist/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (con *activityController) GetAllActivity(c echo.Context) error {
	db := con.db
	response := entities.Response[[]responses.GetActivityResponse]{}

	activities := []entities.Activity{}
	if err := db.Find(&activities).Error; err != nil {
		response.Status = types.FAILED
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}

	for _, activity := range activities {
		response.Data = append(response.Data, responses.GetActivityResponse{
			ID:        activity.ActivityID,
			Title:     activity.Title,
			Email:     activity.Email,
			CreatedAt: activity.CreatedAt,
			UpdatedAt: activity.UpdatedAt,
		})
	}
	response.Status = types.SUCCESS
	response.Message = types.SUCCESS

	return c.JSON(http.StatusOK, response)
}

func (con *activityController) GetOneActivity(c echo.Context) error {
	db := con.db
	response := entities.Response[responses.GetActivityResponse]{}
	id, _ := strconv.Atoi(c.Param("id"))

	activity := entities.Activity{}
	condition := entities.Activity{ActivityID: id}
	if err := db.Where(&condition).Find(&activity).Error; err != nil {
		response.Status = types.FAILED
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}

	if activity == (entities.Activity{}) {
		response.Status = types.NOT_FOUND
		response.Message = "Activity with ID " + c.Param("id") + " Not Found"
		return c.JSON(http.StatusNotFound, response)
	}

	response.Data = responses.GetActivityResponse{
		ID:        activity.ActivityID,
		Title:     activity.Title,
		Email:     activity.Email,
		CreatedAt: activity.CreatedAt,
		UpdatedAt: activity.UpdatedAt,
	}
	response.Status = types.SUCCESS
	response.Message = types.SUCCESS

	return c.JSON(http.StatusOK, response)
}

func (con *activityController) CreateActivity(c echo.Context) error {
	db := con.db
	response := entities.Response[responses.GetActivityResponse]{}
	request := requests.CreateActivityRequest{}

	if err := c.Bind(&request); err != nil {
		response.Status = types.FAILED
		response.Message = types.ERROR_BAD_REQUEST
		return c.JSON(http.StatusBadRequest, response)
	}

	if request.Title == "" {
		response := entities.Response[[]string]{}
		response.Data = make([]string, 0)
		response.Status = types.ERROR_BAD_REQUEST
		response.Message = "title cannot be null"
		return c.JSON(http.StatusBadRequest, response)
	}

	// insert to database
	activity := entities.Activity{
		Title: request.Title,
		Email: request.Email,
	}

	if err := db.Create(&activity).Error; err != nil {
		response.Status = types.FAILED
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Data = responses.GetActivityResponse{
		ID:        activity.ActivityID,
		Title:     activity.Title,
		Email:     activity.Email,
		CreatedAt: activity.CreatedAt,
		UpdatedAt: activity.UpdatedAt,
	}
	response.Status = types.SUCCESS
	response.Message = types.SUCCESS
	return c.JSON(http.StatusCreated, response)

}

func (con *activityController) UpdateActivity(c echo.Context) error {
	db := con.db
	response := entities.Response[responses.GetActivityResponse]{}
	id, _ := strconv.Atoi(c.Param("id"))
	request := requests.UpdateActivityRequest{}

	if err := c.Bind(&request); err != nil {
		response.Status = types.FAILED
		response.Message = types.ERROR_BAD_REQUEST
		return c.JSON(http.StatusBadRequest, response)
	}

	if request.Title == "" {
		response := entities.Response[[]string]{}
		response.Data = make([]string, 0)
		response.Status = types.ERROR_BAD_REQUEST
		response.Message = "title cannot be null"
		return c.JSON(http.StatusBadRequest, response)
	}

	activity := entities.Activity{}
	condition := entities.Activity{ActivityID: id}
	if err := db.Where(&condition).Find(&activity).Error; err != nil {
		response.Status = types.FAILED
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}

	if activity == (entities.Activity{}) {
		response.Status = types.NOT_FOUND
		response.Message = "Activity with ID " + c.Param("id") + " Not Found"
		return c.JSON(http.StatusNotFound, response)
	}

	if err := db.Model(&activity).Updates(entities.Activity{Title: request.Title}).Error; err != nil {
		response.Status = types.FAILED
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Data = responses.GetActivityResponse{
		ID:        activity.ActivityID,
		Title:     activity.Title,
		Email:     activity.Email,
		CreatedAt: activity.CreatedAt,
		UpdatedAt: activity.UpdatedAt,
	}
	response.Status = types.SUCCESS
	response.Message = types.SUCCESS
	return c.JSON(http.StatusOK, response)

}

func (con *activityController) DeleteActivity(c echo.Context) error {
	db := con.db
	response := entities.Response[entities.Nullstruct]{}
	id, _ := strconv.Atoi(c.Param("id"))

	activity := entities.Activity{}
	condition := entities.Activity{ActivityID: id}
	if err := db.Where(&condition).Find(&activity).Error; err != nil {
		response.Status = types.FAILED
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}

	if activity == (entities.Activity{}) {
		response.Status = types.NOT_FOUND
		response.Message = "Activity with ID " + c.Param("id") + " Not Found"
		return c.JSON(http.StatusNotFound, response)
	}

	if err := db.Delete(&activity).Error; err != nil {
		response.Status = types.FAILED
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Data = entities.Nullstruct{}
	response.Status = types.SUCCESS
	response.Message = types.SUCCESS
	return c.JSON(http.StatusOK, response)
}
