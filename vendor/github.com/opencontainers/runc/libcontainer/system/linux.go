// +build linux

package system

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
<<<<<<< HEAD
	"syscall"
	"unsafe"
=======
	"syscall" // only for exec
	"unsafe"

	"golang.org/x/sys/unix"
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
)

// If arg2 is nonzero, set the "child subreaper" attribute of the
// calling process; if arg2 is zero, unset the attribute.  When a
// process is marked as a child subreaper, all of the children
// that it creates, and their descendants, will be marked as
// having a subreaper.  In effect, a subreaper fulfills the role
// of init(1) for its descendant processes.  Upon termination of
// a process that is orphaned (i.e., its immediate parent has
// already terminated) and marked as having a subreaper, the
// nearest still living ancestor subreaper will receive a SIGCHLD
// signal and be able to wait(2) on the process to discover its
// termination status.
const PR_SET_CHILD_SUBREAPER = 36

type ParentDeathSignal int

func (p ParentDeathSignal) Restore() error {
	if p == 0 {
		return nil
	}
	current, err := GetParentDeathSignal()
	if err != nil {
		return err
	}
	if p == current {
		return nil
	}
	return p.Set()
}

func (p ParentDeathSignal) Set() error {
	return SetParentDeathSignal(uintptr(p))
}

func Execv(cmd string, args []string, env []string) error {
	name, err := exec.LookPath(cmd)
	if err != nil {
		return err
	}

	return syscall.Exec(name, args, env)
}

<<<<<<< HEAD
func Prlimit(pid, resource int, limit syscall.Rlimit) error {
	_, _, err := syscall.RawSyscall6(syscall.SYS_PRLIMIT64, uintptr(pid), uintptr(resource), uintptr(unsafe.Pointer(&limit)), uintptr(unsafe.Pointer(&limit)), 0, 0)
=======
func Prlimit(pid, resource int, limit unix.Rlimit) error {
	_, _, err := unix.RawSyscall6(unix.SYS_PRLIMIT64, uintptr(pid), uintptr(resource), uintptr(unsafe.Pointer(&limit)), uintptr(unsafe.Pointer(&limit)), 0, 0)
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	if err != 0 {
		return err
	}
	return nil
}

func SetParentDeathSignal(sig uintptr) error {
<<<<<<< HEAD
	if _, _, err := syscall.RawSyscall(syscall.SYS_PRCTL, syscall.PR_SET_PDEATHSIG, sig, 0); err != 0 {
=======
	if _, _, err := unix.RawSyscall(unix.SYS_PRCTL, unix.PR_SET_PDEATHSIG, sig, 0); err != 0 {
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
		return err
	}
	return nil
}

func GetParentDeathSignal() (ParentDeathSignal, error) {
	var sig int
<<<<<<< HEAD
	_, _, err := syscall.RawSyscall(syscall.SYS_PRCTL, syscall.PR_GET_PDEATHSIG, uintptr(unsafe.Pointer(&sig)), 0)
=======
	_, _, err := unix.RawSyscall(unix.SYS_PRCTL, unix.PR_GET_PDEATHSIG, uintptr(unsafe.Pointer(&sig)), 0)
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	if err != 0 {
		return -1, err
	}
	return ParentDeathSignal(sig), nil
}

func SetKeepCaps() error {
<<<<<<< HEAD
	if _, _, err := syscall.RawSyscall(syscall.SYS_PRCTL, syscall.PR_SET_KEEPCAPS, 1, 0); err != 0 {
=======
	if _, _, err := unix.RawSyscall(unix.SYS_PRCTL, unix.PR_SET_KEEPCAPS, 1, 0); err != 0 {
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
		return err
	}

	return nil
}

func ClearKeepCaps() error {
<<<<<<< HEAD
	if _, _, err := syscall.RawSyscall(syscall.SYS_PRCTL, syscall.PR_SET_KEEPCAPS, 0, 0); err != 0 {
=======
	if _, _, err := unix.RawSyscall(unix.SYS_PRCTL, unix.PR_SET_KEEPCAPS, 0, 0); err != 0 {
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
		return err
	}

	return nil
}

func Setctty() error {
<<<<<<< HEAD
	if _, _, err := syscall.RawSyscall(syscall.SYS_IOCTL, 0, uintptr(syscall.TIOCSCTTY), 0); err != 0 {
=======
	if _, _, err := unix.RawSyscall(unix.SYS_IOCTL, 0, uintptr(unix.TIOCSCTTY), 0); err != 0 {
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
		return err
	}
	return nil
}

// RunningInUserNS detects whether we are currently running in a user namespace.
// Copied from github.com/lxc/lxd/shared/util.go
func RunningInUserNS() bool {
	file, err := os.Open("/proc/self/uid_map")
	if err != nil {
		// This kernel-provided file only exists if user namespaces are supported
		return false
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	l, _, err := buf.ReadLine()
	if err != nil {
		return false
	}

	line := string(l)
	var a, b, c int64
	fmt.Sscanf(line, "%d %d %d", &a, &b, &c)
	/*
	 * We assume we are in the initial user namespace if we have a full
	 * range - 4294967295 uids starting at uid 0.
	 */
	if a == 0 && b == 0 && c == 4294967295 {
		return false
	}
	return true
}

// SetSubreaper sets the value i as the subreaper setting for the calling process
func SetSubreaper(i int) error {
<<<<<<< HEAD
	return Prctl(PR_SET_CHILD_SUBREAPER, uintptr(i), 0, 0, 0)
}

func Prctl(option int, arg2, arg3, arg4, arg5 uintptr) (err error) {
	_, _, e1 := syscall.Syscall6(syscall.SYS_PRCTL, uintptr(option), arg2, arg3, arg4, arg5, 0)
	if e1 != 0 {
		err = e1
	}
	return
=======
	return unix.Prctl(PR_SET_CHILD_SUBREAPER, uintptr(i), 0, 0, 0)
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
}
