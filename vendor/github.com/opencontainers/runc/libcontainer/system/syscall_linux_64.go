// +build linux,arm64 linux,amd64 linux,ppc linux,ppc64 linux,ppc64le linux,s390x

package system

import (
<<<<<<< HEAD
	"syscall"
=======
	"golang.org/x/sys/unix"
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
)

// Setuid sets the uid of the calling thread to the specified uid.
func Setuid(uid int) (err error) {
<<<<<<< HEAD
	_, _, e1 := syscall.RawSyscall(syscall.SYS_SETUID, uintptr(uid), 0, 0)
=======
	_, _, e1 := unix.RawSyscall(unix.SYS_SETUID, uintptr(uid), 0, 0)
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	if e1 != 0 {
		err = e1
	}
	return
}

// Setgid sets the gid of the calling thread to the specified gid.
func Setgid(gid int) (err error) {
<<<<<<< HEAD
	_, _, e1 := syscall.RawSyscall(syscall.SYS_SETGID, uintptr(gid), 0, 0)
=======
	_, _, e1 := unix.RawSyscall(unix.SYS_SETGID, uintptr(gid), 0, 0)
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	if e1 != 0 {
		err = e1
	}
	return
}
