// +build !windows

package system

<<<<<<< HEAD
// DefaultPathEnv is unix style list of directories to search for
// executables. Each directory is separated from the next by a colon
// ':' character .
const DefaultPathEnv = "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"

=======
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
// CheckSystemDriveAndRemoveDriveLetter verifies that a path, if it includes a drive letter,
// is the system drive. This is a no-op on Linux.
func CheckSystemDriveAndRemoveDriveLetter(path string) (string, error) {
	return path, nil
}
