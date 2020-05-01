// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/5/1 - 12:53 上午

package logker

const (
	// Log message format wildcards
	// {level} == logging level
	// {time} == time
	// {position} == runtime caller info
	// {message} == logging message
	// This was modified for v 1.1.5
	// OutPut: [INFO] 2006-01-02 13:05.0006 MP - Position: test.go|main.test:21 - Message: news
	DefaultFormat = "{level} - Date: {time}  {position} - Message: {message}"
)

// Asynchronous Task Interface
// added after v1.1.6
type Async interface {
	sendMsg(msg *message)
	begin()
}

// Asynchronous Task Processing
type AsyncTask struct {
	queueSize int64
	logQueue  chan *message
}

// Initializes The Logging Producer Channel Buffer
// Queue Size Data Type int64
func InitAsync(queueSize int64) *AsyncTask {
	return &AsyncTask{queueSize: queueSize, logQueue: make(chan *message, queueSize)}
}
