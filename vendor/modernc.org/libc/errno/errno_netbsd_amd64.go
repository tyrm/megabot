// Code generated by 'ccgo errno/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -o errno/errno_netbsd_amd64.go -pkgname errno', DO NOT EDIT.

package errno

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
	E2BIG              = 7
	EACCES             = 13
	EADDRINUSE         = 48
	EADDRNOTAVAIL      = 49
	EAFNOSUPPORT       = 47
	EAGAIN             = 35
	EALREADY           = 37
	EAUTH              = 80
	EBADF              = 9
	EBADMSG            = 88
	EBADRPC            = 72
	EBUSY              = 16
	ECANCELED          = 87
	ECHILD             = 10
	ECONNABORTED       = 53
	ECONNREFUSED       = 61
	ECONNRESET         = 54
	EDEADLK            = 11
	EDESTADDRREQ       = 39
	EDOM               = 33
	EDQUOT             = 69
	EEXIST             = 17
	EFAULT             = 14
	EFBIG              = 27
	EFTYPE             = 79
	EHOSTDOWN          = 64
	EHOSTUNREACH       = 65
	EIDRM              = 82
	EILSEQ             = 85
	EINPROGRESS        = 36
	EINTR              = 4
	EINVAL             = 22
	EIO                = 5
	EISCONN            = 56
	EISDIR             = 21
	ELAST              = 96
	ELOOP              = 62
	EMFILE             = 24
	EMLINK             = 31
	EMSGSIZE           = 40
	EMULTIHOP          = 94
	ENAMETOOLONG       = 63
	ENEEDAUTH          = 81
	ENETDOWN           = 50
	ENETRESET          = 52
	ENETUNREACH        = 51
	ENFILE             = 23
	ENOATTR            = 93
	ENOBUFS            = 55
	ENODATA            = 89
	ENODEV             = 19
	ENOENT             = 2
	ENOEXEC            = 8
	ENOLCK             = 77
	ENOLINK            = 95
	ENOMEM             = 12
	ENOMSG             = 83
	ENOPROTOOPT        = 42
	ENOSPC             = 28
	ENOSR              = 90
	ENOSTR             = 91
	ENOSYS             = 78
	ENOTBLK            = 15
	ENOTCONN           = 57
	ENOTDIR            = 20
	ENOTEMPTY          = 66
	ENOTSOCK           = 38
	ENOTSUP            = 86
	ENOTTY             = 25
	ENXIO              = 6
	EOPNOTSUPP         = 45
	EOVERFLOW          = 84
	EPERM              = 1
	EPFNOSUPPORT       = 46
	EPIPE              = 32
	EPROCLIM           = 67
	EPROCUNAVAIL       = 76
	EPROGMISMATCH      = 75
	EPROGUNAVAIL       = 74
	EPROTO             = 96
	EPROTONOSUPPORT    = 43
	EPROTOTYPE         = 41
	ERANGE             = 34
	EREMOTE            = 71
	EROFS              = 30
	ERPCMISMATCH       = 73
	ESHUTDOWN          = 58
	ESOCKTNOSUPPORT    = 44
	ESPIPE             = 29
	ESRCH              = 3
	ESTALE             = 70
	ETIME              = 92
	ETIMEDOUT          = 60
	ETOOMANYREFS       = 59
	ETXTBSY            = 26
	EUSERS             = 68
	EWOULDBLOCK        = 35
	EXDEV              = 18
	X_ERRNO_H_         = 0
	X_FILE_OFFSET_BITS = 64
	X_LP64             = 1
	X_NETBSD_SOURCE    = 1
	X_SYS_CDEFS_ELF_H_ = 0
	X_SYS_CDEFS_H_     = 0
	X_SYS_ERRNO_H_     = 0
	X_X86_64_CDEFS_H_  = 0
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

var _ int8 /* gen.c:2:13: */
