<<<<<<< HEAD
// +build !cgo

package term

import (
	"syscall"
	"unsafe"
)

const (
	getTermios = syscall.TCGETS
	setTermios = syscall.TCSETS
)

// Termios is the Unix API for terminal I/O.
type Termios struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Cc     [20]byte
	Ispeed uint32
	Ospeed uint32
}
=======
package term

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	getTermios = unix.TCGETS
	setTermios = unix.TCSETS
)

// Termios is the Unix API for terminal I/O.
type Termios unix.Termios
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2

// MakeRaw put the terminal connected to the given file descriptor into raw
// mode and returns the previous state of the terminal so that it can be
// restored.
func MakeRaw(fd uintptr) (*State, error) {
	var oldState State
<<<<<<< HEAD
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, getTermios, uintptr(unsafe.Pointer(&oldState.termios))); err != 0 {
=======
	if _, _, err := unix.Syscall(unix.SYS_IOCTL, fd, getTermios, uintptr(unsafe.Pointer(&oldState.termios))); err != 0 {
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
		return nil, err
	}

	newState := oldState.termios

<<<<<<< HEAD
	newState.Iflag &^= (syscall.IGNBRK | syscall.BRKINT | syscall.PARMRK | syscall.ISTRIP | syscall.INLCR | syscall.IGNCR | syscall.ICRNL | syscall.IXON)
	newState.Oflag &^= syscall.OPOST
	newState.Lflag &^= (syscall.ECHO | syscall.ECHONL | syscall.ICANON | syscall.ISIG | syscall.IEXTEN)
	newState.Cflag &^= (syscall.CSIZE | syscall.PARENB)
	newState.Cflag |= syscall.CS8

	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, setTermios, uintptr(unsafe.Pointer(&newState))); err != 0 {
=======
	newState.Iflag &^= (unix.IGNBRK | unix.BRKINT | unix.PARMRK | unix.ISTRIP | unix.INLCR | unix.IGNCR | unix.ICRNL | unix.IXON)
	newState.Oflag &^= unix.OPOST
	newState.Lflag &^= (unix.ECHO | unix.ECHONL | unix.ICANON | unix.ISIG | unix.IEXTEN)
	newState.Cflag &^= (unix.CSIZE | unix.PARENB)
	newState.Cflag |= unix.CS8

	if _, _, err := unix.Syscall(unix.SYS_IOCTL, fd, setTermios, uintptr(unsafe.Pointer(&newState))); err != 0 {
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
		return nil, err
	}
	return &oldState, nil
}
