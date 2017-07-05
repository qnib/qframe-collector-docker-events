// +build windows

package idtools

import (
	"os"

	"github.com/docker/docker/pkg/system"
)

// Platforms such as Windows do not support the UID/GID concept. So make this
// just a wrapper around system.MkdirAll.
func mkdirAs(path string, mode os.FileMode, ownerUID, ownerGID int, mkAll, chownExisting bool) error {
<<<<<<< HEAD
	if err := system.MkdirAll(path, mode); err != nil && !os.IsExist(err) {
=======
	if err := system.MkdirAll(path, mode, ""); err != nil && !os.IsExist(err) {
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
		return err
	}
	return nil
}

// CanAccess takes a valid (existing) directory and a uid, gid pair and determines
// if that uid, gid pair has access (execute bit) to the directory
// Windows does not require/support this function, so always return true
<<<<<<< HEAD
func CanAccess(path string, uid, gid int) bool {
=======
func CanAccess(path string, pair IDPair) bool {
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	return true
}
