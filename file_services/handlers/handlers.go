package handlers

import "gorm.io/gorm"

//Initializing a handler struct to pass db connection between handler funcs.
//With this handler struct we don't have to open a db connection for every func.
type handler struct {
	DB *gorm.DB
}

func NewConn(db *gorm.DB) handler {
	return handler{db}
}
