package models

import "time"

type (
	Animal struct {
		Id        int        `gorm:"column:id;primaryKey"`
		Name      string     `gorm:"column:name"`
		Class     string     `gorm:"column:class"`
		Legs      int        `gorm:"column:legs"`
		CreatedAt time.Time  `gorm:"column:created_at"`
		UpdatedAt time.Time  `gorm:"column:updated_at"`
		DeletedAt *time.Time `gorm:"column:deleted_at"`
	}
)

func (Animal) TableName() string {
	return "animal"
}
