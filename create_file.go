package main

import (
	"fmt"
	"os"
)

//创建文件
func CreateFile(_tableName, s string, outDir string) error {
	exist, err := pathExists(outDir)
	if err != nil {
		return err
	}
	//文件夹不存在 创建
	if !exist {
		fmt.Printf("\"%s\" 输出目录不存在，创建目录...\n", outDir)
		os.Mkdir(outDir, os.ModePerm)
	}

	f, err := os.Create(fmt.Sprintf("%s/%s.go", outDir, _tableName))
	defer f.Close()
	if err != nil {
		fmt.Printf("错误! 创建 %s.go 文件失败，err:%v", _tableName, err.Error())
		return err
	} else {
		_, err = f.Write([]byte(s))
		if err != nil {
			fmt.Printf("错误! 创建 %s.go 文件失败，err:%v", _tableName, err.Error())
			return err
		}
	}
	fmt.Printf("创建 %s.go 文件成功，路径为：%s\n", _tableName, f.Name())
	return nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
