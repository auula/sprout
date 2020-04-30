// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/30 - 11:04 下午

package logker

import (
	"testing"
)

func TestVerifyString(t *testing.T) {
	defaults := "{level} - Date: {time}  {position1} - Message: {message}"
	if err:=verify(defaults);err != nil {
		t.Log(err)
	}
}

func TestConsole_Debug(t *testing.T) {
	defaults := "{level} - Date: {time}  {position} - Message: {message}"
	task := InitAsyncTask(Qs1w)
	logger, e := NewClog(ERROR, Shanghai, defaults, task, t)
	if e != nil {
		t.Log(e)
	}
	for  {
		logger.Error("Debug")
	}
}