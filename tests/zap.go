package main

import (
  "encoding/json"
  "log"

  "go.uber.org/zap"
  "go.uber.org/zap/zapcore"
)

func ZapC1() {
  sampleJSON := []byte(`{
       "level" : "info",
       "encoding": "json",
       "outputPaths":["stdout", "log.log"],
       "errorOutputPaths":["stderr"],
       "encoderConfig": {
           "messageKey":"message",
           "levelKey":"level",
           "levelEncoder":"lowercase"
       }
   }`)

  var cfg zap.Config

  if err := json.Unmarshal(sampleJSON, &cfg); err != nil {
    panic(err)
  }

  logger, err := cfg.Build()

  if err != nil {
    panic(err)
  }
  defer logger.Sync()

  logger.Info("INFO log level message")
  logger.Warn("Warn log level message")
  logger.Error("Error log level message")
}

func ZapC2() {
  logger, _ := zap.NewProduction()

  logger.Info("INFO log level message")
  logger.Warn("Warn log level message")
  logger.Error("Error log level message")
}

// use sugaredLogger
func Zapc3() {
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

  log := logger.Sugar()

  log.Info("INFO log level message")
  log.Infof("INFO log level message -->>> %+v, %+v", "1", "2")

  log.Warn("Warn log level message")
  log.Error("Error log level message")
}
