package supervisor

import (
	"gosupervisor/configuration"
	"gosupervisor/supervisor/mqttsupervisor"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Supervisor struct {
	MqttClient mqtt.Client
}

func NewSupervisor(configuration configuration.Configuration) Supervisor {

	return Supervisor{
		MqttClient: mqttsupervisor.NewMqttSupervisor(configuration),
	}
}
