// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/30 - 9:19 下午

package logker

// This Was Added To v1.2.1
// Log Message Packet
type message struct {
	level   level
	logTime string
	msg     string
	caller  string
	format  string
}

// Package Log Messages
func (c *console) pack(lev level, msg string) *message {
	return &message{level: lev, logTime: c.tz.NowTimeStr(), msg: msg, caller: buildCallerStr(), format: c.formatting}
}

// Package Log Messages
func (f *fileLog) pack(lev level, msg string) *message {
	return &message{level: lev, logTime: f.tz.NowTimeStr(), msg: msg, caller: buildCallerStr(), format: f.formatting}
}
