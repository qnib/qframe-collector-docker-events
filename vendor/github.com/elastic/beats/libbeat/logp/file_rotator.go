package logp

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
<<<<<<< HEAD
=======
	"sync"
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
)

const RotatorMaxFiles = 1024
const DefaultKeepFiles = 7
const DefaultRotateEveryBytes = 10 * 1024 * 1024

type FileRotator struct {
	Path             string
	Name             string
	RotateEveryBytes *uint64
	KeepFiles        *int
<<<<<<< HEAD

	current     *os.File
	currentSize uint64
=======
	Permissions      *uint32

	current     *os.File
	currentSize uint64
	currentLock sync.RWMutex
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
}

func (rotator *FileRotator) CreateDirectory() error {
	fileinfo, err := os.Stat(rotator.Path)
	if err == nil {
		if !fileinfo.IsDir() {
			return fmt.Errorf("%s exists but it's not a directory", rotator.Path)
		}
	}

	if os.IsNotExist(err) {
		err = os.MkdirAll(rotator.Path, 0750)
		if err != nil {
			return err
		}
	}

	return nil
}

func (rotator *FileRotator) CheckIfConfigSane() error {
	if len(rotator.Name) == 0 {
<<<<<<< HEAD
		return fmt.Errorf("File logging requires a name for the file names")
=======
		return fmt.Errorf("file logging requires a name for the file names")
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	}
	if rotator.KeepFiles == nil {
		rotator.KeepFiles = new(int)
		*rotator.KeepFiles = DefaultKeepFiles
	}
	if rotator.RotateEveryBytes == nil {
		rotator.RotateEveryBytes = new(uint64)
		*rotator.RotateEveryBytes = DefaultRotateEveryBytes
	}

	if *rotator.KeepFiles < 2 || *rotator.KeepFiles >= RotatorMaxFiles {
<<<<<<< HEAD
		return fmt.Errorf("The number of files to keep should be between 2 and %d", RotatorMaxFiles-1)
=======
		return fmt.Errorf("the number of files to keep should be between 2 and %d", RotatorMaxFiles-1)
	}

	if rotator.Permissions != nil && (*rotator.Permissions > uint32(os.ModePerm)) {
		return fmt.Errorf("the permissions mask %d is invalid", *rotator.Permissions)
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	}
	return nil
}

func (rotator *FileRotator) WriteLine(line []byte) error {
	if rotator.shouldRotate() {
		err := rotator.Rotate()
		if err != nil {
			return err
		}
	}

	line = append(line, '\n')
<<<<<<< HEAD
	_, err := rotator.current.Write(line)
	if err != nil {
		return err
	}
	rotator.currentSize += uint64(len(line))
=======

	rotator.currentLock.RLock()
	_, err := rotator.current.Write(line)
	rotator.currentLock.RUnlock()

	if err != nil {
		return err
	}

	rotator.currentLock.Lock()
	rotator.currentSize += uint64(len(line))
	rotator.currentLock.Unlock()
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2

	return nil
}

func (rotator *FileRotator) shouldRotate() bool {
<<<<<<< HEAD
=======
	rotator.currentLock.RLock()
	defer rotator.currentLock.RUnlock()

>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	if rotator.current == nil {
		return true
	}

	if rotator.currentSize >= *rotator.RotateEveryBytes {
		return true
	}

	return false
}

func (rotator *FileRotator) FilePath(fileNo int) string {
	if fileNo == 0 {
		return filepath.Join(rotator.Path, rotator.Name)
	}
	filename := strings.Join([]string{rotator.Name, strconv.Itoa(fileNo)}, ".")
	return filepath.Join(rotator.Path, filename)
}

func (rotator *FileRotator) FileExists(fileNo int) bool {
	path := rotator.FilePath(fileNo)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func (rotator *FileRotator) Rotate() error {
<<<<<<< HEAD
=======
	rotator.currentLock.Lock()
	defer rotator.currentLock.Unlock()
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2

	if rotator.current != nil {
		if err := rotator.current.Close(); err != nil {
			return err
		}
	}

	// delete any extra files, normally we shouldn't have any
	for fileNo := *rotator.KeepFiles; fileNo < RotatorMaxFiles; fileNo++ {
		if rotator.FileExists(fileNo) {
			perr := os.Remove(rotator.FilePath(fileNo))
			if perr != nil {
				return perr
			}
		}
	}

	// shift all files from last to first
	for fileNo := *rotator.KeepFiles - 1; fileNo >= 0; fileNo-- {
		if !rotator.FileExists(fileNo) {
			// file doesn't exist, don't rotate
			continue
		}
		path := rotator.FilePath(fileNo)

		if rotator.FileExists(fileNo + 1) {
			// next file exists, something is strange
<<<<<<< HEAD
			return fmt.Errorf("File %s exists, when rotating would overwrite it", rotator.FilePath(fileNo+1))
=======
			return fmt.Errorf("file %s exists, when rotating would overwrite it", rotator.FilePath(fileNo+1))
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
		}

		err := os.Rename(path, rotator.FilePath(fileNo+1))
		if err != nil {
			return err
		}
	}

	// create the new file
	path := rotator.FilePath(0)
<<<<<<< HEAD
	current, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
=======
	current, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(rotator.getPermissions()))
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	if err != nil {
		return err
	}
	rotator.current = current
	rotator.currentSize = 0

	// delete the extra file, ignore errors here
	path = rotator.FilePath(*rotator.KeepFiles)
	os.Remove(path)

	return nil
}
<<<<<<< HEAD
=======

func (rotator *FileRotator) getPermissions() uint32 {
	if rotator.Permissions == nil {
		return 0600
	}
	return *rotator.Permissions
}
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
