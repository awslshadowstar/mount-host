package mountEscape

import (
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
	return os.Remove(target)
}
