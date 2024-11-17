package logwrapper

import (
	"io"
	"log/slog"
	"os"
)

func Setup(baseFolder string) {
	l, err := os.CreateTemp(baseFolder, "days2xmaslet*.log")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	wh := io.Writer(l)
	slog.SetDefault(slog.New(slog.NewTextHandler(wh, nil)))
}
