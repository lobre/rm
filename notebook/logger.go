package notebook

import (
	"bytes"
	"os"

	logger "github.com/apsdehal/go-logger"
)

//const logLevel logger.LogLevel = logger.DebugLevel
const logLevel logger.LogLevel = logger.CriticalLevel

const logFormat string = "▶ %{level} %{message}"

var log *logger.Logger

// Init logger at package initialization
func init() {
	logger.SetDefaultFormat(logFormat)

	var err error
	log, err = logger.New("logger", 1, os.Stdout)
	if err != nil {
		panic(err)
	}
	log.SetLogLevel(logLevel)
}

// Indent text
func indent(indent int, s string) string {
	var buffer bytes.Buffer
	for i := 0; i < indent; i++ {
		buffer.WriteString(" ")
	}
	buffer.WriteString(s)
	return buffer.String()
}
