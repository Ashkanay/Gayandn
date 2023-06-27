package model

import "gorm.io/gorm"

//import h "gayandn/helper"

type User struct {
	//Inheritance
	gorm.Model
	FirstName  string //`json:"firstname"`
	MiddleName string //`json:"middlename"`
	LastName   string //`json:"lastname"`
	LoginName  string //`gorm:"column:loginName; not null"` // json:"loginname"`
	Password   string //`gorm:"not null" json:"password"`
	Gender     string //`gorm:"not null" json:"gender"`
	UserImage  string //`json:"userimage"`
	Mobil      int    //`json:"mobil"`
	Email      string //`json:"email"`
	Note       string //`json:"note"`
}
