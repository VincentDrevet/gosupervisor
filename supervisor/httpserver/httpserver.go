package httpserver

import (
	"fmt"
	"gosupervisor/configuration"
	"gosupervisor/supervisor/database"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HttpServer struct {
	Engine   *gin.Engine
	Addr     string
	Port     int
	Database *gorm.DB
}

func NewHttpServer(configuration configuration.Configuration) HttpServer {

	db, err := database.InitDatabase(configuration)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	db.AutoMigrate(&database.Controller{}, &database.InputDevice{}, &database.Metric{})

	return HttpServer{
		Engine:   gin.Default(),
		Addr:     configuration.HttpAddr,
		Port:     configuration.HttpPort,
		Database: db,
	}
}

func (server *HttpServer) LoadRoutes() {
	api := server.Engine.Group("/api")
	{
		controller := api.Group("/controller")
		{
			controller.POST("", server.NewController)
			controller.GET("", server.GetAllController)
			controller.GET(":id", server.GetControllerByID)
			controller.DELETE(":id", server.DeleteControllerByID)

			/*
				Gestion périphérique d'entrée
			*/
			controller.POST(":id/inputdevice", server.NewInputDevice)
			controller.GET(":id/inputdevice/:id_inputdevice", server.GetInputDeviceByIDAndByController)
			controller.GET(":id/inputdevice", server.GetAllInputDeviceOfController)
			controller.DELETE(":id/inputdevice/:id_inputdevice", server.DeleteInputDevice)
		}
		metric := api.Group("/metric")
		{
			metric.POST("", server.NewMetric)
			metric.GET("", server.getAllMetrics)
		}
	}
}

func (server *HttpServer) Run() error {
	server.LoadRoutes()
	return server.Engine.Run(fmt.Sprintf("%s:%d", server.Addr, server.Port))
}
