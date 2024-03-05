package mountEscape

import (
	"bufio"
	"os"
	"strings"
)

func GetDeviceList() (deviceList [][2]string, err error) { // 打开/sys/block目录
	// Open the /proc/partitions file
	file, err := os.Open("/proc/partitions")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Skip the first two lines as they are headers
	for i := 0; i < 2; i++ {
		scanner.Scan()
	}

	// // Print the header
	// fmt.Println("Major\tMinorSize\tDevice\t")

	// Iterate over each line in the file
	for scanner.Scan() {
		// Split the line by whitespace
		fields := strings.Fields(scanner.Text())

		// Check if it's a partition entry
		if len(fields) >= 4 {
			// Get device name and size
			device := fields[3]
			minor := fields[1]
			major := fields[0]
			deviceList = append(deviceList, [2]string{major + ":" + minor, device})
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return deviceList, err
	}
	return deviceList, nil
}
