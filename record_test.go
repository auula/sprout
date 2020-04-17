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
	dir := "/Users/ding/Documents/GO_CODE_DEV/src/logker"
	flog := new(fileLog)
	flog.logLevel = DEBUG
	flog.directory = dir
	flog.fileName = "logging"
	file, e := flog.initFilePtr()
	fmt.Printf("1 %T %p \n", file, file)
	if e != nil {
		fmt.Println(e)
	}
	flog.file = file
	fmt.Printf("2 %T %p \n", file, file)
	flog.power = 0666
	flog.timeZone = Shanghai
	flog.tz = &timeZone{TimeZoneStr: flog.timeZone}
	flog.outPut(DEBUG, "testetstststtststs")

}
