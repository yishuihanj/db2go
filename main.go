package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/yishuihanj/db2go/builder"
)

//go:generate go build
func main() {
	// 加载命令行参数
	driver, cfg := builder.Load()

	// 执行生成命令
	if err := builder.Generate(driver, cfg); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// 格式化
	cmd := exec.Command("go", "fmt", cfg.Out)
	if err := cmd.Start(); err != nil {
		fmt.Printf("format go files failed,%v", err)
		os.Exit(1)
	}
	fmt.Printf(" ✅  完成任务\n\n")
}
