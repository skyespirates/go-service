package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"rest-api/utils"
	"strconv"
)

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
}

func GetAllTasks(c *gin.Context) {
	var task []utils.Task
	utils.DB.Find(&task)
	utils.JSONResponse(c, http.StatusOK, "success", task)
}

func GetTaskDetail(c *gin.Context) {
	taskId := c.Param("id")

	var task utils.Task
	result := utils.DB.First(&task, taskId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			utils.JSONResponse(c, http.StatusNotFound, "Task not found", nil)
		} else {
			utils.JSONResponse(c, http.StatusInternalServerError, "Error retrieving task", nil)
		}
		return
	}
	utils.JSONResponse(c, http.StatusOK, "success", task)
}

func CreateTask(c *gin.Context) {
	var newTask utils.Task

	// Bind the JSON request body to the newTask struct
	if err := c.ShouldBindJSON(&newTask); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	// Save the new task to the database
	if err := utils.DB.Create(&newTask).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "Error creating task", nil)
		return
	}

	// Return a success response
	utils.JSONResponse(c, http.StatusCreated, "Task created successfully", newTask)
}

func UpdateTask(c *gin.Context) {
	taskId := c.Param("id")
	id, err := strconv.Atoi(taskId)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid task ID", nil)
		return
	}

	var updatedTask utils.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	var task utils.Task
	result := utils.DB.First(&task, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			utils.JSONResponse(c, http.StatusNotFound, "Task not found", nil)
		} else {
			utils.JSONResponse(c, http.StatusInternalServerError, "Error retrieving task", nil)
		}
		return
	}

	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.Completed = updatedTask.Completed

	if err := utils.DB.Save(&task).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "Error updating task", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Task updated successfully", task)
}

func DeleteTask(c *gin.Context) {
	taskId := c.Param("id")
	id, err := strconv.Atoi(taskId)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid task ID", nil)
		return
	}

	var task utils.Task
	result := utils.DB.First(&task, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			utils.JSONResponse(c, http.StatusNotFound, "Task not found", nil)
		} else {
			utils.JSONResponse(c, http.StatusInternalServerError, "Error retrieving task", nil)
		}
		return
	}

	if err := utils.DB.Delete(&task).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "Error deleting task", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Task deleted successfully", nil)
}
