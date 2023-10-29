package service

import (
	"net/http"
	db "todo/pkg/database"
	"todo/pkg/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RequestResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateTodo(c echo.Context) error {
	task, err := validateAndSanitizeCreateTodo(c)
	if err != nil {
		return c.String(400, err.Error())
	}
	queryResult := db.DB_CONNECTION.GetDB().Create(&task)
	return handleQueryResult(queryResult, c, RequestResponse{Message: "Created Successfully", Data: task})
}

func GetTodo(c echo.Context) error {
	var tasks []models.Task
	status, err := validateAndSanitizeGetTodo(c)
	if err != nil {
		return c.String(400, err.Error())
	}
	var queryResult *gorm.DB
	if status == 0 {
		queryResult = db.DB_CONNECTION.GetDB().Find(&tasks)
	} else {
		queryResult = db.DB_CONNECTION.GetDB().Where("status = ?", status).Find(&tasks)
	}
	if queryResult.Error != nil {
		return c.String(400, queryResult.Error.Error())
	}
	return c.JSON(http.StatusOK, tasks)
}

func DeleteTodo(c echo.Context) error {
	var task models.Task

	id, err := validateAndGetId(c.FormValue("id"))
	if err != nil {
		return c.String(400, err.Error())
	}

	queryResult := db.DB_CONNECTION.GetDB().Where("id = ?", id).Delete(&task)
	return handleQueryResult(queryResult, c, RequestResponse{Message: "Deleted Successfully"})
}

func UpdateTodo(c echo.Context) error {
	id, err := validateAndSanitizeUpdateTodoByID(c)

	if err != nil {
		return c.String(400, err.Error())
	}

	queryResult := db.DB_CONNECTION.GetDB().Model(&models.Task{}).Where("id = ?", id).Updates(c.Get("updateObj").(map[string]interface{}))
	return handleQueryResult(queryResult, c, RequestResponse{Message: "Updated Successfully"})
}

func handleQueryResult(queryResult *gorm.DB, c echo.Context, successMessage RequestResponse) error {
	if queryResult.Error != nil {
		return c.String(400, queryResult.Error.Error())
	}
	if queryResult.RowsAffected == 0 {
		return c.String(400, "No such task")
	}
	return c.JSON(http.StatusOK, successMessage)
}
