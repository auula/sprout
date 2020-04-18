<p align="center">
    <a href="https://github.com/Higker/logker/"><img src="https://i.loli.net/2020/04/18/9JnSbmkist8dUaC.png" width="128"/></a>
    <h1 align="center">LogKer</h1>
</p>
<p align="center"><code>
logker</code> 是一个Golang小型简单易用的日志库.👨‍💻‍</p>

<p align="center">
    📚<a href="https://github.com/Higker/logker/blob/master/README.md" target="_blank">English Docs</a> | 
    🤩<a href="https://github.com/Higker/logker/" target="_blank">源代码</a> | 
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


## 开始安装

🔝 最底版本要求 Go version  **1.11**.
🔝 你的包管理必须使用的是go module!!!

```bash
go get github.com/higker/logker v1.0.6
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
	flog := klog.NewFlog(klog.DEBUG, true, klog.Shanghai, dir, "log", 10*1024, 0777)
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
// directory:   dir,	   \\ 日志存储的路径&文件夹
// fileName:    fileName,  \\ 日志的名字 不需要写后缀名
// timeZone:    zone,	   \\ 你需要设置的时区 可以使用内置常量  你可以在下面查看文档链接
// power:       power,     \\ 你的文件系统权限
// fileMaxSize: size,      \\ 单个日志文件大小
```

#### 2. Console Logger

```go
package main

import (
	klog "github.com/higker/logker"
	"time"
)

func main() {
	// New Console logger
	clog := klog.NewClog(klog.DEBUG,klog.Shanghai)
	clog.Debug("DEBUG : %d + %d = %d",1,2,1+2)
	clog.Error("ERROR")
	clog.Warning("WARNING %p",clog)
	clog.Info("INFO %s","Hello LogKer.")
}
```
#### 3. 演示截图:
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