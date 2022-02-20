package httpserver

import (
	"gosupervisor/supervisor/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewMetricRequest struct {
	Name string `json:"name" binding:"required"`
	Unit string `json:"unit" binding:"required"`
}

func (s *HttpServer) NewMetric(c *gin.Context) {
	var request NewMetricRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		ReturnError(http.StatusBadRequest, err, c)
		return
	}

	newmetric := database.Metric{
		Name: request.Name,
		Unit: request.Unit,
	}

	if result := s.Database.Create(&newmetric); result.Error != nil {
		ReturnError(http.StatusInternalServerError, result.Error, c)
		return
	}

	ReturnSuccess(http.StatusOK, newmetric, c)
}

func (s *HttpServer) getAllMetrics(c *gin.Context) {

	var metrics []database.Metric

	if result := s.Database.Find(&metrics); result.Error != nil {
		ReturnError(http.StatusInternalServerError, result.Error, c)
		return
	}

	ReturnSuccess(http.StatusOK, metrics, c)
}
