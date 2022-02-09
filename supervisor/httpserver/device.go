package httpserver

import (
	"gosupervisor/supervisor/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	newDevice := database.Device{
		DeviceID:   uuid.NewString(),
		DeviceName: request.DeviceName,
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
