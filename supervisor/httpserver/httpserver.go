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

	db.AutoMigrate(&database.Device{})

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
		device := api.Group("/device")
		{
			device.POST("", server.NewDevice)
			device.GET("", server.GetAllDevice)
		}
	}
}

func (server *HttpServer) Run() error {
	server.LoadRoutes()
	return server.Engine.Run(fmt.Sprintf("%s:%d", server.Addr, server.Port))
}
