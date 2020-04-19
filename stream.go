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
		color.Blue(c.formatting, DEBUG.toStr(), c.tz.NowTimeStr(), buildCallerStr(SKIP), v)
	case INFO.toStr():
		color.Green(c.formatting, INFO.toStr(), c.tz.NowTimeStr(), buildCallerStr(SKIP), v)
	case WARNING.toStr():
		color.Yellow(c.formatting, WARNING.toStr(), c.tz.NowTimeStr(), buildCallerStr(SKIP), v)
	case ERROR.toStr():
		color.Red(c.formatting, ERROR.toStr(), c.tz.NowTimeStr(), buildCallerStr(SKIP), v)
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
	_, err := f.file.WriteString(fmt.Sprintf(f.formatting+"\n", lev.toStr(), f.tz.NowTimeStr(), buildCallerStr(SKIP+1), v))
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
	for i, mc := range match {
		//用来检测顺序的
		if !(match[i] == mc) {
			panic("你的匹配符顺序不对！！！正确的顺序是: {level} > {time} > {position} > {message}!!!")
		}
		if strings.Contains(string(str), mc) {
			if mc == match[0] {
				str = strings.Replace(string(str), mc, "[%s]", -1)
			}
			str = strings.Replace(string(str), mc, "%s", -1)
		} else {
			panic("YourLogMessageFormatIsMissing:" + mc + "Tag!!")
		}
	}
	return str
}
