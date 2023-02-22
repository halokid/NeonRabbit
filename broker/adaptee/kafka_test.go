package adaptee

import (
    "github.com/mattn/go-colorable"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "testing"
)

func TestKafka_Pub(t *testing.T) {
    k := NewKafka()
    k.Pub("neon_rabbit", "test message")

    // --------------------
    z := zap.NewDevelopmentEncoderConfig()
    z.EncodeLevel = zapcore.CapitalColorLevelEncoder
    zc := zap.New(zapcore.NewCore(
        zapcore.NewConsoleEncoder(z),
        zapcore.AddSync(colorable.NewColorableStdout()),
        zapcore.DebugLevel,
    ))

    zc.Warn("-->>> Zap sample test")
}
