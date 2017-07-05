package qutils

import (
	"bytes"
	"runtime"
	"strconv"
)


<<<<<<< HEAD
func GetGID() uint64 {
=======
func GetGID() int {
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
<<<<<<< HEAD
	n, _ := strconv.ParseUint(string(b), 10, 64)
=======
	n, _ := strconv.Atoi(string(b))
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	return n
}

