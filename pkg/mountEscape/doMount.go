package mountEscape

import (
	"fmt"
	"syscall"
)

func DoMount(device, target, fsType string) error {

	flags := syscall.MS_NOATIME // 可以根据需要添加更多的挂载选项

	// 进行挂载
	if err := syscall.Mount(device, target, fsType, uintptr(flags), ""); err != nil {
		return err
	}

	return nil
}

func Umount(target string) error {
	err := syscall.Unmount(target, 0)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
