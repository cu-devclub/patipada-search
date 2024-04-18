package server

type Server interface {
	Start()
	GetGatewayArch() *GatewayArch
}
