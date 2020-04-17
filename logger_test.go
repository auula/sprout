// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/16  9:41 下午

package logker

import (
	"fmt"
	"testing"
)

func TestLevel_ToStr(t *testing.T) {
	fmt.Println(DEBUG.toStr())
	fmt.Println(INFO.toStr())
	fmt.Println(WARNING.toStr())
	fmt.Println(ERROR.toStr())
}
func TestMB(t *testing.T) {
	fmt.Println(MB10)
	fmt.Println(MB100)
	fmt.Println(GB1)
}
