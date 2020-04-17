// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
package logker

import (
	"github.com/fatih/color"
	"testing"
)

// thank you very much https://github.com/fatih/color ^_^
func TestColor(t *testing.T) {
	// Print with default helper functions
	color.Cyan("Prints text in cyan.")
	// A newline will be appended automatically
	color.Blue("Prints %s in blue.", "text")
	// These are using the default foreground colors
	color.Red("We have red")
	color.Magenta("And many others ..")
}
