// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/16 - 9:36 下午

package logker

import "fmt"

// console Logger
type console struct {
	// Logging level
	logLevel level
	// set running Time Zone
	timeZone logTimeZone
	// customize of time zone type
	tz *timeZone
	// Whether enable console color
	// code : WheColor bool
	// MessageMatchingCard
	formatting string
}

func (c *console) initTime() {
	// set customize time zone
	c.tz = &timeZone{TimeZoneStr: c.timeZone}
}

// TODO: Whether enable current level
func (c *console) isEnableLevel(lev level) bool {
	// debug<info<warn<Error
	return c.logLevel <= lev
}

//func (c *console) IsEnableColor() bool {
//	return c.WheColor
//}

func (c *console) Info(value string, args ...interface{}) {
	if c.isEnableLevel(INFO) {
		c.outPutMessage(INFO, fmt.Sprintf(value, args...))
	}
}
func (c *console) Debug(value string, args ...interface{}) {
	if c.isEnableLevel(DEBUG) {
		c.outPutMessage(DEBUG, fmt.Sprintf(value, args...))
	}
}
func (c *console) Error(value string, args ...interface{}) {
	if c.isEnableLevel(ERROR) {
		c.outPutMessage(ERROR, fmt.Sprintf(value, args...))
	}
}
func (c *console) Warning(value string, args ...interface{}) {
	if c.isEnableLevel(WARNING) {
		c.outPutMessage(WARNING, fmt.Sprintf(value, args...))
	}
}
