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

type fileLog struct {
	logLevel  level
	wheError  bool        // Whether enable error log file.
	directory string      // Log save directory
	fileName  string      // Log file name
	file      *os.File    // Log file  pointer
	errFile   *os.File    // Error file log pointer
	tz        *timeZone   // Customize of time zone type
	timeZone  logTimeZone // Set running Time Zone
	power     os.FileMode // File system Power
}

// Initialization error file pointer
func (f *fileLog) initErrPtr() (*os.File, error) {
	savePath := path.Join(f.directory, "error_", f.fileName, ".log")
	file, e := os.OpenFile(savePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, f.power)
	if e == nil {
		return nil, errors.New("open file fail :" + savePath)
	}
	return file, nil
}

//	Initialization file pointer
func (f *fileLog) initFilePtr() (*os.File, error) {
	savePath := path.Join(f.directory, f.fileName, ".log")
	file, e := os.OpenFile(savePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, f.power)
	if e == nil {
		return nil, errors.New("open file fail :" + savePath)
	}
	return file, nil
}

func (f *fileLog) isEnableErr() bool {
	return f.wheError
}

func (f *fileLog) Info(value string, args ...interface{}) {

}
func (f *fileLog) Debug(value string, args ...interface{}) {

}
func (f *fileLog) Error(value string, args ...interface{}) {

}
func (f *fileLog) Warning(value string, args ...interface{}) {

}
