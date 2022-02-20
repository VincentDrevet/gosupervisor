package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Controller struct {
	ControllerID   uuid.UUID `gorm:"primaryKey"`
	ControllerName string
	InputDevices   []InputDevice `gorm:"foreignKey:ControllerID"`
}

func (Controller *Controller) BeforeCreate(tx *gorm.DB) (err error) {
	Controller.ControllerID = uuid.New()

	return nil
}
