package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InputDevice struct {
	InputDeviceID   uuid.UUID `gorm:"primaryKey"`
	InputDeviceName string
	Metrics         []Metric `gorm:"many2many:inputdevice_metrics;"`
	ControllerID    string
}

func (InputDevice *InputDevice) BeforeCreate(tx *gorm.DB) (err error) {
	InputDevice.InputDeviceID = uuid.New()

	return nil
}
