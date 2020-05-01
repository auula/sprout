// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/1 - 1:04 上午

package logker

import (
	"strings"
)

const (
	// Log file size
	// You can also customize. Unit is kb.
	MB10  int64 = 10 * 1024 * 1024
	MB100 int64 = 100 * 1024 * 1024
	GB1   int64 = 10 * MB100
)

// send pack msg to log queue
func (f *fileLog) sendMsg(msg *message) {
	f.asyncTask.logQueue <- msg
}

// Execute this function to start an asynchronous task
func (f *fileLog) begin() {
	go f.asyncOutPutTask()
}

// Asynchronous Output Logging Tasks
func (f *fileLog) asyncOutPutTask() {
	for {
		msg := <-f.asyncTask.logQueue
		if f.checkSize() {
			// division file
			f.divisionLogFile(plain)
		}
		_, _ = f.file.WriteString(f.parsePacket(msg))
		// 如果是error单独文件开启就执行
		// zh_CN:	检测error独立文件输出开关是否开启 如果开启就往error独立文件输出内容
		// usa_EN:	Check whether the error independent file output switch is turned on.
		// If it is turned on, output the content to the error independent file
		if msg.level == ERROR {
			if f.checkErrSize() {
				// division error file
				f.divisionLogFile(major)
			}
			if f.isEnableErr() {
				_, _ = f.errFile.WriteString(f.parsePacket(msg))
			}
		}
	}
}

// ReplaceOurCustomMessageFormatIdentifier
// This function was added at 23:50:28 on April 19, 2020 in v1.1.5
func (f *fileLog) parsePacket(msg *message) string {
	str := msg.format
	str = strings.Replace(str, "{level}", msg.level.toStr(), -1)
	str = strings.Replace(str, "{time}", msg.logTime, -1)
	str = strings.Replace(str, "{position}", msg.caller, -1)
	str = strings.Replace(str, "{message}", msg.msg, -1)
	return str + "\n"
}
