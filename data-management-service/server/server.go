package server

type Server interface {
	Start()
	GetRequestArch() *RequestArch
}
