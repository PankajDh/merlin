package merlin_logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type CustomeErrorFormatter struct {
	NextFormatter log.Formatter
}

func NewCustomErrorFormatter(formatter log.Formatter) *CustomeErrorFormatter {
	return &CustomeErrorFormatter{
		NextFormatter: formatter,
	}
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

type causeChainError interface {
	Cause() error
}

func (f *CustomeErrorFormatter) Format(entry *log.Entry) ([]byte, error) {
	data := make(log.Fields)

	for k, v := range entry.Data {
		switch v.(type) {
		case error:
			data[k] = v.(error).Error()
			// https://github.com/sirupsen/logrus/issues/137
			causeString := ""
			for {
				if stackErr, ok := v.(stackTracer); ok {
					x := ""
					for _, f := range stackErr.StackTrace() {
						y := fmt.Sprintf("%s:%d", f, f)
						x += y
					}

					causeString += x
					if causalError, ok := v.(causeChainError); ok {
						if causalError.Cause() == nil {
							break
						}
						causeString += " ; Caused by "
						v = causalError.Cause()
					} else {
						break
					}
				}
			}
			if causeString != "" {
				data["stacktrace"] = causeString
			}
		default:
			switch val := reflect.ValueOf(v); val.Kind() {
			case reflect.Array, reflect.Struct, reflect.Slice, reflect.Map, reflect.Invalid, reflect.Chan,
				reflect.Uintptr, reflect.Func, reflect.Interface, reflect.Ptr, reflect.UnsafePointer:
				data[k] = fmt.Sprintf("%+v", v)
			default:
				data[k] = v
			}
		}
	}

	nextEntry := *entry
	nextEntry.Data = data

	return f.NextFormatter.Format(&nextEntry)
}

type ContextHook struct{}

func (hook ContextHook) Levels() []log.Level {
	return log.AllLevels
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.WithFields(log.Fields{"name": name, "elapsed": elapsed}).Info(fmt.Sprintf("%s tookTime %s", name, elapsed))
}

type LogCaller struct {
	Func string `json:"func"`
	File string `json:"file"`
	Line int    `json:"line"`
	Ok   bool   `json:"ok"`
}

func GetCaller(level int) LogCaller {
	pc, file, line, ok := runtime.Caller(level)
	return LogCaller{runtime.FuncForPC(pc).Name(), file, line, ok}
}

func Initialise() {
	log.SetReportCaller(true)

	log.SetLevel(log.InfoLevel)
	env := os.Getenv("ENVIRONMENT")

	if env == "PROD" {
		formatter := NewCustomErrorFormatter(&log.JSONFormatter{})
		log.SetFormatter(formatter)

		presentWorkingDir, err := os.Getwd()
		if err != nil {
			log.WithError(err).Errorln("Error while finding the working dir for logging")
			return
		}

		fileName := filepath.Join(presentWorkingDir, "log", "merlin.log")
		_, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_SYNC, 0755)
		if err != nil {
			log.WithError(err).Errorln("Error while logging to file, defaulting to stderr")
			return
		}

		mw := io.MultiWriter(os.Stdout, &lumberjack.Logger{
			Filename: fileName,
			MaxSize:  100,
			MaxAge:   3,
			Compress: false,
		})

		log.SetOutput(mw)
	} else {
		log.SetLevel(log.InfoLevel)
		log.SetOutput(os.Stdout)
		textFormatter := &log.TextFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000",
			FullTimestamp:   true,
			ForceColors:     true,
		}

		formatter := NewCustomErrorFormatter(textFormatter)
		log.SetFormatter(formatter)
	}

	log.WithFields(log.Fields{}).Infoln("Logging Initialised")
}
