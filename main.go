package main

import (
	"fmt"

	"github.com/awslshadowstar/mount-host/pkg/mountEscape"
)

func main() {
	if err := mountEscape.MountHost(); err != nil {
		fmt.Println(err)
	}
}
