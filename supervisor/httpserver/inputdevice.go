package httpserver

import (
	"gosupervisor/supervisor/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InputDeviceBodyRequest struct {
	InputDeviceName string            `json:"inputdevicename" binding:"required"`
	Metrics         []database.Metric `json:"metrics" binding:"required"`
}

func (server *HttpServer) NewInputDevice(c *gin.Context) {
	var controllerid ControllerURIRequest

	if err := c.ShouldBindUri(&controllerid); err != nil {
		ReturnError(http.StatusBadRequest, err, c)
		return
	}

	var InputDevicebody InputDeviceBodyRequest

	if err := c.ShouldBindJSON(&InputDevicebody); err != nil {
		ReturnError(http.StatusBadRequest, err, c)
		return
	}

	if result := server.Database.Where("controller_id = ?", controllerid.Id).First(&database.Controller{}); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			ReturnError(http.StatusNotFound, result.Error, c)
			return
		}
		ReturnError(http.StatusInternalServerError, result.Error, c)
		return
	}

	newInputDevice := database.InputDevice{
		InputDeviceName: InputDevicebody.InputDeviceName,
		Metrics:         InputDevicebody.Metrics,
		ControllerID:    controllerid.Id,
	}

	if result := server.Database.Create(&newInputDevice); result.Error != nil {
		ReturnError(http.StatusInternalServerError, result.Error, c)
		return
	}

	ReturnSuccess(http.StatusOK, newInputDevice, c)
}

type InputDeviceURIRequest struct {
	ControllerURIRequest
	IdInputDevice string `uri:"id_inputdevice" binding:"required"`
}

func (server *HttpServer) GetInputDeviceByIDAndByController(c *gin.Context) {
	var urirequest InputDeviceURIRequest

	if err := c.ShouldBindUri(&urirequest); err != nil {
		ReturnError(http.StatusBadRequest, err, c)
		return
	}

	var InputDevice database.InputDevice

	if result := server.Database.Where("controller_id = ? AND input_device_id = ?", urirequest.Id, urirequest.IdInputDevice).Preload("Metrics").First(&InputDevice); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			ReturnError(http.StatusNotFound, result.Error, c)
			return
		}
		ReturnError(http.StatusInternalServerError, result.Error, c)
		return
	}

	ReturnSuccess(http.StatusOK, InputDevice, c)

}

func (server *HttpServer) GetAllInputDeviceOfController(c *gin.Context) {
	var requesturi ControllerURIRequest

	if err := c.ShouldBindUri(&requesturi); err != nil {
		ReturnError(http.StatusBadRequest, err, c)
		return
	}

	var InputDevices []database.InputDevice

	if result := server.Database.Debug().Where("controller_id = ?", requesturi.Id).Preload("Metrics").Find(&InputDevices); result.Error != nil {
		ReturnError(http.StatusInternalServerError, result.Error, c)
		return
	}

	if len(InputDevices) == 0 {
		ReturnError(http.StatusNotFound, gorm.ErrRecordNotFound, c)
		return
	}

	ReturnSuccess(http.StatusOK, InputDevices, c)

}

func (server *HttpServer) DeleteInputDevice(c *gin.Context) {
	var request InputDeviceURIRequest

	if err := c.ShouldBindUri(&request); err != nil {
		ReturnError(http.StatusBadRequest, err, c)
		return
	}

	if result := server.Database.Where("controller_id = ? AND input_device_id = ?", request.Id, request.IdInputDevice).Delete(&database.InputDevice{}); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			ReturnError(http.StatusNotFound, result.Error, c)
			return
		}

		ReturnError(http.StatusInternalServerError, result.Error, c)
		return
	}

	ReturnSuccess(http.StatusAccepted, nil, c)
}
