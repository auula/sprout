// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/30 - 9:16 下午

package logker

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

// Channel Buffer Size
const (
	_    = iota
	Qs1w = 10000 * iota
	Qs2w
	Qs3w
	Qs4w
	Qs5w
	Qs6w
)

const (
	// Log message format wildcards
	// {level} == logging level
	// {time} == time
	// {position} == runtime caller info
	// {message} == logging message
	// OutPut: [INFO] 2006-01-02 13:05.0006 MP - Position: test.go|main.test:21 - Message: news
	DefaultFormat = "{level} - Date: {time}  {position} - Message: {message}" //This was modified for v 1.1.5
)

type AsyncTask struct {
	queueSize int64
	logQueue  chan *message
}

// Initializes The Producer Channel Buffer
func InitAsyncTask(queueSize int64) *AsyncTask {
	return &AsyncTask{queueSize: queueSize, logQueue: make(chan *message, queueSize)}
}

// send pack msg to log queue
func (a *AsyncTask) sendMsg(msg *message) {
	a.logQueue <- msg
}

func (a *AsyncTask) begin(t *testing.T) {
	go a.asyncOutPutTask(t)
	t.Log("atlog 启动了")
}

func (a *AsyncTask) asyncOutPutTask(t *testing.T) {
	file, _ := os.OpenFile("/Users/ding/Documents/GO_CODE_DEV/src/logker/text.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	writer := bufio.NewWriter(file)
	for {
		msg := <-a.logQueue
		_, _ = writer.WriteString(msg.msg + msg.caller + msg.logTime + msg.level.toStr() + "\n")
	}
}

// ReplaceOurCustomMessageFormatIdentifier
// This function was added at 23:50:28 on April 19, 2020 in v1.1.5
func (a *AsyncTask) parsePacket(msg message) string {
	str := msg.format
	str = strings.Replace(str, "{level}", msg.level.toStr(), -1)
	str = strings.Replace(str, "{time}", msg.logTime, -1)
	str = strings.Replace(str, "{position}", msg.caller, -1)
	str = strings.Replace(str, "{message}", msg.msg, -1)
	return str
}
