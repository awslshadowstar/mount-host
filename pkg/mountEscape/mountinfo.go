package mountEscape

import (
	"fmt"
	"os"
	"strings"
)

func GetMountInfo() (MajorAndMinor, fstype string, err error) {
	data, err := os.ReadFile("/proc/self/mountinfo")
	if err != nil {
		return "", "", err
	}

	lines := strings.Split(string(data), "\n")

	// 遍历每一行，找到第五列为 /etc/hosts 的行并输出第三列
	for _, line := range lines {
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 8 {
			continue
		}
		if fields[4] == "/etc/hosts" {
			return fields[2], fields[7], nil
		}
	}
	return "", "", fmt.Errorf("not found device mount to /etc/hosts")
}
