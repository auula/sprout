// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/30 - 9:13 下午

package logker

import (
	"errors"
	"fmt"
	"strings"
)

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
	asyncTask  *AsyncTask
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

func (c *console) Info(value string, args ...interface{}) {
	if c.isEnableLevel(INFO) {
		c.sendMsg(c.pack(INFO, fmt.Sprintf(value, args...)))
	}
}
func (c *console) Debug(value string, args ...interface{}) {
	if c.isEnableLevel(DEBUG) {
		c.sendMsg(c.pack(DEBUG, fmt.Sprintf(value, args...)))
	}
}
func (c *console) Error(value string, args ...interface{}) {
	if c.isEnableLevel(ERROR) {
		c.sendMsg(c.pack(ERROR, fmt.Sprintf(value, args...)))
	}
}
func (c *console) Warning(value string, args ...interface{}) {
	if c.isEnableLevel(WARNING) {
		c.sendMsg(c.pack(WARNING, fmt.Sprintf(value, args...)))
	}
}

// repair issues : https://github.com/Higker/logker/issues/1
func verify(str string) error {
	match := []string{"{level}", "{time}", "{position}", "{message}"}
	for _, mc := range match {
		if !strings.Contains(string(str), mc) {
			return errors.New("Lack Log format Tag :" + mc + "!!")
		}
	}
	return nil
}
