package server

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server interface {
	Start()
	GetHandler() *echo.Echo
	GetDB() *gorm.DB
}
