// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/16 - 10:59 下午

package logker

import (
	"fmt"
	"github.com/fatih/color"
)

const (
	//[INFO] 2006-01-02 13:05.0006 MP - Position: test.go|main.test:21 - Message: news
	format     = "[%s] - Date: %s  %s - Message: %s"
	fileFormat = format + "\n"
)

// Logging record
type logRecord interface {
	OutPutMessage(v string)
}

func (c *console) OutPutMessage(model level, v string) {
	switch model.toStr() {
	case DEBUG.toStr():
		// blue color of log message.
		// format log message output console.
		color.Blue(format, DEBUG.toStr(), c.tz.NowTimeStr(), buildCallerStr(SKIP), v)
	case INFO.toStr():
		color.Green(format, INFO.toStr(), c.tz.NowTimeStr(), buildCallerStr(SKIP), v)
	case WARNING.toStr():
		color.Yellow(format, WARNING.toStr(), c.tz.NowTimeStr(), buildCallerStr(SKIP), v)
	case ERROR.toStr():
		color.Red(format, ERROR.toStr(), c.tz.NowTimeStr(), buildCallerStr(SKIP), v)
	default:
		// Log Level Type Error
		// Program automatically set to debug
		color.Red("-----------------------------------------------------------------")
		color.Red("！！！Log Level Type Error,Program automatically set to debug！！！")
		color.Red("-----------------------------------------------------------------")
		c.logLevel = DEBUG
		// recursion
		c.OutPutMessage(model, v)
	}
}

func (f *fileLog) OutPutMessage(model level, v string) {
	switch model.toStr() {
	case DEBUG.toStr():
		_ = f.outPut(DEBUG, v)
	case INFO.toStr():
		_ = f.outPut(INFO, v)
	case WARNING.toStr():
		_ = f.outPut(WARNING, v)
	case ERROR.toStr():
		_ = f.outPut(ERROR, v)
	default:
		// Log Level Type Error
		// Program automatically set to debug
		f.logLevel = DEBUG
		// recursion
		f.OutPutMessage(model, v)
	}
}

func (f *fileLog) outPut(lev level, v string) error {
	_, err := f.file.WriteString(fmt.Sprintf(fileFormat, lev.toStr(), f.tz.NowTimeStr(), buildCallerStr(SKIP), v))
	if err != nil {
		newErr := f.outPut(lev, v)
		if newErr != nil {
			panic("output message to log file fail. filePath:" + f.directory + "/" + f.fileName)
		}
		return err
	}
	return nil
}
