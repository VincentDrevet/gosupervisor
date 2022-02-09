package supervisor

import (
	"gosupervisor/configuration"
	"gosupervisor/supervisor/httpserver"
	"gosupervisor/supervisor/mqttsupervisor"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Supervisor struct {
	MqttClient mqtt.Client
	HttpServer httpserver.HttpServer
}

func NewSupervisor(configuration configuration.Configuration) Supervisor {

	return Supervisor{
		MqttClient: mqttsupervisor.NewMqttSupervisor(configuration),
		HttpServer: httpserver.NewHttpServer(configuration),
	}
}
