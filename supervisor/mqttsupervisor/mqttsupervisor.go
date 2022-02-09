package mqttsupervisor

import (
	"fmt"
	"gosupervisor/configuration"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func NewMqttSupervisor(configuration configuration.Configuration) mqtt.Client {
	return mqtt.NewClient(mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%d", configuration.BrokerAddr, configuration.BrokerPort)).SetClientID(configuration.MqttSupervisorClientID))
}
