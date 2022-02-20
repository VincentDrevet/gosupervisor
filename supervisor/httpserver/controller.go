package httpserver

import (
	"gosupervisor/supervisor/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NewControllerRequest struct {
	ControllerName string `json:"Controllername" binding:"required"`
}

func (server *HttpServer) NewController(c *gin.Context) {
	var request NewControllerRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		ReturnError(http.StatusBadRequest, err, c)
		return
	}

	newController := database.Controller{
		ControllerName: request.ControllerName,
	}

	if result := server.Database.Create(&newController); result.Error != nil {
		ReturnError(http.StatusInternalServerError, result.Error, c)
		return
	}

	ReturnSuccess(http.StatusOK, newController, c)

}

func (server *HttpServer) GetAllController(c *gin.Context) {

	var Controllers []database.Controller

	if result := server.Database.Preload("InputDevices.Metrics").Find(&Controllers); result.Error != nil {
		ReturnError(http.StatusInternalServerError, result.Error, c)
		return
	}

	ReturnSuccess(http.StatusOK, Controllers, c)

}

type ControllerURIRequest struct {
	Id string `uri:"id" binding:"required"`
}

func (server *HttpServer) GetControllerByID(c *gin.Context) {
	var urirequest ControllerURIRequest

	if err := c.ShouldBindUri(&urirequest); err != nil {
		ReturnError(http.StatusBadRequest, err, c)
		return
	}

	var Controller database.Controller

	//if result := server.Database.Select("ControllerID, ControllerName, InputDevices.InputDeviceID, InputDevices.InputDeviceName, InputDevices.Metrics.Name, InputDevices.Metrics.Unit")
	if result := server.Database.Preload("InputDevices.Metrics").Where("Controller_id = ?", urirequest.Id).First(&Controller); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			ReturnError(http.StatusNotFound, result.Error, c)
			return
		}
		ReturnError(http.StatusInternalServerError, result.Error, c)
		return
	}

	ReturnSuccess(http.StatusOK, Controller, c)
}

func (server *HttpServer) DeleteControllerByID(c *gin.Context) {
	var request ControllerURIRequest

	if err := c.ShouldBindUri(&request); err != nil {
		ReturnError(http.StatusBadRequest, err, c)
		return
	}

	if result := server.Database.Where("Controller_id = ?", request.Id).Delete(&database.Controller{}); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			ReturnError(http.StatusNotFound, result.Error, c)
			return
		}
		ReturnError(http.StatusInternalServerError, result.Error, c)
		return
	}

	ReturnSuccess(http.StatusAccepted, nil, c)
}
