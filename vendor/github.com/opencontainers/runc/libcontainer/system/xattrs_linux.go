package system

<<<<<<< HEAD
import (
	"syscall"
	"unsafe"
)

var _zero uintptr

// Returns the size of xattrs and nil error
// Requires path, takes allocated []byte or nil as last argument
func Llistxattr(path string, dest []byte) (size int, err error) {
	pathBytes, err := syscall.BytePtrFromString(path)
	if err != nil {
		return -1, err
	}
	var newpathBytes unsafe.Pointer
	if len(dest) > 0 {
		newpathBytes = unsafe.Pointer(&dest[0])
	} else {
		newpathBytes = unsafe.Pointer(&_zero)
	}

	_size, _, errno := syscall.Syscall6(syscall.SYS_LLISTXATTR, uintptr(unsafe.Pointer(pathBytes)), uintptr(newpathBytes), uintptr(len(dest)), 0, 0, 0)
	size = int(_size)
	if errno != 0 {
		return -1, errno
	}

	return size, nil
}
=======
import "golang.org/x/sys/unix"
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2

// Returns a []byte slice if the xattr is set and nil otherwise
// Requires path and its attribute as arguments
func Lgetxattr(path string, attr string) ([]byte, error) {
	var sz int
<<<<<<< HEAD
	pathBytes, err := syscall.BytePtrFromString(path)
	if err != nil {
		return nil, err
	}
	attrBytes, err := syscall.BytePtrFromString(attr)
	if err != nil {
		return nil, err
	}

	// Start with a 128 length byte array
	sz = 128
	dest := make([]byte, sz)
	destBytes := unsafe.Pointer(&dest[0])
	_sz, _, errno := syscall.Syscall6(syscall.SYS_LGETXATTR, uintptr(unsafe.Pointer(pathBytes)), uintptr(unsafe.Pointer(attrBytes)), uintptr(destBytes), uintptr(len(dest)), 0, 0)

	switch {
	case errno == syscall.ENODATA:
		return nil, errno
	case errno == syscall.ENOTSUP:
		return nil, errno
	case errno == syscall.ERANGE:
		// 128 byte array might just not be good enough,
		// A dummy buffer is used ``uintptr(0)`` to get real size
		// of the xattrs on disk
		_sz, _, errno = syscall.Syscall6(syscall.SYS_LGETXATTR, uintptr(unsafe.Pointer(pathBytes)), uintptr(unsafe.Pointer(attrBytes)), uintptr(unsafe.Pointer(nil)), uintptr(0), 0, 0)
		sz = int(_sz)
		if sz < 0 {
			return nil, errno
		}
		dest = make([]byte, sz)
		destBytes := unsafe.Pointer(&dest[0])
		_sz, _, errno = syscall.Syscall6(syscall.SYS_LGETXATTR, uintptr(unsafe.Pointer(pathBytes)), uintptr(unsafe.Pointer(attrBytes)), uintptr(destBytes), uintptr(len(dest)), 0, 0)
		if errno != 0 {
			return nil, errno
		}
	case errno != 0:
		return nil, errno
	}
	sz = int(_sz)
	return dest[:sz], nil
}

func Lsetxattr(path string, attr string, data []byte, flags int) error {
	pathBytes, err := syscall.BytePtrFromString(path)
	if err != nil {
		return err
	}
	attrBytes, err := syscall.BytePtrFromString(attr)
	if err != nil {
		return err
	}
	var dataBytes unsafe.Pointer
	if len(data) > 0 {
		dataBytes = unsafe.Pointer(&data[0])
	} else {
		dataBytes = unsafe.Pointer(&_zero)
	}
	_, _, errno := syscall.Syscall6(syscall.SYS_LSETXATTR, uintptr(unsafe.Pointer(pathBytes)), uintptr(unsafe.Pointer(attrBytes)), uintptr(dataBytes), uintptr(len(data)), uintptr(flags), 0)
	if errno != 0 {
		return errno
	}
	return nil
}
=======
	// Start with a 128 length byte array
	dest := make([]byte, 128)
	sz, errno := unix.Lgetxattr(path, attr, dest)

	switch {
	case errno == unix.ENODATA:
		return nil, errno
	case errno == unix.ENOTSUP:
		return nil, errno
	case errno == unix.ERANGE:
		// 128 byte array might just not be good enough,
		// A dummy buffer is used to get the real size
		// of the xattrs on disk
		sz, errno = unix.Lgetxattr(path, attr, []byte{})
		if errno != nil {
			return nil, errno
		}
		dest = make([]byte, sz)
		sz, errno = unix.Lgetxattr(path, attr, dest)
		if errno != nil {
			return nil, errno
		}
	case errno != nil:
		return nil, errno
	}
	return dest[:sz], nil
}
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
