// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
	"testing"
)

// 格式: "\033[风格;前景色;背景色m内容\033[0m"
//go:generate echo -e "\033[4;31;42m你好\033[0m"
func TestColor(t *testing.T) {
	fmt.Printf("\033[%dm%s\033[0m\n", FgMagenta, "带前景色的字体")
	fmt.Printf("\033[%d;%dm%s\033[0m\n", Bold, FgBlue, "带前景色和样式的字体")
	fmt.Printf("\033[%d;%dm%s\033[0m\n", FgBlue, BgGreen, "带前景色和背景色的字体")
	fmt.Printf("\033[%d;%d;%dm%s\033[0m\n", Underline, FgWhite, BgMagenta, "带前景色、背景色和样式的字体")
}

func TestPascal(t *testing.T) {
	fmt.Println(Pascal("_user_name"))
	fmt.Println(Pascal("_username"))
	fmt.Println(Pascal("_user_last_name"))
	fmt.Println(Pascal("user_name"))
	fmt.Println(Pascal("User_Name"))
	fmt.Println(Pascal("USERNAME"))
	fmt.Println(Pascal("username"))
	fmt.Println(Pascal("user__name"))
	fmt.Println(Pascal("_user_name_"))
	fmt.Println(Pascal("_user_name_99"))
}
