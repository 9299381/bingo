package logger

import (
	"errors"
	"github.com/9299381/bingo/package/config"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
	"sync"
	"time"
)

var logger *logrus.Logger
var once sync.Once

func instance() *logrus.Logger {
	once.Do(func() {
		logger = InitLogger(viper.GetString("config.mode"))
	})
	return logger
}
func GetInstance() *logrus.Logger {
	return instance()
}

func InitLogger(mode string) *logrus.Logger {
	var logger = logrus.New()
	if mode == "prod" {
		//写入文件
		src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			panic(err)
		}
		logger.SetOutput(src)
		//设置日志级别
		logger.SetLevel(logrus.InfoLevel)
		logger.AddHook(getLogHook())

	} else {
		src := os.Stdout
		logger.SetOutput(src)
		logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
		logger.SetLevel(logrus.DebugLevel)
	}
	return logger
}

func getLogHook() *lfshook.LfsHook {
	logWriter, _ := getLogWriter()
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	return lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func getLogWriter() (*rotatelogs.RotateLogs, error) {
	logFilePath := config.EnvString("log.file_path", "./logs")
	logFileName := config.EnvString("log.file_name", "log")

	//这里应该放到log中
	exist, err := pathExists(logFilePath)
	if err != nil {
		panic(errors.New("9999::日志目录配置有问题"))
	}
	if !exist {
		err := os.Mkdir(logFilePath, os.ModePerm)
		if err != nil {
			panic(errors.New("9999::创建日志目录失败"))
		}
	}

	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	logWriter, err := rotatelogs.New(
		fileName+".%Y%m%d.log",
		rotatelogs.WithLinkName(fileName),         // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	return logWriter, err
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
