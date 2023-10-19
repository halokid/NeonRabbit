package pkg

import (
  "log"

  "go.uber.org/zap"
  "go.uber.org/zap/zapcore"
)

// var Logger *zap.SugaredLogger
type Logger struct {
  L     *zap.SugaredLogger
  Level zapcore.Level
}

func LoggerInit(pkgx *Pkg) {
  log.Println("-->>> Pkg logger init", pkgx)
  //config := zap.NewProductionConfig()
  config := zap.NewDevelopmentConfig()
  config.EncoderConfig.TimeKey = "timestamp"
  //config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("Jan 02 15:04:05.000000000")
  config.EncoderConfig.StacktraceKey = "" // to hide stacktrace info
  // logLevel := zap.NewAtomicLevel()		// TODO: default log level is `info`
  level := zapcore.InfoLevel
  logLevel := zap.NewAtomicLevelAt(level) // TODO: set the log level to `warn`
  config.Level = logLevel
  config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

  logger, err := config.Build()
  if err != nil {
    log.Fatal(err)
  }

  // Logger = logger.Sugar()
  //pkgx.Logger = logger.Sugar()
  pkgx.Logger = &Logger{}
  pkgx.Logger.L = logger.Sugar()
  pkgx.Logger.Level = level

  // log := logger.Sugar()
  // log.Info("INFO log level message")
  // log.Warn("Warn log level message")
  // log.Error("Error log level message")
}

func (l *Logger) Infof(template string, args ...interface{}) {
  l.L.Infof(template, args)
}

func (l *Logger) Info(args ...interface{}) {
  l.L.Info(args)
}

func (l *Logger) CheckErr(err error, template string, args ...interface{}) {
  if l.Level == zapcore.InfoLevel {
    l.L.Infof(template, args)
  } else {
    if err != nil {
      l.L.Errorf(template, args, err)
    }
  }
}
