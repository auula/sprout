// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/16  9:38 下午

package logker

import (
	"errors"
	"os"
	"path"
)

//Customize system power
type Power os.FileMode

type fileLog struct {
	WheError  bool        // Whether enable error log file.
	Directory string      // Log save directory
	FileName  string      // Log file name
	File      *os.File    // Log file  pointer
	ErrFile   *os.File    // Error file log pointer
	tz        *timeZone   // Customize of time zone type
	timeZone  logTimeZone // Set running Time Zone
	Pw        Power       // File Power
}

// Initialization error file pointer
func (f *fileLog) initErrPtr() (*os.File, error) {
	savePath := path.Join(f.Directory, "error_", f.FileName, ".log")
	file, e := os.OpenFile(savePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, f.Pw)
	if e == nil {
		return nil, errors.New("open file fail :" + savePath)
	}
	return file, nil
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
