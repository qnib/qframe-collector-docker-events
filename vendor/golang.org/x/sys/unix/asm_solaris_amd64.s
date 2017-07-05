// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !gccgo

#include "textflag.h"

//
// System calls for amd64, Solaris are implemented in runtime/syscall_solaris.go
//

<<<<<<< HEAD
TEXT ·sysvicall6(SB),NOSPLIT,$0-64
	JMP	syscall·sysvicall6(SB)

TEXT ·rawSysvicall6(SB),NOSPLIT,$0-64
=======
TEXT ·sysvicall6(SB),NOSPLIT,$0-88
	JMP	syscall·sysvicall6(SB)

TEXT ·rawSysvicall6(SB),NOSPLIT,$0-88
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	JMP	syscall·rawSysvicall6(SB)
