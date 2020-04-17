// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/17 - 8:21 下午

package logker

import (
	"os"
	"testing"
)

func Test_console_OutPutMessage(t *testing.T) {
	type fields struct {
		logLevel level
		timeZone logTimeZone
		tz       *timeZone
	}
	type args struct {
		model level
		v     string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &console{
				logLevel: tt.fields.logLevel,
				timeZone: tt.fields.timeZone,
				tz:       tt.fields.tz,
			}
		})
	}
}

func Test_fileLog_OutPutMessage(t *testing.T) {
	type fields struct {
		logLevel  level
		wheError  bool
		directory string
		fileName  string
		file      *os.File
		errFile   *os.File
		tz        *timeZone
		timeZone  logTimeZone
		power     os.FileMode
	}
	type args struct {
		model level
		v     string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fileLog{
				logLevel:  tt.fields.logLevel,
				wheError:  tt.fields.wheError,
				directory: tt.fields.directory,
				fileName:  tt.fields.fileName,
				file:      tt.fields.file,
				errFile:   tt.fields.errFile,
				tz:        tt.fields.tz,
				timeZone:  tt.fields.timeZone,
				power:     tt.fields.power,
			}
		})
	}
}

func Test_fileLog_outPut(t *testing.T) {
	type fields struct {
		logLevel  level
		wheError  bool
		directory string
		fileName  string
		file      *os.File
		errFile   *os.File
		tz        *timeZone
		timeZone  logTimeZone
		power     os.FileMode
	}
	type args struct {
		lev level
		v   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fileLog{
				logLevel:  tt.fields.logLevel,
				wheError:  tt.fields.wheError,
				directory: tt.fields.directory,
				fileName:  tt.fields.fileName,
				file:      tt.fields.file,
				errFile:   tt.fields.errFile,
				tz:        tt.fields.tz,
				timeZone:  tt.fields.timeZone,
				power:     tt.fields.power,
			}
			if err := f.outPut(tt.args.lev, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("outPut() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
