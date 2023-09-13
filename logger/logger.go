package logger

import (
	"github.com/rs/zerolog"
	"os"
)

var Log *zerolog.Logger

func Init() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	logger := zerolog.New(consoleWriter).With().Timestamp().Logger()
	Log = &logger
}
