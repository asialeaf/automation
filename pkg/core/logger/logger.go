package logger

import (
	"fmt"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"github.com/asialeaf/automation/pkg/core/common"
)

var Log *AutoLog

type AutoLog struct {
	logrus.FieldLogger
	OutputPath string
	Verbose    bool
}

func NewLogger(outputPath string, verbose bool) *AutoLog {
	logger := logrus.New()

	formatter := &Formatter{
		HideKeys:               true,
		TimestampFormat:        "15:04:05 MST",
		NoColors:               true,
		ShowLevel:              logrus.WarnLevel,
		FieldsDisplayWithOrder: []string{common.Pipeline, common.Module, common.Task, common.Node},
	}
	logger.SetFormatter(formatter)

	path := filepath.Join(outputPath, "./automation.log")
	writer, _ := rotatelogs.New(
		path+".%Y%m%d",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	logWriters := lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}

	if verbose {
		logger.SetLevel(logrus.DebugLevel)
		logWriters[logrus.DebugLevel] = writer
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	logger.Hooks.Add(lfshook.NewHook(logWriters, formatter))
	return &AutoLog{logger, outputPath, verbose}
}

func (k *AutoLog) Message(node, str string) {
	Log.Infof("message: [%s]\n%s", node, str)
}

func (k *AutoLog) Messagef(node, format string, args ...interface{}) {
	Log.Infof("message: [%s]\n%s", node, fmt.Sprintf(format, args...))
}
