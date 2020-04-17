// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/16  9:38 下午

package logker

import "os"

type FileLog struct {
	WheError  bool     // whether enable error log file.
	Directory string   // Log save directory
	FileName  string   // Log file name
	File      *os.File // Log file  pointer
	ErrFile   *os.File // Error file log pointer
}

func (f *FileLog) isEnableErr() bool {
	return f.WheError
}

func (f *FileLog) Info(value string, args ...interface{}) {

}
func (f *FileLog) Debug(value string, args ...interface{}) {

}
func (f *FileLog) Error(value string, args ...interface{}) {

}
func (f *FileLog) Warning(value string, args ...interface{}) {

}
