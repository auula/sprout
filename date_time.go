// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/17 - 12:26 下午

package logker

import (
	"strconv"
	"time"
)

type LogTimeZone string

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
	EST       LogTimeZone = "EST5EDT"
	PST       LogTimeZone = "PST8PDT"
	CST       LogTimeZone = "CST6CDT"
	GMT       LogTimeZone = "Europe/London"
	Thailand  LogTimeZone = "Asia/Bangkok"
	Vietnam   LogTimeZone = "Asia/Ho_Chi_Minh"
	Singapore LogTimeZone = "Asia/Singapore"
	Taiwan    LogTimeZone = "Asia/Taipei"
	HongKong  LogTimeZone = "Asia/Hong_Kong"
	Japan     LogTimeZone = "Asia/Tokyo"
	Korea     LogTimeZone = "Asia/Pyongyang"
	Shanghai  LogTimeZone = "Asia/Shanghai" // Shanghai China

	timeFormat = "2006-01-02 15:04:05.0000 PM"
)

// 自定义时间类型 ：customize time zone type struct
type timeZone struct {
	TimeZoneStr LogTimeZone // your custom time zone,recommend use this.
}

// get time str fmt: 2006-01-02 15:04:05.0000 PM
func (tz *TimeZone) NowTimeStr() (ts string) {
	return tz.localTime().Format(timeFormat)
}

// ret: nanosecond string
func (tz *TimeZone) NanoSecondStr() string {
	return strconv.FormatInt(tz.localTime().UnixNano(), 10)
}

//set customize time zone
func (tz *TimeZone) localTime() time.Time {
	//  set time zone
	ltz, err := time.LoadLocation(string(tz.TimeZoneStr))
	if err != nil {
		panic("core init set time zone fail.")
	}
	return time.Now().In(ltz)
}
