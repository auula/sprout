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
	outPutMessage(v string) // Record logging msg.
	outPutErrMessage(v string)
}

func (c *console) outPutMessage(model level, v string) {
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
		c.outPutMessage(model, v)
	}
}

func (f *fileLog) outPutMessage(model level, v string) {
	switch model.toStr() {
	case DEBUG.toStr():
		f.outPut(DEBUG, v)
	case INFO.toStr():
		f.outPut(INFO, v)
	case WARNING.toStr():
		f.outPut(WARNING, v)
	case ERROR.toStr():
		//f.outPutErr(ERROR, v)
		f.outPut(ERROR, v)
	default:
		// Log Level Type Error
		// Program automatically set to debug
		f.logLevel = DEBUG
		f.Error("Log Level Type Error! Program automatically set to debug!!!")
		// recursion
		f.outPutMessage(model, v)
	}
}

func (f *fileLog) outPut(lev level, v string) {
	_, err := f.file.WriteString(fmt.Sprintf(fileFormat, lev.toStr(), f.tz.NowTimeStr(), buildCallerStr(SKIP+1), v))
	_ = f.file.Sync()
	if err != nil {
		_ = f.file.Close()
		panic("output message to log file fail. filePath:" + f.directory + "/" + f.fileName + ".log")
	}
	//fmt.Println(n)
}
func (f *fileLog) outPutErrMessage(model level, v string) {
	switch model.toStr() {
	case DEBUG.toStr():
		f.outPutErr(DEBUG, v)
	case INFO.toStr():
		f.outPutErr(INFO, v)
	case WARNING.toStr():
		f.outPutErr(WARNING, v)
	case ERROR.toStr():
		f.outPutErr(ERROR, v)
	default:
		// Log Level Type Error
		// Program automatically set to debug
		f.logLevel = DEBUG
		f.Error("Log Level Type Error! Program automatically set to debug!!!")
		// recursion
		f.outPutErrMessage(model, v)
	}
}
func (f *fileLog) outPutErr(model level, v string) {
	_, err := f.errFile.WriteString(fmt.Sprintf(fileFormat, model.toStr(), f.tz.NowTimeStr(), buildCallerStr(SKIP+1), v))
	_ = f.errFile.Sync()
	if err != nil {
		_ = f.errFile.Close()
		panic("output message to log file fail. filePath:" + f.directory + "/" + f.fileName + ".log")
	}
}
