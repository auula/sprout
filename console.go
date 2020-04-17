// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/16 - 9:36 下午

package logker

import "fmt"

// console Logger
type console struct {
	// Logging level
	logLevel Level
	// set running Time Zone
	timeZone LogTimeZone
	// customize of time zone type
	tz *TimeZone
	// Whether enable console color
	// code : WheColor bool
}

func (c *console) initTime() {
	// set customize time zone
	c.tz = &TimeZone{TimeZoneStr: c.timeZone}
}

// TODO: Whether enable current level
func (c *console) isEnableLevel() bool {
	return c.logLevel >= DEBUG
}

//func (c *console) IsEnableColor() bool {
//	return c.WheColor
//}

func (c *console) Info(value string, args ...interface{}) {
	if c.isEnableLevel() {
		c.OutPutMessage(INFO, fmt.Sprintf(value, args...))
	}
}
func (c *console) Debug(value string, args ...interface{}) {
	if c.isEnableLevel() {
		c.OutPutMessage(DEBUG, fmt.Sprintf(value, args...))
	}
}
func (c *console) Error(value string, args ...interface{}) {
	if c.isEnableLevel() {
		c.OutPutMessage(ERROR, fmt.Sprintf(value, args...))
	}
}
func (c *console) Warning(value string, args ...interface{}) {
	if c.isEnableLevel() {
		c.OutPutMessage(WARNING, fmt.Sprintf(value, args...))
	}
}
