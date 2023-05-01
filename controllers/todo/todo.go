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

func (con *todoController) GetAllTodo(c echo.Context) error {
	db := con.db
	response := entities.Response[[]responses.GetTodoResponse]{}

	todos := []entities.Todo{}
	if activity_group_id, err := strconv.Atoi(c.QueryParam("activity_group_id")); err == nil {
		condition := entities.Todo{ActivityGroupID: activity_group_id}
		if err := db.Where(&condition).Find(&todos).Error; err != nil {
			response.Status = types.FAILED
			response.Message = types.ERROR_INTERNAL_SERVER
			return c.JSON(http.StatusInternalServerError, response)
		}
	} else {
		if err := db.Find(&todos).Error; err != nil {
			response.Status = types.FAILED
			response.Message = types.ERROR_INTERNAL_SERVER
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	for _, todo := range todos {
		response.Data = append(response.Data, responses.GetTodoResponse{
			ID:               todo.TodoID,
			ActitvityGroupID: todo.ActivityGroupID,
			Title:            todo.Title,
			IsActive:         todo.IsActive,
			Priority:         todo.Priority,
			CreatedAt:        todo.CreatedAt,
			UpdatedAt:        todo.UpdatedAt,
		})
	}

	response.Status = types.SUCCESS
	response.Message = types.SUCCESS
	return c.JSON(http.StatusOK, response)
}

func (con *todoController) GetOneTodo(c echo.Context) error {
	db := con.db
	response := entities.Response[responses.GetTodoResponse]{}
	id, _ := strconv.Atoi(c.Param("id"))

	todo := entities.Todo{}
	condition := entities.Todo{TodoID: id}
	if err := db.Where(&condition).Find(&todo).Error; err != nil {
		response.Status = types.FAILED
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}

	if todo == (entities.Todo{}) {
		response.Status = types.FAILED
		response.Message = "Todo with ID " + c.Param("id") + " Not Found"
		return c.JSON(http.StatusNotFound, response)
	}

	response.Data = responses.GetTodoResponse{
		ID:               todo.TodoID,
		ActitvityGroupID: todo.ActivityGroupID,
		Title:            todo.Title,
		IsActive:         todo.IsActive,
		Priority:         todo.Priority,
		CreatedAt:        todo.CreatedAt,
		UpdatedAt:        todo.UpdatedAt,
	}

	response.Status = types.SUCCESS
	response.Message = types.SUCCESS
	return c.JSON(http.StatusOK, response)
}

func (con *todoController) CreateTodo(c echo.Context) error {
	db := con.db
	response := entities.Response[responses.GetTodoResponse]{}
	request := requests.CreateTodoRequest{}

	if err := c.Bind(&request); err != nil {
		response.Status = types.FAILED
		response.Message = types.ERROR_BAD_REQUEST
		return c.JSON(http.StatusBadRequest, response)
	}

	if request.Title == "" {
		response.Status = types.ERROR_BAD_REQUEST
		response.Message = "title cannot be null"
		response.Data = responses.GetTodoResponse{}
		return c.JSON(http.StatusBadRequest, response)
	}

	if request.ActivityGroupID == 0 {
		response.Status = types.ERROR_BAD_REQUEST
		response.Message = "activity_group_id cannot be null"
		response.Data = responses.GetTodoResponse{}
		return c.JSON(http.StatusBadRequest, response)
	}

	todo := entities.Todo{
		Title:           request.Title,
		ActivityGroupID: request.ActivityGroupID,
		IsActive:        request.IsActive,
		Priority:        request.Priority,
	}

	// if !request.IsActive {
	// 	todo.IsActive = true
	// }

	if request.Priority == "" {
		todo.Priority = "very-high"
	}

	if err := db.Create(&todo).Error; err != nil {
		response.Status = types.FAILED
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Data = responses.GetTodoResponse{
		ID:               todo.TodoID,
		ActitvityGroupID: todo.ActivityGroupID,
		Title:            todo.Title,
		IsActive:         todo.IsActive,
		Priority:         todo.Priority,
		CreatedAt:        todo.CreatedAt,
		UpdatedAt:        todo.UpdatedAt,
	}
	response.Status = types.SUCCESS
	response.Message = types.SUCCESS
	return c.JSON(http.StatusOK, response)
}

func (con *todoController) UpdateTodo(c echo.Context) error {
	db := con.db
	response := entities.Response[responses.GetTodoResponse]{}
	id, _ := strconv.Atoi(c.Param("id"))
	request := requests.UpdateTodoRequest{}

	if err := c.Bind(&request); err != nil {
		response.Status = types.FAILED
		response.Message = types.ERROR_BAD_REQUEST
		return c.JSON(http.StatusBadRequest, response)
	}

	todo := entities.Todo{}
	condition := entities.Todo{TodoID: id}
	if err := db.Where(&condition).Find(&todo).Error; err != nil {
		response.Status = types.FAILED
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}

	if todo == (entities.Todo{}) {
		response.Status = types.FAILED
		response.Message = "Todo with ID " + c.Param("id") + " Not Found"
		return c.JSON(http.StatusNotFound, response)
	}

	if err := db.Model(&todo).Updates(entities.Todo{Title: request.Title, Priority: request.Priority, IsActive: request.IsActive}).Error; err != nil {
		response.Status = types.FAILED
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Data = responses.GetTodoResponse{
		ID:               todo.TodoID,
		ActitvityGroupID: todo.ActivityGroupID,
		Title:            todo.Title,
		IsActive:         todo.IsActive,
		Priority:         todo.Priority,
		CreatedAt:        todo.CreatedAt,
		UpdatedAt:        todo.UpdatedAt,
	}
	response.Status = types.SUCCESS
	response.Message = types.SUCCESS
	return c.JSON(http.StatusOK, response)
}

func (con *todoController) DeleteTodo(c echo.Context) error {
	db := con.db
	response := entities.Response[responses.GetTodoResponse]{}
	id, _ := strconv.Atoi(c.Param("id"))

	todo := entities.Todo{}
	condition := entities.Todo{TodoID: id}
	if err := db.Where(&condition).Find(&todo).Error; err != nil {
		response.Status = types.FAILED
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}

	if todo == (entities.Todo{}) {
		response.Status = types.FAILED
		response.Message = "Todo with ID " + c.Param("id") + " Not Found"
		return c.JSON(http.StatusNotFound, response)
	}

	if err := db.Delete(&todo).Error; err != nil {
		response.Status = types.FAILED
		response.Message = types.ERROR_INTERNAL_SERVER
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Status = types.SUCCESS
	response.Message = types.SUCCESS
	return c.JSON(http.StatusOK, response)
}
