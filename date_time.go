// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/17 - 12:26 下午

package logker

import (
	"strconv"
	"time"
)

type logTimeZone string

const (
	//https://blog.imdst.com/guo-bf/
	//美东     EST         /usr/share/zoneinfo/EST5EDT
	//美西     PST         /usr/share/zoneinfo/PST8PDT
	//澳洲     CST         /usr/share/zoneinfo/CST6CDT
	//欧洲     GMT         /usr/share/zoneinfo/Europe/London
	//泰国     Thailand    /usr/share/zoneinfo/Asia/Bangkok
	//越南     Vietnam     /usr/share/zoneinfo/Asia/Ho_Chi_Minh
	//新加坡   Singapore   /usr/share/zoneinfo/Asia/Singapore
	//台湾     Taiwan      /usr/share/zoneinfo/Asia/Taipei
	//香港     HongKong    /usr/share/zoneinfo/Asia/Hong_Kong
	//日本     Japan       /usr/share/zoneinfo/Asia/Tokyo
	//韩国     Korea       /usr/share/zoneinfo/Asia/Pyongyang
	EST       logTimeZone = "EST5EDT"
	PST       logTimeZone = "PST8PDT"
	CST       logTimeZone = "CST6CDT"
	GMT       logTimeZone = "Europe/London"
	Thailand  logTimeZone = "Asia/Bangkok"
	Vietnam   logTimeZone = "Asia/Ho_Chi_Minh"
	Singapore logTimeZone = "Asia/Singapore"
	Taiwan    logTimeZone = "Asia/Taipei"
	HongKong  logTimeZone = "Asia/Hong_Kong"
	Japan     logTimeZone = "Asia/Tokyo"
	Korea     logTimeZone = "Asia/Pyongyang"
	Shanghai  logTimeZone = "Asia/Shanghai" // Shanghai China

	//timeFormat    = "2006-01-02 15:04:05.0000 PM"
	timeFormat    = "2006-01-02 15:04:05.0000"
	logTimeFormat = "2006_01_02_15_04_05.0000"
)

// 自定义时间类型 ：customize time zone type struct
type timeZone struct {
	TimeZoneStr logTimeZone // your custom time zone,recommend use this.
}

// get time str fmt: 2006-01-02 15:04:05.0000 PM
func (tz *timeZone) NowTimeStr() (ts string) {
	return tz.localTime().Format(timeFormat)
}

// Log file name format : log_2006_01_02_15_04_05_0000.log
func (tz *timeZone) NowTimeStrLogName() (ts string) {
	return tz.localTime().Format(logTimeFormat)
}

// ret: nanosecond string
func (tz *timeZone) NanoSecondStr() string {
	return strconv.FormatInt(tz.localTime().UnixNano(), 10)
}

//set customize time zone
func (tz *timeZone) localTime() time.Time {
	//  set time zone
	ltz, err := time.LoadLocation(string(tz.TimeZoneStr))
	if err != nil {
		panic("core init set time zone fail.")
	}
	return time.Now().In(ltz)
}
