// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/16 - 10:59 下午

package logker

// STREAM v1.0.6
// Format console and file log information.
// IO read write operation.
import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

var (
	//[INFO] 2006-01-02 13:05.0006 MP - Position: test.go|main.test:21 - Message: news
	format     = "{level} - Date: {time}  {position} - Message: {message}" //This version was modified from v 1.1.0
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
		c.replaceMsg(DEBUG, v)
		color.Blue(c.formatting)
	case INFO.toStr():
		c.replaceMsg(DEBUG, v)
		color.Green(c.formatting)
	case WARNING.toStr():
		c.replaceMsg(DEBUG, v)
		color.Yellow(c.formatting)
	case ERROR.toStr():
		c.replaceMsg(DEBUG, v)
		color.Red(c.formatting)
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
	f.replaceMsg(lev, v)
	_, err := f.file.WriteString(f.formatting + "\n")
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

// repair issues : https://github.com/Higker/logker/issues/1
func buildFormat(str string) string {
	match := []string{"{level}", "{time}", "{position}", "{message}"}
	for _, mc := range match {
		if !strings.Contains(string(str), mc) {
			panic("YourLogMessageFormatIsMissing:" + mc + "Tag!!")
		}
	}
	return str
}

// ReplaceOurCustomMessageFormatIdentifier
// This function was added at 23:50:28 on April 19, 2020 in v1.1.0
func (c *console) replaceMsg(lev level, v string) {
	c.formatting = strings.Replace(c.formatting, "{level}", lev.toStr(), -1)
	c.formatting = strings.Replace(c.formatting, "{time}", c.tz.NowTimeStr(), -1)
	c.formatting = strings.Replace(c.formatting, "{position}", buildCallerStr(SKIP), -1)
	c.formatting = strings.Replace(c.formatting, "{message}", v, -1)
}
func (f *fileLog) replaceMsg(lev level, v string) {
	f.formatting = strings.Replace(f.formatting, "{level}", lev.toStr(), -1)
	f.formatting = strings.Replace(f.formatting, "{time}", f.tz.NowTimeStr(), -1)
	f.formatting = strings.Replace(f.formatting, "{position}", buildCallerStr(SKIP), -1)
	f.formatting = strings.Replace(f.formatting, "{message}", v, -1)
}
