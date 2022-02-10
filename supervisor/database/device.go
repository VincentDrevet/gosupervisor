package database

type Device struct {
	DeviceID     string
	DeviceName   string
	ErrorTopic   string
	CommandTopic string
	DataTopic    string
}
