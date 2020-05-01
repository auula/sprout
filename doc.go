/*
LogKer  for Golang simple  logging library.

The simplest way to use LogKer is simply the package-level exported logger:

Console Logging :

  package main

  import (
    klog "github.com/higker/logker"
  )

  func main() {
		// Custom logging message template
		format := "{level} - 时间 {time}  - 位置 {position} - 消息 {message}"
    	//Create Log Buffer
    	task := InitAsync(Qs1w)
		log,e := klog.NewClog(klog.DEBUG, klog.Shanghai,format,task)
 		if e != nil{
        	..... //Custom Operation
    	}
		log.Debug("DEBUG %s","自定义日志消息匹配符测试")
		log.Info("%v", log)
		log.Warning("%v", logker.Shanghai)
		log.Error("ERROR")
  }

Output:
  [ DEBUG ] - 时间 2020-04-20 11:57:23.8927  - 位置 main.go|main.main:23 - 消息 DEBUG 自定义日志消息匹配符测试
  [  INFO ] - 时间 2020-04-20 11:57:23.8928  - 位置 main.go|main.main:24 - 消息 &{0 Asia/Shanghai 0xc00008e220 {level} - 时间 {t位置 {position} - 消息 {message}}
  [WARNING] - 时间 2020-04-20 11:57:23.8928  - 位置 main.go|main.main:25 - 消息 Asia/Shanghai
  [ ERROR ] - 时间 2020-04-20 11:57:23.8929  - 位置 main.go|main.main:26 - 消息 ERROR




File Logging :

package main

import (
	klog "github.com/higker/logker"
	"time"
)

func main() {
	// Specify file location! Create folder in advance!!
	dir := "/Users/ding/Documents/test_log"
	// New file logger
	// File Max size : You can also use built-in constants
	// klog.GB1  	= 1GB
	// klog.MB10  	= 10MB
 	// klog.MB100	= 100MB
	format := "{level} - DATE {time}  - POS {position} - MSG {message}"
	//Create Log Buffer
    	task := InitAsync(Qs1w)
   	flog,e := klog.NewFlog(klog.DEBUG, true, klog.Shanghai, dir, "log", 10*1024, 0777,format,task)
    	if e != nil{
       	    .... //Custom Operation
   	}
	// Analog output log
	for {
		flog.Debug("DEBUG : %d + %d = %d",1,2,1+2)
		flog.Error("ERROR")
		flog.Warning("WARNING %p",flog)
		flog.Info("INFO %s","Hello LogKer.")
		time.Sleep(2 * time.Second)
	}
}

For a full guide visit https://github.com/higker/logker
*/
package logker
