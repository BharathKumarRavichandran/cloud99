package models

import (
	"github.com/jinzhu/gorm"
)

// Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into model `Container`
type Container struct {
	gorm.Model
	ContainerID string  `gorm:"type:varchar(100);NOT NULL"`
	Image       string  `gorm:"type:varchar(100);NOT NULL"`
	DropletID   Droplet `gorm:"foreignkey:DropletID"`
}
