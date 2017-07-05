// Based on ssh/terminal:
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux darwin freebsd openbsd netbsd dragonfly
// +build !appengine

package logrus

import (
<<<<<<< HEAD
=======
	"io"
	"os"
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	"syscall"
	"unsafe"
)

// IsTerminal returns true if stderr's file descriptor is a terminal.
<<<<<<< HEAD
func IsTerminal() bool {
	fd := syscall.Stderr
	var termios Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), ioctlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
=======
func IsTerminal(f io.Writer) bool {
	var termios Termios
	switch v := f.(type) {
	case *os.File:
		_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(v.Fd()), ioctlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
		return err == 0
	default:
		return false
	}
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
}
