// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/16 - 10:59 下午

package logker

import "github.com/fatih/color"

const (
	//[INFO] 2006-01-02 13:05.0006 MP - Position: test.go|main.test:21 - Message: news
	format = "[%s] - Date: %s  %s - Message: %s"
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
