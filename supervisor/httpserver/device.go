package httpserver

import (
	"fmt"
	"gosupervisor/supervisor/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NewDeviceRequest struct {
	DeviceName string `json:"devicename" binding:"required"`
}

func (server *HttpServer) NewDevice(c *gin.Context) {
	var request NewDeviceRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		ReturnError(http.StatusBadRequest, err, c)
		return
	}

	deviceid := uuid.NewString()

	newDevice := database.Device{
		DeviceID:     deviceid,
		DeviceName:   request.DeviceName,
		DataTopic:    fmt.Sprintf("%s/%d", deviceid, 0),
		ErrorTopic:   fmt.Sprintf("%s/%d", deviceid, 1),
		CommandTopic: fmt.Sprintf("%s/%d", deviceid, 2),
	}

	if result := server.Database.Create(&newDevice); result.Error != nil {
		ReturnError(http.StatusInternalServerError, result.Error, c)
		return
	}

	ReturnSuccess(http.StatusOK, newDevice, c)

}

func (server *HttpServer) GetAllDevice(c *gin.Context) {

	var devices []database.Device

	if result := server.Database.Find(&devices); result.Error != nil {
		ReturnError(http.StatusInternalServerError, result.Error, c)
		return
	}

	ReturnSuccess(http.StatusOK, devices, c)

}

type DeviceURIRequest struct {
	Id string `uri:"id" binding:"required"`
}

func (server *HttpServer) GetDeviceByID(c *gin.Context) {
	var urirequest DeviceURIRequest

	if err := c.ShouldBindUri(&urirequest); err != nil {
		ReturnError(http.StatusBadRequest, err, c)
		return
	}

	var device database.Device

	if result := server.Database.Where("device_id = ?", urirequest.Id).First(&device); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			ReturnError(http.StatusNotFound, result.Error, c)
			return
		}
		ReturnError(http.StatusInternalServerError, result.Error, c)
		return
	}

	ReturnSuccess(http.StatusOK, device, c)
}
