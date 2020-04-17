// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/16  9:38 下午

package logker

import "os"

type fileLog struct {
	WheError  bool     // whether enable error log file.
	Directory string   // Log save directory
	FileName  string   // Log file name
	File      *os.File // Log file  pointer
	ErrFile   *os.File // Error file log pointer
}

func (f *fileLog) isEnableErr() bool {
	return f.WheError
}

func (f *fileLog) Info(value string, args ...interface{}) {

}
func (f *fileLog) Debug(value string, args ...interface{}) {

}
func (f *fileLog) Error(value string, args ...interface{}) {

}
func (f *fileLog) Warning(value string, args ...interface{}) {

}
