// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/16  9:38 下午

package logker

import (
	"errors"
	"fmt"
	"os"
	"path"
)

const (
	errPerfix = "error_"
	suffix    = ".log"
	bakSuffix = "_bak.log"
	bakPerfix = "log_"
)

type logFileType int

const (
	plain logFileType = iota // Common plain Logging file type
	major                    // Major error logging file type
)

type fileLog struct {
	logLevel    level
	wheError    bool        // Whether enable error log file.
	directory   string      // Log save directory
	fileName    string      // Log file name
	file        *os.File    // Log file  pointer
	errFile     *os.File    // Error file log pointer
	tz          *timeZone   // Customize of time zone type
	timeZone    logTimeZone // Set running Time Zone
	power       os.FileMode // File system Power
	fileMaxSize int64       // File Max size
}

// Initialization error file pointer
func (f *fileLog) initErrPtr() (*os.File, error) {
	savePath := path.Join(f.directory, errPerfix+f.fileName+suffix)
	file, e := os.OpenFile(savePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, f.power)
	if e == nil {
		return nil, errors.New("open file fail :" + savePath)
	}
	return file, nil
}

//	Initialization file pointer
func (f *fileLog) initFilePtr() (*os.File, error) {
	savePath := path.Join(f.directory, f.fileName+".log")
	// fmt.Println(savePath)
	file, e := os.OpenFile(savePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, f.power)
	// fmt.Printf("3 %T %p \n", file, file)
	if e != nil {
		return nil, errors.New("open file fail :" + savePath)
	}
	return file, nil
}

func (f *fileLog) isEnableErr() bool {
	return f.wheError
}

// TODO: Whether enable current level
func (f *fileLog) isEnableLevel(lev level) bool {
	// debug<info<warn<Error
	return f.logLevel <= lev
}

// TODO: Check log size & whether division log file
func (f *fileLog) checkSize() bool {
	info, e := f.file.Stat()
	if e != nil {
		return false
	}
	return info.Size() >= f.fileMaxSize
}
func (f *fileLog) checkErrSize() bool {
	info, e := f.errFile.Stat()
	if e != nil {
		return false
	}
	return info.Size() >= f.fileMaxSize
}
func (f *fileLog) Info(value string, args ...interface{}) {
	if f.isEnableLevel(INFO) {
		if f.checkSize() {
			// division file
			f.divisionLogFile(plain)
		}
		f.outPutMessage(INFO, fmt.Sprintf(value, args...))
	}
}
func (f *fileLog) Debug(value string, args ...interface{}) {
	if f.isEnableLevel(DEBUG) {
		if f.checkSize() {
			// division file
			f.divisionLogFile(plain)
		}
		f.outPutMessage(DEBUG, fmt.Sprintf(value, args...))
	}
}
func (f *fileLog) Error(value string, args ...interface{}) {
	if f.isEnableLevel(ERROR) {
		if f.checkSize() {
			// division file
			f.divisionLogFile(plain)
		}
		// 检测error独立文件输出开关是否开启 如果开启就往error独立文件输出内容
		if f.isEnableErr() {
			fmt.Println("1进来了", f.wheError)
			if f.checkErrSize() {
				fmt.Println("2进来了", f.wheError)
				// division error file
				f.divisionLogFile(major)
			}
			fmt.Println("4进来了", f.wheError)
			f.outPutErrMessage(ERROR, fmt.Sprintf(value, args...))
		}
		f.outPutMessage(ERROR, fmt.Sprintf(value, args...))
	}
}
func (f *fileLog) Warning(value string, args ...interface{}) {
	if f.isEnableLevel(WARNING) {
		if f.checkSize() {
			// division file
			f.divisionLogFile(plain)
		}

		f.outPutMessage(WARNING, fmt.Sprintf(value, args...))
	}
}

// Division logging file.
func (f *fileLog) divisionLogFile(fileType logFileType) {
	switch fileType {
	case plain:
		_ = f.file.Close()
		srcPath := path.Join(f.directory, f.fileName+suffix)
		newPath := path.Join(f.directory, bakPerfix+f.tz.NowTimeStrLogName()+bakSuffix)
		_ = os.Rename(srcPath, newPath)
		f.file, _ = f.initFilePtr()
	case major:
		fmt.Println("3进来了", f.wheError)
		_ = f.errFile.Close()
		srcPath := path.Join(f.directory, errPerfix+f.fileName+suffix)
		newPath := path.Join(f.directory, errPerfix+f.tz.NowTimeStrLogName()+bakSuffix)
		_ = os.Rename(srcPath, newPath)
		f.errFile, _ = f.initErrPtr()
	default:
		f.Error("division log file fail. Type: %v", fileType)
	}

}
