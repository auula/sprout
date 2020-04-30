// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/30 - 9:16 下午

package logker

import (
	"github.com/fatih/color"
	"strings"
)

// Channel Buffer Size
const (
	_    = iota
	Qs1w = 10000 * iota
	Qs2w
	Qs3w
	Qs4w
	Qs5w
	Qs6w
)

// send pack msg to log queue
func (c *console) sendMsg(msg *message) {
	c.asyncTask.logQueue <- msg
}

func (c *console) begin() {
	go c.asyncOutPutTask()
}

func (c *console) asyncOutPutTask() {
	for {
		msg := <-c.asyncTask.logQueue
		switch msg.level {
		case DEBUG:
			// blue color of log message.
			// format log message output console.
			color.Blue(c.parsePacket(msg))
		case INFO:
			color.Green(c.parsePacket(msg))
		case WARNING:
			color.Yellow(c.parsePacket(msg))
		case ERROR:
			color.Red(c.parsePacket(msg))
		default:
			// Log Level Type Error
			// Program automatically set to debug
			color.Red("-----------------------------------------------------------------")
			color.Red("！！！Log Level Type Error,Program automatically set to debug！！！")
			color.Red("-----------------------------------------------------------------")
			c.logLevel = DEBUG
			// recursion
			c.sendMsg(c.pack(ERROR,"Log Level Type Error,Program automatically set to debug！！！"))
		}
	}
}

// ReplaceOurCustomMessageFormatIdentifier
// This function was added at 23:50:28 on April 19, 2020 in v1.1.5
func (c *console) parsePacket(msg *message) string {
	str := msg.format
	str = strings.Replace(str, "{level}", msg.level.toStr(), -1)
	str = strings.Replace(str, "{time}", msg.logTime, -1)
	str = strings.Replace(str, "{position}", msg.caller, -1)
	str = strings.Replace(str, "{message}", msg.msg, -1)
	return str + "\n"
}
