<p align="center">
    <a href="https://github.com/Higker/logker/"><img src="https://i.loli.net/2020/04/18/9JnSbmkist8dUaC.png" width="128"/></a>
    <h1 align="center">LogKer</h1>
</p>
<p align="center"><code>
logker</code> 是一个Golang小型简单易用的日志库.👨‍💻‍</p>

<p align="center">
    📚<a href="https://github.com/Higker/logker/blob/master/README.md" target="_blank">English Docs</a> | 
    🤩<a href="https://pkg.go.dev/github.com/higker/logker?tab=doc" target="_blank">pkg.go.dev</a> | 
    👨‍💻‍ <a href="https://github.com/Higker/logker/releases" target="_blank">Release版本</a> 
</p>

---

## 特性&简介🌲

- 支持文件自动备份。
- 保存到到指定位置。
- 支持彩色日志信息打印。
- 自定义日志文件存储大小。
- 支持设置日志文件的时区。
- `Error`级别的日志单独输出到指定文件。
- 支持自动根据设置的文件大小切割日志文件。
- 日志级别有4种 `debug` `info` `error` `warning` 。
- `未来会支持: 网络传输存储 websoket🙏.`
- `本项目会不断更新着. 如果你喜欢😍请把你的star给我吧 Thanks♪(･ω･)ﾉ!`
- 这是你吗?😜

![](https://tva1.sinaimg.cn/large/007S8ZIlgy1ge3na9rkzwg308c04oe5c.gif)

## 升级目录
- [v1.1.5](https://github.com/Higker/logker/releases/tag/v1.1.5) 支持自定义消息输出格式  [issues1](https://github.com/Higker/logker/issues/1)
- [v1.1.6](https://github.com/Higker/logker/releases/tag/v1.1.6) 支持异步写日志了，代码库更小了. 


## 历史文档
> 当前文档版本 v1.1.6
- [v1.1.5 Doc](https://github.com/Higker/logker/blob/master/readme_v1.1.5.md)

## 开始安装

🔝 最底版本要求 Go version  **1.11**.
🔝 你的包管理必须使用的是go module!!!

```shell script
go get github.com/higker/logker
```
或者
```shell script
go get -u github.com/higker/logker
```

运行时也可以添加使用参数 `-u` 来获取更新.


## 使用演示
#### 1. File Logger
```go
package main

import (
	klog "github.com/higker/logker"
	"time"
)

func main() {
	// 在初始化的时候必须自己先创建好存储日志的文件夹!!
	dir := "/Users/ding/Documents/test_log"
	// New file logger
	// File Max size : 这里是单个的日志文件大小 你可以自定义也可以使用内置的常量
	// klog.GB1  	= 1GB
	// klog.MB10  	= 10MB
 	// klog.MB100	= 100MB
	format := "{level} - 时间 {time}  - 位置 {position} - 消息 {message}" 
  	//创建日志缓冲区
  	task := InitAsync(Qs1w) //This version was modified from v 1.1.6
  	flog,e := klog.NewFlog(klog.DEBUG, true, klog.Shanghai, dir, "log", 10*1024, 0777,format,task)
  	if e != nil{
       		.... //根据自己情况自定义操作
  	}
	// 模拟日志输出
	for {
		flog.Debug("DEBUG : %d + %d = %d",1,2,1+2)
		flog.Error("ERROR")
		flog.Warning("WARNING %p",flog)
		flog.Info("INFO %s","Hello LogKer.")
		time.Sleep(2 * time.Second)
	}
}
```
> 👆参数列表:
```go

// Build File logger
// Args note
// logLevel:    lev,       \\ 日志等级
// wheError:    wheErr,    \\ 是否开启error级别的日志信息单独输出到一个error_开头文件
// directory:   dir,	     \\ 日志存储的路径&文件夹
// fileName:    fileName,  \\ 日志的名字 不需要写后缀名
// timeZone:    zone,	     \\ 你需要设置的时区 可以使用内置常量  你可以在下面查看文档链接
// power:       power,     \\ 你的文件系统权限
// fileMaxSize: size,      \\ 单个日志文件大小
// task := InitAsync(Qs1w) \\协程日志缓冲区
// format:
	//现在upgrade到v1.1.5版本即可自定义消息输出格式~
	//自定义标签名字必须是{level} {time} {position} {message}
	//自定义标签的位置就是程序运行时输出对应的日志消息的位置！！！！
	//例如下面我自定义的
	// 1. //format := "{level} - {time} - {position} - {message}"
	format := "{level} - 时间 {time}  - 位置 {position} - 消息 {message}" //This version was modified from v 1.1.5
```

#### 2. Console Logger

```go

package main

import (
	"github.com/higker/logker"
	"strings"
)

//type Formatting string

// logKer library Test
func main() {
	// Now upgrade to version v 1.1.5 to customize the message output format ~
	// TheCustomTagNameMustBe {level} {time} {position} {message}
	// The location of the custom label is the location
	// where the program outputs the corresponding log message at run time！！！！
	// forExampleHereIsMyCustom
	// 1. //format := "{level} - {time} - {position} - {message}"
	format := "{level} - 时间 {time}  - 位置 {position} - 消息 {message}" //This version was modified from v 1.1.5
	//日志缓冲区
  	task := InitAsync(Qs1w) //This version was modified from v 1.1.6
	log,e := logker.NewClog(logker.DEBUG, logker.Shanghai, format,task)
  	if e != nil{
       		..... //自定义操作
  	}
	log.Debug("DEBUG %s","自定义日志消息匹配符测试")
	log.Info("%v", log)
	log.Warning("%v", logker.Shanghai)
	log.Error("ERROR")
}
```
## OutPut
```shell
[ DEBUG ] - 时间 2020-04-20 11:57:23.8927  - 位置 main.go|main.main:23 - 消息 DEBUG 自定义日志消息匹配符测试
[  INFO ] - 时间 2020-04-20 11:57:23.8928  - 位置 main.go|main.main:24 - 消息 &{0 Asia/Shanghai 0xc00008e220 {level} - 时间 {t位置 {position} - 消息 {message}}
[WARNING] - 时间 2020-04-20 11:57:23.8928  - 位置 main.go|main.main:25 - 消息 Asia/Shanghai
[ ERROR ] - 时间 2020-04-20 11:57:23.8929  - 位置 main.go|main.main:26 - 消息 ERROR

```
#### 3. 演示截图:
> 截图没有更新~ 当前截图为 v1.0.9 version,你可以自己安装库使用查看效果,当然可能是新版本的牛逼一点2333333
![LogKerGolang](https://i.loli.net/2020/04/18/Jjv82WDsyGtCaEH.png)

![log,golang,logKer](https://i.loli.net/2020/04/18/mJnvBp7oXwd8KSU.png)

## 鸣谢名单🤝
- fatih (https://github.com/fatih/color)
- Icon Mafia (Logo & Banner 设计)
- When are you? 😜

## 其他帮助 
- `欢迎提交issues👏 不一定能及时处理因为开源项目,我还有其他工作需要做。`
- `欢迎🇨🇳中国区小伙伴加入代码贡献.不对应该会中文交流就可以23333💐`
- `Pkg Docs: `[click to pkg.go.dev](https://pkg.go.dev/github.com/higker/logker?tab=doc)

## 开源协议
This project open source is MIT License
. See the [LICENSE](LICENSE) file content.
