package utils

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"vendetta/internal/app/config"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

const (
	levelInfo  = "info"
	levelDebug = "debug"
	levelWarn  = "warn"
	levelError = "error"
	levelFatal = "fatal"
	levelPanic = "panic"
)

type DefaultLogger interface {
	Info(v ...interface{})
	Debug(v ...interface{})
	Warn(v ...interface{})
	Fatal(v ...interface{})
	Panic(v ...interface{})
}

type Logger struct {
	DefaultLogger
	Mode          string
	fileLogger    *zap.Logger
	consoleLogger *log.Logger
}

func (l *Logger) Default() *Logger {
	return l
}

func (l *Logger) Println(v ...interface{}) {
	l.fileLogger.Info(l.format(true, levelInfo, v...))
	l.consoleLogger.Println(l.format(false, levelInfo, v...))
}

func (l *Logger) Debug(v ...interface{}) {
	if l.Mode == config.AppModeDevelopment || l.Mode == config.AppModeDebug {
		l.fileLogger.Debug(l.format(true, levelDebug, v...))
		l.consoleLogger.Println(l.format(false, levelDebug, v...))
	}
}

func (l *Logger) DebugT(trace string, v ...interface{}) {
	if l.Mode == config.AppModeDevelopment || l.Mode == config.AppModeDebug {
		l.fileLogger.Debug(l.formatT(true, "debug", trace, v...))
		l.consoleLogger.Println(l.formatT(false, "debug", trace, v...))
	}
}

func (l *Logger) Info(v ...interface{}) {
	l.fileLogger.Info(l.format(true, levelInfo, v...))
	l.consoleLogger.Println(l.format(false, levelInfo, v...))
}

func (l *Logger) InfoT(trace string, v ...interface{}) {
	l.fileLogger.Info(l.formatT(true, levelInfo, trace, v...))
	l.consoleLogger.Println(l.formatT(false, levelInfo, trace, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.fileLogger.Warn(l.format(true, levelWarn, v...))
	l.consoleLogger.Println(l.format(false, levelWarn, v...))
}

func (l *Logger) WarnT(trace string, v ...interface{}) {
	l.fileLogger.Warn(l.formatT(true, levelWarn, trace, v...))
	l.consoleLogger.Println(l.formatT(false, levelWarn, trace, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.fileLogger.Error(l.format(true, levelError, v...))
	l.consoleLogger.Println(l.format(false, levelError, v...))
}

func (l *Logger) ErrorT(trace string, v ...interface{}) {
	l.fileLogger.Error(l.formatT(true, levelError, trace, v...))
	l.consoleLogger.Println(l.formatT(false, levelError, trace, v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.fileLogger.Fatal(l.format(true, levelFatal, v...))
	l.consoleLogger.Fatal(l.format(false, levelFatal, v...))
}

func (l *Logger) FatalT(trace string, v ...interface{}) {
	l.fileLogger.Fatal(l.formatT(true, levelFatal, trace, v...))
	l.consoleLogger.Fatal(l.formatT(false, levelFatal, trace, v...))
}

func (l *Logger) Panic(v ...interface{}) {
	l.fileLogger.Panic(l.format(true, levelPanic, v...))
	l.consoleLogger.Panicln(l.format(false, levelPanic, v...))
}

func (l *Logger) PanicT(trace string, v ...interface{}) {
	l.fileLogger.Panic(l.formatT(true, levelPanic, trace, v...))
	l.consoleLogger.Panicln(l.formatT(false, levelPanic, trace, v...))
}

func (l *Logger) format(isFile bool, level string, v ...interface{}) string {
	if len(v) == 0 {
		return ""
	}

	str := fmt.Sprintf("%v", v)
	if len(str) == 0 {
		return ""
	}

	str = str[1:]
	str = str[:len(str)-1]

	if isFile {
		return str
	}

	str = Reset + str + Reset
	return fmt.Sprintf("%s %s", l.getLevelTag(level), str)
}

func (l *Logger) formatT(isFile bool, level string, trace string, v ...interface{}) string {
	if len(v) == 0 {
		return ""
	}

	str := fmt.Sprintf("%v", v)
	if len(str) == 0 {
		return ""
	}

	str = str[1:]
	str = str[:len(str)-1]

	if isFile {
		return fmt.Sprintf("%s: %s", trace, str)
	}

	trace = Cyan + trace + Reset
	str = Reset + str + Reset

	return fmt.Sprintf("%s %s: %s", l.getLevelTag(level), trace, str)
}

func (l *Logger) getLevelTag(level string) string {
	var levelTag string
	switch level {
	case levelInfo:
		levelTag = Green + "[INFO]" + Green
	case levelDebug:
		levelTag = Gray + "[DEBUG]" + Gray
	case levelWarn:
		levelTag = Yellow + "[WARN]" + Yellow
	case levelError:
		levelTag = Red + "[ERROR]" + Red
	case levelFatal:
		levelTag = Magenta + "[FATAL]" + Magenta
	case levelPanic:
		levelTag = Red + "[PANIC]" + Red
	}

	return levelTag
}

func NewDefaultLogger(mode string) *Logger {

	level := zap.InfoLevel
	if mode == config.AppModeDebug {
		level = zap.DebugLevel
	}

	pe := zap.NewProductionEncoderConfig()
	fileEncoder := zapcore.NewJSONEncoder(pe)
	pe.EncodeTime = zapcore.ISO8601TimeEncoder

	ioWriter := &lumberjack.Logger{
		Filename:   "logs/logs.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     28,
		LocalTime:  true,
		Compress:   false,
	}

	core := zapcore.NewCore(fileEncoder, zapcore.AddSync(ioWriter), level)
	fileLogger := zap.New(core)

	return &Logger{
		Mode:          mode,
		fileLogger:    fileLogger,
		consoleLogger: log.Default(),
	}
}
