package app

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (app *application) InitLogger() error {
	app.log = app.buildLogger()
	return nil
}

func (app *application) buildLogger() *zap.Logger {
	var encoderCfg zapcore.EncoderConfig

	encoderCfg = zap.NewDevelopmentEncoderConfig()

	// Set time encoding to human-readable format
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)

	var logLevel zapcore.Level
	logLevel = zapcore.DebugLevel

	stdout := zapcore.AddSync(os.Stdout)

	cores := make([]zapcore.Core, 0)
	cores = append(cores, zapcore.NewCore(consoleEncoder, stdout, logLevel))

	core := zapcore.NewTee(cores...)

	return zap.New(core)
}
