package libevdev

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/unix"
)

func ioctl(fd uintptr, name int, data interface{}) error {
	var v uintptr

	switch dd := data.(type) {
	case unsafe.Pointer:
		v = uintptr(dd)

	case int:
		v = uintptr(dd)

	case uintptr:
		v = dd

	default:
		return fmt.Errorf("ioctl: Invalid argument: %T", data)
	}

	_, _, errno := unix.RawSyscall(unix.SYS_IOCTL, fd, uintptr(name), v)
	if errno == 0 {
		return nil
	}

	return errno
}
