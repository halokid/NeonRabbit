package pkg

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func LoggerInit() {
	log.Println("-->>> Pkg logger init")
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("Jan 02 15:04:05.000000000")
	config.EncoderConfig.StacktraceKey = "" // to hide stacktrace info
	// logLevel := zap.NewAtomicLevel()		// TODO: default log level is `info`
	logLevel := zap.NewAtomicLevelAt(zapcore.WarnLevel) // TODO: set the log level to `warn`
	config.Level = logLevel

	logger, err := config.Build()
	if err != nil {
		log.Fatal(err)
	}

	Logger = logger.Sugar()

	// log := logger.Sugar()
	// log.Info("INFO log level message")
	// log.Warn("Warn log level message")
	// log.Error("Error log level message")
}
