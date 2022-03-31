package handlers

import "gorm.io/gorm"

type handler struct {
	DB *gorm.DB
}

func NewConn(db *gorm.DB) handler {
	return handler{db}
}
