package mountEscape

import (
	"fmt"
	"syscall"
)

func DoMount(device, target, fsType string) error {
	fmt.Println(device)
	fmt.Println(target)

	flags := syscall.MS_NOATIME // 可以根据需要添加更多的挂载选项

	// 进行挂载
	if err := syscall.Mount(device, target, fsType, uintptr(flags), ""); err != nil {
		return err
	}

	fmt.Println("设备成功挂载到", target)
	return nil
}

func Umount(target string) error {
	err := syscall.Unmount(target, syscall.MNT_DETACH)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
