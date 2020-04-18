<p align="center">
    <a href="https://github.com/Higker/logker/"><img src="https://i.loli.net/2020/04/18/9JnSbmkist8dUaC.png" width="128"/></a>
    <h1 align="center">LogKer</h1>
</p>
<p align="center"><code>
logker</code> It's a log Library of Golang. It's easy to use.👨‍💻‍</p>

<p align="center">
    📚<a href="https://github.com/Higker/logker/readme_zh.md" target="_blank">中文说明</a> | 
    🤩<a href="https://github.com/Higker/logker/" target="_blank">Source Code</a> | 
    👨‍💻‍ <a href="https://github.com/Higker/logker/releases" target="_blank">Release</a> 
</p>

---

## Features

- Support file backup.
- Simple and easy to use.
- Support file storage log.
- Support console color printing.
- Custom log file storage location.
- Support setting time and time zone.
- Support log file size for split storage.
- Error level output to specified file separately.
- Four levels of `debug` `info` `error` `warning` are supported.
- `Future support: The remote computer stores the websoket output🙏.`
- `The project is constantly maintained and updated. I like 😍A kind of Please click star Thanks♪(･ω･)ﾉ!`


## Installation

🔝 The minimum requirement of Go version is **1.11**.
🔝 Your project also uses go module!!!

```bash
go get github.com/higker/logker v1.0.6
```

command add `-u` flag to update in the future.


## Use Example
#### 1. File Logger
```go
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
	flog := klog.NewFlog(klog.DEBUG, true, klog.Shanghai, dir, "log", 10*1024, 0777)
	// Analog output log
	for {
		flog.Debug("DEBUG : %d + %d = %d",1,2,1+2)
		flog.Error("ERROR")
		flog.Warning("WARNING %p",flog)
		flog.Info("INFO %s","Hello LogKer.")
		time.Sleep(2 * time.Second)
	}
}
```
> 👆Parameter note List:
```go

// Build File logger
// Args note
// logLevel:    lev,       \\ logging level
// wheError:    wheErr,    \\ whether enable  error alone file
// directory:   dir,	   \\ logging file save directory
// fileName:    fileName,  \\ logging save file name
// timeZone:    zone,	   \\ load time zone format
// power:       power,     \\ file system power
// fileMaxSize: size,      \\ logging alone file max size
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
#### 3. Effect:
![LogKerGolang](https://i.loli.net/2020/04/18/Jjv82WDsyGtCaEH.png)

![log,golang,logKer](https://i.loli.net/2020/04/18/mJnvBp7oXwd8KSU.png)

## Thank list🤝
- fatih (https://github.com/fatih/color)
- Icon Mafia (Logo & Banner design)
- When are you? 😜

## Other 
- `Welcome to submit work order👏`
- `Pkg Docs: `[click to pkg.go.dev](https://pkg.go.dev/github.com/higker/logker?tab=doc)
## License

This project open source is MIT License
. See the [LICENSE](LICENSE) file content.