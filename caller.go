// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/17 - 1:54 下午

package logker

import (
	"fmt"
	"path"
	"runtime"
)

// Runtime caller skip
const (
	skip = 3
)

// Gets The Program Caller Location
func buildCallerStr(skip int) string {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		panic("error runtime caller.")
	}
	// get function name
	funName := runtime.FuncForPC(pc).Name()
	// file path
	filePath := path.Base(file)
	// build string runtime caller info
	//return fmt.Sprintf("- Pc:%d - Position: %s|%s:%d", pc, filePath, funName, line)
	return fmt.Sprintf("%s|%s:%d", filePath, funName, line)
}
