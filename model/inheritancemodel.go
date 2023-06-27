package model

import (
	"time"
)

// User struct
type Inheritance struct {
	Id           uint       //`gorm:"not null" json:"id"`
	CreatedDate  time.Time  //`gorm:"not null" json:"created_date"`
	CreatedBy    string     //`json:"created_by"`
	ModifiedDate *time.Time //`json:"modified_date"`
	ModifiedBy   *string    //`json:"modified_by"`
	IsActive     bool       `gorm:"default:true"`
	IsDeleted    bool       `gorm:"default:false"`
	Deleted      bool       `gorm:"default:false"`
}
