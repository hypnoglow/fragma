package slogfu

import "log/slog"

func Named(logger *slog.Logger, name string) *slog.Logger {
	return logger.With(slog.String("logger", name))
}

func Error(err error) slog.Attr {
	if err == nil {
		return slog.String("error", "<nil>")
	}
	return slog.String("error", err.Error())
}
