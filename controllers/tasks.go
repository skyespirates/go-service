package controllers

import (
	"github.com/gin-gonic/gin"
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

var tasks = []Task{
	{Id: 1, Title: "Task 1", Description: "Description 1", IsCompleted: false},
	{Id: 2, Title: "Task 2", Description: "Description 2", IsCompleted: true},
	{Id: 3, Title: "Task 3", Description: "Description 3", IsCompleted: false},
}

func GetAllTasks(c *gin.Context) {
	utils.JSONResponse(c, http.StatusOK, "success", tasks)
}

func GetTaskDetail(c *gin.Context) {
	taskId := c.Param("id")
	id, err := strconv.Atoi(taskId)
	// Check if the id is not a number
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "error", nil)
		return
	}

	// Check if the id is out of range
	if id < 0 || id >= len(tasks) {
		utils.JSONResponse(c, http.StatusNotFound, "error", nil)
		return
	}

	task := tasks[id]
	utils.JSONResponse(c, http.StatusOK, "success", task)
}

func CreateTask(c *gin.Context) {
	var newTask Task

	// Bind the JSON request body to the newTask struct
	if err := c.ShouldBindJSON(&newTask); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	// Set the new task ID
	newTask.Id = len(tasks) + 1

	// Append the new task to the tasks slice
	tasks = append(tasks, newTask)

	// Return a success response
	utils.JSONResponse(c, http.StatusCreated, "Task created successfully", newTask)
}

func UpdateTask(c *gin.Context) {
	taskId := c.Param("id")
	id, err := strconv.Atoi(taskId)
	// Check if the id is not a number
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid task ID", nil)
		return
	}

	// Check if the id is out of range
	if id < 1 || id > len(tasks) {
		utils.JSONResponse(c, http.StatusNotFound, "Task not found", nil)
		return
	}

	var updatedTask Task
	// Bind the JSON request body to the updatedTask struct
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	// Update the task fields
	tasks[id-1].Title = updatedTask.Title
	tasks[id-1].Description = updatedTask.Description
	tasks[id-1].IsCompleted = updatedTask.IsCompleted

	// Return a success response
	utils.JSONResponse(c, http.StatusOK, "Task updated successfully", tasks[id-1])
}
func DeleteTask(c *gin.Context) {
	taskId := c.Param("id")
	id, err := strconv.Atoi(taskId)
	// Check if the id is not a number
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid task ID", nil)
		return
	}

	// Check if the id is out of range
	if id < 1 || id > len(tasks) {
		utils.JSONResponse(c, http.StatusNotFound, "Task not found", nil)
		return
	}

	// Remove the task from the slice
	tasks = append(tasks[:id-1], tasks[id:]...)

	// Return a success response
	utils.JSONResponse(c, http.StatusOK, "Task deleted successfully", nil)
}
