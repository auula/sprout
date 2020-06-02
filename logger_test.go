// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/30 - 8:52 下午

package logker

import "testing"

func TestSize(t *testing.T) {
	// TestByteSize
	t.Log(KB)
	t.Log(MB)
	t.Log(GB)
	t.Log(TB)
}

func TestInfo(t *testing.T) {
	Error("test %d", 1)
}

func Test_level_toStr(t *testing.T) {
	tests := []struct {
		name string
		lev  level
		want string
	}{
		{name: "DEBUG", lev: DEBUG, want: DEBUG.toStr()},
		{name: "ERROR", lev: ERROR, want: DEBUG.toStr()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lev.toStr(); got != tt.want {
				t.Errorf("toStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
