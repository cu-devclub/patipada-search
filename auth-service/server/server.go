package server

import "github.com/labstack/echo/v4"

type Server interface {
	Start()
	GetHandler() *echo.Echo
}
