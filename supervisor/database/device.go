package database

import "gorm.io/gorm"

type Device struct {
	gorm.Model
	DeviceID   string
	DeviceName string
}
