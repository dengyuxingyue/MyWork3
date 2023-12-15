package api

import (
	"work/pkg/utils"
	"work/service"

	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createTask); err == nil {
		res := createTask.Create(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}

}

func ShowTask(c *gin.Context) {
	var showtask service.ShowTaskService
	if err := c.ShouldBind(&showtask); err == nil {
		res := showtask.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

func ShowAlltasks(c *gin.Context) {
	var showAlltasks service.ShowAllTasksService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showAlltasks); err == nil {
		res := showAlltasks.List(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}

}

func UpdateTask(c *gin.Context) {
	var upDateTask service.UpdateTaskService
	if err := c.ShouldBind(&upDateTask); err == nil {
		res := upDateTask.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

func SearchTask(c *gin.Context) {
	var searchTask service.SearchTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTask); err == nil {
		res := searchTask.Search(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}

}

func DeleteTask(c *gin.Context) {
	var deletetask service.DeleteTaskService
	if err := c.ShouldBind(&deletetask); err == nil {
		res := deletetask.Delete(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}
}

func GetCompletedTasks(c *gin.Context) {
	var getCompletedTasks service.GetCompletedTasksService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&getCompletedTasks); err == nil {
		res := getCompletedTasks.Listcompleted(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}

}

func GetInCompletedTasks(c *gin.Context) {
	var getInCompletedTasks service.GetInCompletedTasksService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&getInCompletedTasks); err == nil {
		res := getInCompletedTasks.ListIncompleted(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error(err)
		c.JSON(400, ErrorResponse(err))
	}

}
