// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

// 颜色
type Color int

// 前景色
const (
	FgBlack   Color = iota + 30 // 30: 黑色
	FgRed                       // 31: 红色
	FgGreen                     // 32: 绿色
	FgYellow                    // 33: 黄色
	FgBlue                      // 34: 蓝色
	FgMagenta                   // 35: 品红/洋紫
	FgCyan                      // 36: 青色
	FgWhite                     // 37: 白色
)

// 背景色
const (
	BgBlack   Color = iota + 40 // 40: 黑色
	BgRed                       // 41: 红色
	BgGreen                     // 42: 绿色
	BgYellow                    // 43: 黄色
	BgBlue                      // 44: 蓝色
	BgMagenta                   // 45: 品红/洋紫
	BgCyan                      // 46: 青色
	BgWhite                     // 47: 白色
)

// 样式
type Style int

const (
	Reset        Style = iota // 0: 重置
	Bold                      // 1: 加粗
	Faint                     // 2: 模糊
	Italic                    // 3: 斜体
	Underline                 // 4: 下划线
	BlinkSlow                 // 5: 慢速闪烁
	BlinkRapid                // 6: 快速闪烁
	ReverseVideo              // 7: 反白/反向显示
	Concealed                 // 8: 隐藏/暗格
	CrossedOut                // 9: 删除
)

