package handlers

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

type HandlerOpts struct {
	Name      string `json:"name"`
	Method    string `json:"method"`
	Params    any    `json:"params"`
	Time      string `json:"time"`
	RemoteIP  string `json:"remote_ip"`
	Host      string `json:"host"`
	UserAgent string `json:"user_agent"`
}

func NewHandlerOpts(c *gin.Context) *HandlerOpts {
	start := time.Now()
	return &HandlerOpts{
		Name:      c.Request.URL.Path,
		Method:    c.Request.Method,
		Params:    nil,
		Time:      start.Format(time.RFC3339Nano),
		RemoteIP:  c.Request.RemoteAddr,
		Host:      c.Request.Host,
		UserAgent: c.Request.UserAgent(),
	}
}

func (h HandlerOpts) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("name", h.Name),
		slog.String("method", h.Method),
		slog.Any("params", h.Params),
		slog.String("time", h.Time),
		slog.String("remote_ip", h.RemoteIP),
		slog.String("host", h.Host),
		slog.String("user_agent", h.UserAgent),
	)
}

type Handler interface {
	Text2Vec(c *gin.Context)
}
