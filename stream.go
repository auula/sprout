// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/16 - 10:59 下午

package logker

// STREAM v1.0.6
// Format console and file log information.
// IO read write operation.
import (
	"github.com/fatih/color"
	"strings"
)

var (
	// Log message format wildcards
	// {level} == logging level
	// {time} == time
	// {position} == runtime caller info
	// {message} == logging message
	// OutPut: [INFO] 2006-01-02 13:05.0006 MP - Position: test.go|main.test:21 - Message: news
	Format = "{level} - Date: {time}  {position} - Message: {message}" //This was modified for v 1.1.4

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
		color.Blue(c.replaceMsg(DEBUG, v) + "\n")
	case INFO.toStr():
		color.Green(c.replaceMsg(INFO, v) + "\n")
	case WARNING.toStr():
		color.Yellow(c.replaceMsg(WARNING, v) + "\n")
	case ERROR.toStr():
		color.Red(c.replaceMsg(ERROR, v) + "\n")
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

	_, err := f.file.WriteString(f.replaceMsg(lev, v, SKIP+2) + "\n")
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
	_, err := f.errFile.WriteString(f.replaceMsg(model, v, SKIP+2) + "\n")
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
// This function was added at 23:50:28 on April 19, 2020 in v1.1.4
func (c *console) replaceMsg(lev level, v string) string {
	c.formatting = strings.Replace(c.formatting, "{level}", "["+lev.toStr()+"]", -1)
	c.formatting = strings.Replace(c.formatting, "{time}", c.tz.NowTimeStr(), -1)
	c.formatting = strings.Replace(c.formatting, "{position}", buildCallerStr(SKIP+1), -1)
	c.formatting = strings.Replace(c.formatting, "{message}", v, -1)
	return c.formatting
}
func (f *fileLog) replaceMsg(lev level, v string, skip int) string {
	f.formatting = strings.Replace(f.formatting, "{level}", "["+lev.toStr()+"]", -1)
	f.formatting = strings.Replace(f.formatting, "{time}", f.tz.NowTimeStr(), -1)
	f.formatting = strings.Replace(f.formatting, "{position}", buildCallerStr(skip), -1)
	f.formatting = strings.Replace(f.formatting, "{message}", v, -1)
	return f.formatting
}

// 第一次写这么整个的开源项目啊哈哈哈哈 解决了issues 有点激动啊...
// ds 牛逼 牛逼牛逼 牛逼 牛逼 我要让这端注释运行在使用这个库的人电脑上或者服务器~~~~
