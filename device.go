package libevdev

import (
	"encoding/binary"
	"os"
	"syscall"
)

type Device struct {
	fn string // path to input device devnode
	fd *os.File
}

// Open
func Open(devnode string) (*Device, error) {
	fd, err := os.OpenFile(devnode, os.O_RDONLY|syscall.O_NONBLOCK, 0644)
	if err != nil {
		return nil, err
	}
	dev := Device{
		fn: devnode,
		fd: fd,
	}

	return &dev, nil
}

// Close closes the underlying device node
func (d *Device) Close() error {
	d.Release()
	if d.fd != nil {
		return d.fd.Close()
	}
	return nil
}

// Grab grabs the device. This prevents other clients from
// receiving events from this device.
// This is generally a bad idea. Don't do this
func (d *Device) Grab() bool {
	return ioctl(d.fd.Fd(), EVIOCGRAB, 1) == nil
}

// Release releases a lock, previously obtained through `Device.Grab`.
func (d *Device) Release() bool {
	return ioctl(d.fd.Fd(), EVIOCGRAB, 0) == nil
}

// ReadOne read and  return a single input event
func (self *Device) ReadOne() (*InputEvent, error) {
	event := InputEvent{}

	err := binary.Read(self.fd, binary.LittleEndian, &event)
	if err != nil {
		return &event, err
	}

	return &event, err
}
