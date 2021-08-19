package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Employee struct {
	gorm.Model
	Name   string `gorm:"unique" json:"name"`
	City   string `json:"city"`
	Age    int    `json:"age"`
	Status bool   `json:"status"`
}

func (e *Employee) Disable() {
	e.Status = false
}

func (e *Employee) Enable() {
	e.Status = true
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Employee{})
	return db
}
