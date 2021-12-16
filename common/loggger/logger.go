package loggger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var GLogger *logrus.Logger

func InitLogger() {
	GLogger = logrus.StandardLogger()
	logConfig := viper.GetStringMapString("log")

	name := logConfig["name"]
	level := logConfig["level"]
	tf := logConfig["time_format"]
	file_number := logConfig["file_number"]

	if name == "" {
		name = "log.log"
	}

	if level == "" {
		level = "info"
	}

	if tf == "" {
		tf = "2006-01002 15:04:05.999"
	}

	l := logrus.InfoLevel
	switch strings.ToLower(level) {
	case "trace":
		l = logrus.TraceLevel
	case "debug":
		l = logrus.DebugLevel
	case "info":
		l = logrus.InfoLevel
	case "warn":
		l = logrus.WarnLevel
	case "error":
		l = logrus.ErrorLevel
	case "fatal":
		l = logrus.FatalLevel
	}
	GLogger.SetLevel(l)
	GLogger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: tf,
	})

	if file_number == "true" {
		GLogger.SetReportCaller(true)
	}

	pwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	logName := filepath.Join(pwd, name)

	//w, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModeAppend|os.ModePerm)
	//
	//if err != nil{
	//	panic(err)
	//}

	w, err := rotatelogs.New(
		logName+".%Y%m%d%H%M",
		rotatelogs.WithClock(rotatelogs.UTC),
		// 为最新的日志，建立软连接
		rotatelogs.WithLinkName(logName),
		// 日志分割时间，每24小时分割一次
		rotatelogs.WithRotationTime(24*time.Hour),
		// 日志最长保留时间，30天
		rotatelogs.WithMaxAge(30*24*time.Hour),
		// 日志清理时最大保留个数；和WithMaxAge 只能二选一设置
		//rotatelogs.WithRotationCount(1000),
		// 日志分割的大小, 1M
		rotatelogs.WithRotationSize(8*1024*1024),
	)
	logrus.SetOutput(w)
}
