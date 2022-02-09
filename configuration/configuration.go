package configuration

import (
	"encoding/json"
	"io/ioutil"
)

type Configuration struct {
	BrokerAddr             string `json:"broker_addr"`
	BrokerPort             int    `json:"broker_port"`
	MqttSupervisorClientID string `json:"mqtt_supervisor_clientid"`
}

func LoadConfiguration(filepath string) (Configuration, error) {

	filecontent, err := ioutil.ReadFile(filepath)
	if err != nil {
		return Configuration{}, err
	}

	var configuration Configuration

	if err := json.Unmarshal(filecontent, &configuration); err != nil {
		return Configuration{}, err
	}

	return configuration, nil
}
