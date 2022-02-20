package database

type Metric struct {
	Name string `gorm:"primaryKey"`
	Unit string
}
