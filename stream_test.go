// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/17 - 8:21 下午

package logker

import (
	"fmt"
	"testing"
)

func Test_fileLog_outPut(t *testing.T) {
	dir := "/Users/ding/Documents/test_log/"
	flog := new(fileLog)
	flog.logLevel = DEBUG
	flog.directory = dir
	flog.fileName = "logging"
	file := flog.initFilePtr()
	fmt.Printf("1 %T %p \n", file, file)

	flog.file = file
	fmt.Printf("2 %T %p \n", file, file)
	flog.power = 0666
	flog.timeZone = Shanghai
	flog.tz = &timeZone{TimeZoneStr: flog.timeZone}
	fmt.Println(flog.tz)
}
