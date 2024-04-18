package logging

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/fatih/color"
)

type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true
	})

	var b []byte
	var err error
	if len(fields) > 0 {
		b, err = json.Marshal(fields)
		if err != nil {
			return err
		}
	}
	// IST is 7 hours ahead of UTC => Thailand
	location := time.FixedZone("IST", 7*60*60)
	timeStr := r.Time.In(location).Format("[2006-01-02 15:04:05]")
	
	msg := color.CyanString(r.Message)

	if len(b) > 0 {
		h.l.Println(timeStr, level, msg, color.WhiteString(string(b)))
	} else {
		h.l.Println(timeStr, level, msg)
	}

	return nil
}

func NewPrettyHandler(
	out io.Writer,
	opts PrettyHandlerOptions,
) *PrettyHandler {
	var sloghandler slog.Handler = slog.NewJSONHandler(out, &opts.SlogOpts)

	h := &PrettyHandler{
		Handler: sloghandler,
		l:       log.New(out, "", 0),
	}

	return h
}

func NewSLogger() {
	opts := PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}
	handler := NewPrettyHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
