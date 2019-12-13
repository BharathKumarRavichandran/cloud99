package models

// Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into model `Droplet`
type Droplet struct {
	DropletID   uint   `gorm:"index:droplet_id;PRIMARY_KEY"`
	UserID      User   `gorm:"foreignkey:UserID"`
	Ram         string `gorm:"type:varchar(100);NOT NULL"`
	DiskStorage string `gorm:"type:varchar(100);NOT NULL"`
	OS          string `gorm:"type:varchar(100);NOT NULL"`
	Status      string `gorm:"type:varchar(100);NOT NULL"`
}
