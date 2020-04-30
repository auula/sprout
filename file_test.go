// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/1 - 2:03 上午

package logker

import "testing"

func TestOutPutFileLog(t *testing.T) {
	// Specify file location! Create folder in advance!!
	dir := "/Users/ding/Documents/test_log"
	defaults := "{level} - Date: {time}  {position} - Message: {message}"
	task := InitAsync(Qs1w)
	logger, _ := NewFlog(DEBUG, true, Shanghai, dir, "log", 10*1024, 0777, defaults, task)
	for  {
		logger.Info("INFO")
		logger.Debug("DEBUG")
		logger.Warning("WARNING")
		logger.Error("ERROR")
	}
}