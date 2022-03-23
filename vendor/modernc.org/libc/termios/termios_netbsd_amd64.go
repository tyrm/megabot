// Code generated by 'ccgo termios/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -o termios/termios_netbsd_amd64.go -pkgname termios', DO NOT EDIT.

package termios

import (
	"math"
	"reflect"
	"sync/atomic"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ atomic.Value
var _ unsafe.Pointer

const (
	ALTWERASE                 = 0x00000200
	ARG_MAX                   = 262144
	B0                        = 0
	B1000000                  = 1000000
	B110                      = 110
	B115200                   = 115200
	B1200                     = 1200
	B134                      = 134
	B14400                    = 14400
	B150                      = 150
	B1500000                  = 1500000
	B1800                     = 1800
	B19200                    = 19200
	B200                      = 200
	B2000000                  = 2000000
	B230400                   = 230400
	B2400                     = 2400
	B2500000                  = 2500000
	B28800                    = 28800
	B300                      = 300
	B3000000                  = 3000000
	B3500000                  = 3500000
	B38400                    = 38400
	B4000000                  = 4000000
	B460800                   = 460800
	B4800                     = 4800
	B50                       = 50
	B500000                   = 500000
	B57600                    = 57600
	B600                      = 600
	B7200                     = 7200
	B75                       = 75
	B76800                    = 76800
	B921600                   = 921600
	B9600                     = 9600
	BC_DIM_MAX                = 65535
	BRKINT                    = 0x00000002
	CCTS_OFLOW                = 65536
	CDISCARD                  = 15
	CDSUSP                    = 25
	CDTRCTS                   = 0x00020000
	CEOF                      = 4
	CEOT                      = 4
	CERASE                    = 0177
	CFLUSH                    = 15
	CHILD_MAX                 = 160
	CHWFLOW                   = 1245184
	CIGNORE                   = 0x00000001
	CINTR                     = 3
	CKILL                     = 21
	CLNEXT                    = 22
	CLOCAL                    = 0x00008000
	CMIN                      = 1
	COLL_WEIGHTS_MAX          = 2
	CQUIT                     = 034
	CREAD                     = 0x00000800
	CREPRINT                  = 18
	CRPRNT                    = 18
	CRTSCTS                   = 0x00010000
	CRTS_IFLOW                = 65536
	CS5                       = 0x00000000
	CS6                       = 0x00000100
	CS7                       = 0x00000200
	CS8                       = 0x00000300
	CSIZE                     = 0x00000300
	CSTART                    = 17
	CSTATUS                   = 20
	CSTOP                     = 19
	CSTOPB                    = 0x00000400
	CSUSP                     = 26
	CTIME                     = 0
	CWERASE                   = 23
	ECHO                      = 0x00000008
	ECHOCTL                   = 0x00000040
	ECHOE                     = 0x00000002
	ECHOK                     = 0x00000004
	ECHOKE                    = 0x00000001
	ECHONL                    = 0x00000010
	ECHOPRT                   = 0x00000020
	EXPR_NEST_MAX             = 32
	EXTA                      = 19200
	EXTB                      = 38400
	EXTPROC                   = 0x00000800
	FLUSHO                    = 0x00800000
	GID_MAX                   = 2147483647
	HDLCDISC                  = 9
	HUPCL                     = 0x00004000
	ICANON                    = 0x00000100
	ICRNL                     = 0x00000100
	IEXTEN                    = 0x00000400
	IGNBRK                    = 0x00000001
	IGNCR                     = 0x00000080
	IGNPAR                    = 0x00000004
	IMAXBEL                   = 0x00002000
	INLCR                     = 0x00000040
	INPCK                     = 0x00000010
	IOCGROUP_SHIFT            = 8
	IOCPARM_MASK              = 0x1fff
	IOCPARM_SHIFT             = 16
	IOV_MAX                   = 1024
	ISIG                      = 0x00000080
	ISTRIP                    = 0x00000020
	IXANY                     = 0x00000800
	IXOFF                     = 0x00000400
	IXON                      = 0x00000200
	LINE_MAX                  = 2048
	LINK_MAX                  = 32767
	LOGIN_NAME_MAX            = 17
	MAX_CANON                 = 255
	MAX_INPUT                 = 255
	MDMBUF                    = 0x00100000
	NAME_MAX                  = 511
	NCCS                      = 20
	NGROUPS_MAX               = 16
	NOFLSH                    = 0x80000000
	NOKERNINFO                = 0x02000000
	NZERO                     = 20
	OCRNL                     = 0x00000010
	ONLCR                     = 0x00000002
	ONLRET                    = 0x00000040
	ONOCR                     = 0x00000020
	ONOEOT                    = 0x00000008
	OPEN_MAX                  = 128
	OPOST                     = 0x00000001
	OXTABS                    = 0x00000004
	PARENB                    = 0x00001000
	PARMRK                    = 0x00000008
	PARODD                    = 0x00002000
	PATH_MAX                  = 1024
	PENDIN                    = 0x20000000
	PIPE_BUF                  = 512
	PPPDISC                   = 5
	RE_DUP_MAX                = 255
	SLIPDISC                  = 4
	STRIPDISC                 = 6
	TABLDISC                  = 3
	TCIFLUSH                  = 1
	TCIOFF                    = 3
	TCIOFLUSH                 = 3
	TCION                     = 4
	TCOFLUSH                  = 2
	TCOOFF                    = 1
	TCOON                     = 2
	TCSADRAIN                 = 1
	TCSAFLUSH                 = 2
	TCSANOW                   = 0
	TCSASOFT                  = 0x10
	TIOCFLAG_CDTRCTS          = 0x10
	TIOCFLAG_CLOCAL           = 0x02
	TIOCFLAG_CRTSCTS          = 0x04
	TIOCFLAG_MDMBUF           = 0x08
	TIOCFLAG_SOFTCAR          = 0x01
	TIOCM_CAR                 = 0100
	TIOCM_CD                  = 64
	TIOCM_CTS                 = 0040
	TIOCM_DSR                 = 0400
	TIOCM_DTR                 = 0002
	TIOCM_LE                  = 0001
	TIOCM_RI                  = 128
	TIOCM_RNG                 = 0200
	TIOCM_RTS                 = 0004
	TIOCM_SR                  = 0020
	TIOCM_ST                  = 0010
	TIOCPKT_DATA              = 0x00
	TIOCPKT_DOSTOP            = 0x20
	TIOCPKT_FLUSHREAD         = 0x01
	TIOCPKT_FLUSHWRITE        = 0x02
	TIOCPKT_IOCTL             = 0x40
	TIOCPKT_NOSTOP            = 0x10
	TIOCPKT_START             = 0x08
	TIOCPKT_STOP              = 0x04
	TOSTOP                    = 0x00400000
	TTLINEDNAMELEN            = 32
	TTYDEF_CFLAG              = 19200
	TTYDEF_IFLAG              = 11010
	TTYDEF_LFLAG              = 1483
	TTYDEF_OFLAG              = 7
	TTYDEF_SPEED              = 9600
	TTYDISC                   = 0
	UID_MAX                   = 2147483647
	VDISCARD                  = 15
	VDSUSP                    = 11
	VEOF                      = 0
	VEOL                      = 1
	VEOL2                     = 2
	VERASE                    = 3
	VINTR                     = 8
	VKILL                     = 5
	VLNEXT                    = 14
	VMIN                      = 16
	VQUIT                     = 9
	VREPRINT                  = 6
	VSTART                    = 12
	VSTATUS                   = 18
	VSTOP                     = 13
	VSUSP                     = 10
	VTIME                     = 17
	VWERASE                   = 4
	X_AMD64_INT_TYPES_H_      = 0
	X_FILE_OFFSET_BITS        = 64
	X_LP64                    = 1
	X_NETBSD_SOURCE           = 1
	X_NETBSD_SYS_TTYCOM_H_    = 0
	X_PATH_PTMDEV             = "/dev/ptm"
	X_POSIX_SYS_TTYCOM_H_     = 0
	X_SYS_ANSI_H_             = 0
	X_SYS_CDEFS_ELF_H_        = 0
	X_SYS_CDEFS_H_            = 0
	X_SYS_COMMON_ANSI_H_      = 0
	X_SYS_COMMON_INT_TYPES_H_ = 0
	X_SYS_IOCCOM_H_           = 0
	X_SYS_SYSLIMITS_H_        = 0
	X_SYS_TERMIOS_H_          = 0
	X_SYS_TTYDEFAULTS_H_      = 0
	X_X86_64_CDEFS_H_         = 0
)

type Ptrdiff_t = int64 /* <builtin>:3:26 */

type Size_t = uint64 /* <builtin>:9:23 */

type Wchar_t = int32 /* <builtin>:15:24 */

type X__int128_t = struct {
	Flo int64
	Fhi int64
} /* <builtin>:21:43 */ // must match modernc.org/mathutil.Int128
type X__uint128_t = struct {
	Flo uint64
	Fhi uint64
} /* <builtin>:22:44 */ // must match modernc.org/mathutil.Int128

type X__builtin_va_list = uintptr /* <builtin>:46:14 */
type X__float128 = float64        /* <builtin>:47:21 */

// return true if value 'a' fits in type 't'

//	$NetBSD: int_types.h,v 1.7 2014/07/25 21:43:13 joerg Exp $

// -
// Copyright (c) 1990 The Regents of the University of California.
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
// 3. Neither the name of the University nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
// OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
// OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
// SUCH DAMAGE.
//
//	from: @(#)types.h	7.5 (Berkeley) 3/9/91

//	$NetBSD: common_int_types.h,v 1.1 2014/07/25 21:43:13 joerg Exp $

// -
// Copyright (c) 2014 The NetBSD Foundation, Inc.
// All rights reserved.
//
// This code is derived from software contributed to The NetBSD Foundation
// by Joerg Sonnenberger.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE NETBSD FOUNDATION, INC. AND CONTRIBUTORS
// ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED
// TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
// PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL THE FOUNDATION OR CONTRIBUTORS
// BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

// 7.18.1 Integer types

// 7.18.1.1 Exact-width integer types

type X__int8_t = int8     /* common_int_types.h:45:27 */
type X__uint8_t = uint8   /* common_int_types.h:46:27 */
type X__int16_t = int16   /* common_int_types.h:47:27 */
type X__uint16_t = uint16 /* common_int_types.h:48:27 */
type X__int32_t = int32   /* common_int_types.h:49:27 */
type X__uint32_t = uint32 /* common_int_types.h:50:27 */
type X__int64_t = int64   /* common_int_types.h:51:27 */
type X__uint64_t = uint64 /* common_int_types.h:52:27 */

// 7.18.1.4 Integer types capable of holding object pointers

type X__intptr_t = int64   /* common_int_types.h:58:27 */
type X__uintptr_t = uint64 /* common_int_types.h:59:26 */

// Types which are fundamental to the implementation and may appear in
// more than one standard header are defined here.  Standard headers
// then use:
//	#ifdef	_BSD_SIZE_T_
//	typedef	_BSD_SIZE_T_ size_t;
//	#undef	_BSD_SIZE_T_
//	#endif

type X__caddr_t = uintptr        /* ansi.h:37:14 */ // core address
type X__gid_t = X__uint32_t      /* ansi.h:38:20 */ // group id
type X__in_addr_t = X__uint32_t  /* ansi.h:39:20 */ // IP(v4) address
type X__in_port_t = X__uint16_t  /* ansi.h:40:20 */ // "Internet" port number
type X__mode_t = X__uint32_t     /* ansi.h:41:20 */ // file permissions
type X__off_t = X__int64_t       /* ansi.h:42:19 */ // file offset
type X__pid_t = X__int32_t       /* ansi.h:43:19 */ // process id
type X__sa_family_t = X__uint8_t /* ansi.h:44:19 */ // socket address family
type X__socklen_t = uint32       /* ansi.h:45:22 */ // socket-related datum length
type X__uid_t = X__uint32_t      /* ansi.h:46:20 */ // user id
type X__fsblkcnt_t = X__uint64_t /* ansi.h:47:20 */ // fs block count (statvfs)
type X__fsfilcnt_t = X__uint64_t /* ansi.h:48:20 */
type X__wctrans_t = uintptr      /* ansi.h:51:32 */
type X__wctype_t = uintptr       /* ansi.h:54:31 */

// mbstate_t is an opaque object to keep conversion state, during multibyte
// stream conversions.  The content must not be referenced by user programs.
type X__mbstate_t = struct {
	F__mbstateL  X__int64_t
	F__ccgo_pad1 [120]byte
} /* ansi.h:63:3 */

type X__va_list = X__builtin_va_list /* ansi.h:72:27 */

//	$NetBSD: featuretest.h,v 1.10 2013/04/26 18:29:06 christos Exp $

// Written by Klaus Klein <kleink@NetBSD.org>, February 2, 1998.
// Public domain.
//
// NOTE: Do not protect this header against multiple inclusion.  Doing
// so can have subtle side-effects due to header file inclusion order
// and testing of e.g. _POSIX_SOURCE vs. _POSIX_C_SOURCE.  Instead,
// protect each CPP macro that we want to supply.

// Feature-test macros are defined by several standards, and allow an
// application to specify what symbols they want the system headers to
// expose, and hence what standard they want them to conform to.
// There are two classes of feature-test macros.  The first class
// specify complete standards, and if one of these is defined, header
// files will try to conform to the relevant standard.  They are:
//
// ANSI macros:
// _ANSI_SOURCE			ANSI C89
//
// POSIX macros:
// _POSIX_SOURCE == 1		IEEE Std 1003.1 (version?)
// _POSIX_C_SOURCE == 1		IEEE Std 1003.1-1990
// _POSIX_C_SOURCE == 2		IEEE Std 1003.2-1992
// _POSIX_C_SOURCE == 199309L	IEEE Std 1003.1b-1993
// _POSIX_C_SOURCE == 199506L	ISO/IEC 9945-1:1996
// _POSIX_C_SOURCE == 200112L	IEEE Std 1003.1-2001
// _POSIX_C_SOURCE == 200809L   IEEE Std 1003.1-2008
//
// X/Open macros:
// _XOPEN_SOURCE		System Interfaces and Headers, Issue 4, Ver 2
// _XOPEN_SOURCE_EXTENDED == 1	XSH4.2 UNIX extensions
// _XOPEN_SOURCE == 500		System Interfaces and Headers, Issue 5
// _XOPEN_SOURCE == 520		Networking Services (XNS), Issue 5.2
// _XOPEN_SOURCE == 600		IEEE Std 1003.1-2001, XSI option
// _XOPEN_SOURCE == 700		IEEE Std 1003.1-2008, XSI option
//
// NetBSD macros:
// _NETBSD_SOURCE == 1		Make all NetBSD features available.
//
// If more than one of these "major" feature-test macros is defined,
// then the set of facilities provided (and namespace used) is the
// union of that specified by the relevant standards, and in case of
// conflict, the earlier standard in the above list has precedence (so
// if both _POSIX_C_SOURCE and _NETBSD_SOURCE are defined, the version
// of rename() that's used is the POSIX one).  If none of the "major"
// feature-test macros is defined, _NETBSD_SOURCE is assumed.
//
// There are also "minor" feature-test macros, which enable extra
// functionality in addition to some base standard.  They should be
// defined along with one of the "major" macros.  The "minor" macros
// are:
//
// _REENTRANT
// _ISOC99_SOURCE
// _ISOC11_SOURCE
// _LARGEFILE_SOURCE		Large File Support
//		<http://ftp.sas.com/standards/large.file/x_open.20Mar96.html>

// Special Control Characters
//
// Index into c_cc[] character array.
//
//	Name	     Subscript	Enabled by
//			7	   spare 1
//			19	   spare 2

// Input flags - software input processing

// Output flags - software output processing

// Control flags - hardware control of terminal

// "Local" flags - dumping ground for other state
//
// Warning: some flags in this structure begin with
// the letter "I" and look like they belong in the
// input flag.

type Tcflag_t = uint32 /* termios.h:188:22 */
type Cc_t = uint8      /* termios.h:189:23 */
type Speed_t = uint32  /* termios.h:190:22 */

type Termios = struct {
	Fc_iflag  Tcflag_t
	Fc_oflag  Tcflag_t
	Fc_cflag  Tcflag_t
	Fc_lflag  Tcflag_t
	Fc_cc     [20]Cc_t
	Fc_ispeed int32
	Fc_ospeed int32
} /* termios.h:192:1 */

// Commands passed to tcsetattr() for setting the termios structure.

// Standard speeds

type Pid_t = X__pid_t /* termios.h:265:18 */

// Include tty ioctl's that aren't just for backwards compatibility
// with the old tty driver.  These ioctl definitions were previously
// in <sys/ioctl.h>.   Most of this appears only when _NETBSD_SOURCE
// is defined, but (at least) struct winsize has been made standard,
// and needs to be visible here (as well as via the old <sys/ioctl.h>.)
//	$NetBSD: ttycom.h,v 1.21 2017/10/25 06:32:59 kre Exp $

// -
// Copyright (c) 1982, 1986, 1990, 1993, 1994
//	The Regents of the University of California.  All rights reserved.
// (c) UNIX System Laboratories, Inc.
// All or some portions of this file are derived from material licensed
// to the University of California by American Telephone and Telegraph
// Co. or Unix System Laboratories, Inc. and are reproduced herein with
// the permission of UNIX System Laboratories, Inc.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
// 3. Neither the name of the University nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
// OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
// OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
// SUCH DAMAGE.
//
//	@(#)ttycom.h	8.1 (Berkeley) 3/28/94

//	$NetBSD: syslimits.h,v 1.28 2015/08/21 07:19:39 uebayasi Exp $

// Copyright (c) 1988, 1993
//	The Regents of the University of California.  All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
// 3. Neither the name of the University nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
// OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
// OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
// SUCH DAMAGE.
//
//	@(#)syslimits.h	8.1 (Berkeley) 6/2/93

//	$NetBSD: featuretest.h,v 1.10 2013/04/26 18:29:06 christos Exp $

// Written by Klaus Klein <kleink@NetBSD.org>, February 2, 1998.
// Public domain.
//
// NOTE: Do not protect this header against multiple inclusion.  Doing
// so can have subtle side-effects due to header file inclusion order
// and testing of e.g. _POSIX_SOURCE vs. _POSIX_C_SOURCE.  Instead,
// protect each CPP macro that we want to supply.

// Feature-test macros are defined by several standards, and allow an
// application to specify what symbols they want the system headers to
// expose, and hence what standard they want them to conform to.
// There are two classes of feature-test macros.  The first class
// specify complete standards, and if one of these is defined, header
// files will try to conform to the relevant standard.  They are:
//
// ANSI macros:
// _ANSI_SOURCE			ANSI C89
//
// POSIX macros:
// _POSIX_SOURCE == 1		IEEE Std 1003.1 (version?)
// _POSIX_C_SOURCE == 1		IEEE Std 1003.1-1990
// _POSIX_C_SOURCE == 2		IEEE Std 1003.2-1992
// _POSIX_C_SOURCE == 199309L	IEEE Std 1003.1b-1993
// _POSIX_C_SOURCE == 199506L	ISO/IEC 9945-1:1996
// _POSIX_C_SOURCE == 200112L	IEEE Std 1003.1-2001
// _POSIX_C_SOURCE == 200809L   IEEE Std 1003.1-2008
//
// X/Open macros:
// _XOPEN_SOURCE		System Interfaces and Headers, Issue 4, Ver 2
// _XOPEN_SOURCE_EXTENDED == 1	XSH4.2 UNIX extensions
// _XOPEN_SOURCE == 500		System Interfaces and Headers, Issue 5
// _XOPEN_SOURCE == 520		Networking Services (XNS), Issue 5.2
// _XOPEN_SOURCE == 600		IEEE Std 1003.1-2001, XSI option
// _XOPEN_SOURCE == 700		IEEE Std 1003.1-2008, XSI option
//
// NetBSD macros:
// _NETBSD_SOURCE == 1		Make all NetBSD features available.
//
// If more than one of these "major" feature-test macros is defined,
// then the set of facilities provided (and namespace used) is the
// union of that specified by the relevant standards, and in case of
// conflict, the earlier standard in the above list has precedence (so
// if both _POSIX_C_SOURCE and _NETBSD_SOURCE are defined, the version
// of rename() that's used is the POSIX one).  If none of the "major"
// feature-test macros is defined, _NETBSD_SOURCE is assumed.
//
// There are also "minor" feature-test macros, which enable extra
// functionality in addition to some base standard.  They should be
// defined along with one of the "major" macros.  The "minor" macros
// are:
//
// _REENTRANT
// _ISOC99_SOURCE
// _ISOC11_SOURCE
// _LARGEFILE_SOURCE		Large File Support
//		<http://ftp.sas.com/standards/large.file/x_open.20Mar96.html>

// kept in sync with MAXNAMLEN

// IEEE Std 1003.1c-95, adopted in X/Open CAE Specification Issue 5 Version 2

// X/Open CAE Specification Issue 5 Version 2

//	$NetBSD: ioccom.h,v 1.13 2019/05/26 10:21:33 hannken Exp $

// -
// Copyright (c) 1982, 1986, 1990, 1993, 1994
//	The Regents of the University of California.  All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
// 3. Neither the name of the University nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
// OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
// OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
// SUCH DAMAGE.
//
//	@(#)ioccom.h	8.3 (Berkeley) 1/9/95

// Ioctl's have the command encoded in the lower word, and the size of
// any in or out parameters in the upper word.  The high 3 bits of the
// upper word are used to encode the in/out status of the parameter.
//
//	 31 29 28                     16 15            8 7             0
//	+---------------------------------------------------------------+
//	| I/O | Parameter Length        | Command Group | Command       |
//	+---------------------------------------------------------------+

// no parameters
// copy parameters out
// copy parameters in
// copy parameters in and out
// mask for IN/OUT/VOID

// this should be _IORW, but stdio got there first

// Tty ioctl's except for those supported only for backwards compatibility
// with the old tty driver.

// Window/terminal size structure.  This information is stored by the kernel
// in order to provide a consistent interface, but is not used by the kernel.
type Winsize = struct {
	Fws_row    uint16
	Fws_col    uint16
	Fws_xpixel uint16
	Fws_ypixel uint16
} /* ttycom.h:54:1 */

// The following are not exposed when imported via <termios.h>
// when _POSIX_SOURCE (et.al.) is defined (and hence _NETBSD_SOURCE
// is not, unless that is added manually.)

// ptmget, for /dev/ptm pty getting ioctl TIOCPTMGET, and for TIOCPTSNAME
type Ptmget = struct {
	Fcfd int32
	Fsfd int32
	Fcn  [1024]int8
	Fsn  [1024]int8
} /* ttycom.h:74:1 */

// 8-10 compat
// 15 unused
// 17-18 compat

// This is the maximum length of a line discipline's name.
type Linedn_t = [32]int8 /* ttycom.h:111:14 */

// END OF PROTECTED INCLUDE.

//	$NetBSD: ttydefaults.h,v 1.16 2008/05/24 14:06:39 yamt Exp $

// -
// Copyright (c) 1982, 1986, 1993
//	The Regents of the University of California.  All rights reserved.
// (c) UNIX System Laboratories, Inc.
// All or some portions of this file are derived from material licensed
// to the University of California by American Telephone and Telegraph
// Co. or Unix System Laboratories, Inc. and are reproduced herein with
// the permission of UNIX System Laboratories, Inc.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
// 3. Neither the name of the University nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
// OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
// OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
// SUCH DAMAGE.
//
//	@(#)ttydefaults.h	8.4 (Berkeley) 1/21/94

// System wide defaults for terminal state.

// Defaults on "first" open.

// Control Character Defaults
// compat

// PROTECTED INCLUSION ENDS HERE

// #define TTYDEFCHARS to include an array of default control characters.
var _ int8 /* gen.c:2:13: */
