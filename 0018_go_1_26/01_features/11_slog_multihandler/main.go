package main

import (
	"log/slog"
	"os"
)

func main() {
	stdoutHandler := slog.NewTextHandler(os.Stdout, nil)

	file, _ := os.OpenFile("./tmp.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()
	fileHandler := slog.NewJSONHandler(file, nil)

	multiHandler := slog.NewMultiHandler(stdoutHandler, fileHandler)
	logger := slog.New(multiHandler)

	logger.Info("test message",
		slog.Int("id", 42),
	)
}
