package singleton

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LLvlDevelopment = iota + 1
	LLvlProduction
	LStdDir  = "log"        // initial value for log's directory
	LStdFile = "2006-01-02" // log's default filename
)

var (
	instance      *zap.Logger
	sugarInstance *zap.SugaredLogger
	once          sync.Once
	initWithDir   bool // flag identify package declare with dir or not
)

// Options represent option to custom-zap logger
//
// Level set log's level logger, either development or production
// Time set log's time location being used, default is "Asia/Jakarta".
// Use according to Time Zone database, such as "America/New_York".
// Output file is another output file. If you want logger to write log
// to multiple file, add other source here. add "stdout" for console log.
type Options struct {
	Level int
	// Time       *time.Location
	OutputFile []string
}

// newLogger return new custom zap-logger
// set default logger: logs to os.stdout, production level,
// with Asia/Jakarta time. default log filename yyyy-mm-dd.log
func newLoggerWithDir(dir string, prefix string, opt *Options) *zap.Logger {
	var filename = "stdout"

	if opt == nil {
		opt = &Options{}
	}

	if opt.Level < 1 {
		opt.Level = LLvlProduction
	}

	if dir != "" {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Printf("create folder in: %s\n", dir)
			if err = os.MkdirAll(dir, os.ModePerm|os.ModeAppend); err != nil {
				panic(fmt.Sprintf("[log] failed to create directory: %v", err))
			}
		}

		logFile := fmt.Sprintf("%s.%s", time.Now().Format(LStdFile), LStdDir)
		if prefix != "" {
			filename = fmt.Sprintf("%s/%s-%s", dir, prefix, logFile)
		}
		filename = fmt.Sprintf("%s/%s", dir, logFile)
	}

	logger, err := opt.newConfig(filename).Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	return logger
}

// newConfig set config for custom-zap logger
// set log's file to logFile
// set log's time with timeLocation
func (opt *Options) newConfig(logFile string) (cfg zap.Config) {
	if opt.Level > LLvlDevelopment {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.StacktraceKey = ""
	}

	cfg.OutputPaths = []string{logFile}
	cfg.ErrorOutputPaths = []string{logFile}

	if len(opt.OutputFile) > 0 {
		for _, out := range opt.OutputFile {
			cfg.OutputPaths = append(cfg.OutputPaths, out)
			cfg.ErrorOutputPaths = append(cfg.ErrorOutputPaths, out)
		}
	}

	return cfg
}

func InitLoggerWithDir(dir string, prefix string, opt *Options) {
	initWithDir = true
	instance = newLoggerWithDir(dir, prefix, opt)
}

func getLogger() *zap.Logger {
	once.Do(func() {
		if !initWithDir {
			instance = newLogger()
		}
	})

	return instance
}

func GetSugaredLogger() *zap.SugaredLogger {
	if sugarInstance == nil {
		sugarInstance = getLogger().Sugar()
	}
	return sugarInstance
}

func newLogger() *zap.Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.OutputPaths = []string{"stdout"}
	if logger, err := loggerConfig.Build(); err != nil {
		Errorf("failed to create new logger with error: %s", err)
		panic(err)
	} else {
		return logger
	}
}

// Debug logs the message at debug level with additional fields, if any
func Debug(message string, fields ...zap.Field) {
	getLogger().Debug(message, fields...)
}

// Debugf allows Sprintf style formatting and logs at debug level
func Debugf(template string, args ...interface{}) {
	GetSugaredLogger().Debugf(template, args...)
}

// Error logs is equivalent to Error() followed by a call to panic().
func Panic(message string, fields ...zap.Field) {
	getLogger().Error(message, fields...)
	panic(message)
}

// Error logs the message at error level and prints stacktrace with additional fields, if any
func Error(err error, fields ...zap.Field) {
	getLogger().Error(err.Error(), fields...)
}

// Errorf allows Sprintf style formatting, logs at error level and prints stacktrace
func Errorf(template string, args ...interface{}) {
	GetSugaredLogger().Errorf(template, args...)
}

// Fatal logs the message at fatal level with additional fields, if any and exits
func Fatal(err error, fields ...zap.Field) {
	getLogger().Fatal(err.Error(), fields...)
}

// Fatalf allows Sprintf style formatting, logs at fatal level and exits
func Fatalf(template string, args ...interface{}) {
	GetSugaredLogger().Fatalf(template, args...)
}

// Info logs the message at info level with additional fields, if any
func Info(message string, fields ...zap.Field) {
	getLogger().Info(message, fields...)
}

// Infof allows Sprintf style formatting and logs at info level
func Infof(template string, args ...interface{}) {
	GetSugaredLogger().Infof(template, args...)
}

// Warn logs the message at warn level with additional fields, if any
func Warn(message string, fields ...zap.Field) {
	getLogger().Warn(message, fields...)
}

// Warnf allows Sprintf style formatting and logs at warn level
func Warnf(template string, args ...interface{}) {
	GetSugaredLogger().Warnf(template, args...)
}

// AddHook adds func(zapcore.Entry) error) to the logger lifecycle
func AddHook(hook func(zapcore.Entry) error) {
	instance = getLogger().WithOptions(zap.Hooks(hook))
	sugarInstance = instance.Sugar()
}

func WithRequestID(reqID string) *zap.Logger {
	return getLogger().With(
		zap.String("requestID", reqID),
	)
}

// WithRequest takes in a http.Request and logs the message with request's Method, Host and Path
// and returns zap.logger
func WithRequest(r *http.Request) *zap.Logger {
	return getLogger().With(
		zap.Any("method", r.Method),
		zap.Any("host", r.Host),
		zap.Any("path", r.URL.Path),
	)
}

// SugaredWithRequest takes in a http.Request and logs the message with request's Method, Host and Path
// and returns zap.SugaredLogger to support Sprintf styled logging
func SugaredWithRequest(r *http.Request) *zap.SugaredLogger {
	return GetSugaredLogger().With(
		zap.Any("method", r.Method),
		zap.Any("host", r.Host),
		zap.Any("path", r.URL.Path),
	)
}
