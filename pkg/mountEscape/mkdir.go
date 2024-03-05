package mountEscape

import (
	"fmt"
	"os"
)

func Mkdir(target string) error {
	// 创建挂载点目录
	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}
	return nil
}

func RemoveDir(target string) error {
	f, err := os.Open(target)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // 尝试读取一个文件或目录
	if err == nil {
		// 目录为空，可以删除
		err := os.Remove(target)
		if err != nil {
			return err
		}
		fmt.Println("Directory removed successfully.")
		return nil
	}

	// 目录不为空，返回错误
	return fmt.Errorf("directory not empty")
}
