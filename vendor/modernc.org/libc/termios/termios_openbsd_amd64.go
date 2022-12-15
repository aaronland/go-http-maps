// Code generated by 'ccgo termios/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -header -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -ignore-unsupported-alignment -o termios/termios_openbsd_amd64.go -pkgname termios', DO NOT EDIT.

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
	ALTWERASE            = 0x00000200 // termios.h:172:1:
	B0                   = 0          // termios.h:212:1:
	B110                 = 110        // termios.h:215:1:
	B115200              = 115200     // termios.h:234:1:
	B1200                = 1200       // termios.h:221:1:
	B134                 = 134        // termios.h:216:1:
	B14400               = 14400      // termios.h:230:1:
	B150                 = 150        // termios.h:217:1:
	B1800                = 1800       // termios.h:222:1:
	B19200               = 19200      // termios.h:226:1:
	B200                 = 200        // termios.h:218:1:
	B230400              = 230400     // termios.h:235:1:
	B2400                = 2400       // termios.h:223:1:
	B28800               = 28800      // termios.h:231:1:
	B300                 = 300        // termios.h:219:1:
	B38400               = 38400      // termios.h:227:1:
	B4800                = 4800       // termios.h:224:1:
	B50                  = 50         // termios.h:213:1:
	B57600               = 57600      // termios.h:232:1:
	B600                 = 600        // termios.h:220:1:
	B7200                = 7200       // termios.h:229:1:
	B75                  = 75         // termios.h:214:1:
	B76800               = 76800      // termios.h:233:1:
	B9600                = 9600       // termios.h:225:1:
	BRKINT               = 0x00000002 // termios.h:91:1:
	CCTS_OFLOW           = 65536      // termios.h:145:1:
	CDISCARD             = 15         // ttydefaults.h:73:1:
	CDSUSP               = 25         // ttydefaults.h:69:1:
	CEOF                 = 4          // ttydefaults.h:59:1:
	CEOT                 = 4          // ttydefaults.h:76:1:
	CERASE               = 0177       // ttydefaults.h:61:1:
	CFLUSH               = 15         // ttydefaults.h:80:1:
	CHWFLOW              = 1114112    // termios.h:147:1:
	CIGNORE              = 0x00000001 // termios.h:129:1:
	CINTR                = 3          // ttydefaults.h:62:1:
	CKILL                = 21         // ttydefaults.h:64:1:
	CLNEXT               = 22         // ttydefaults.h:72:1:
	CLOCAL               = 0x00008000 // termios.h:141:1:
	CMIN                 = 1          // ttydefaults.h:65:1:
	CQUIT                = 034        // ttydefaults.h:66:1:
	CREAD                = 0x00000800 // termios.h:137:1:
	CREPRINT             = 18         // ttydefaults.h:75:1:
	CRPRNT               = 18         // ttydefaults.h:79:1:
	CRTSCTS              = 0x00010000 // termios.h:143:1:
	CRTS_IFLOW           = 65536      // termios.h:144:1:
	CS5                  = 0x00000000 // termios.h:132:1:
	CS6                  = 0x00000100 // termios.h:133:1:
	CS7                  = 0x00000200 // termios.h:134:1:
	CS8                  = 0x00000300 // termios.h:135:1:
	CSIZE                = 0x00000300 // termios.h:131:1:
	CSTART               = 17         // ttydefaults.h:70:1:
	CSTOP                = 19         // ttydefaults.h:71:1:
	CSTOPB               = 0x00000400 // termios.h:136:1:
	CSUSP                = 26         // ttydefaults.h:67:1:
	CTIME                = 0          // ttydefaults.h:68:1:
	CWERASE              = 23         // ttydefaults.h:74:1:
	ECHO                 = 0x00000008 // termios.h:163:1:
	ECHOCTL              = 0x00000040 // termios.h:167:1:
	ECHOE                = 0x00000002 // termios.h:161:1:
	ECHOK                = 0x00000004 // termios.h:162:1:
	ECHOKE               = 0x00000001 // termios.h:159:1:
	ECHONL               = 0x00000010 // termios.h:164:1:
	ECHOPRT              = 0x00000020 // termios.h:166:1:
	ENDRUNDISC           = 9          // ttycom.h:149:1:
	EXTA                 = 19200      // termios.h:236:1:
	EXTB                 = 38400      // termios.h:237:1:
	EXTPROC              = 0x00000800 // termios.h:175:1:
	FLUSHO               = 0x00800000 // termios.h:178:1:
	HUPCL                = 0x00004000 // termios.h:140:1:
	ICANON               = 0x00000100 // termios.h:170:1:
	ICRNL                = 0x00000100 // termios.h:98:1:
	IEXTEN               = 0x00000400 // termios.h:174:1:
	IGNBRK               = 0x00000001 // termios.h:90:1:
	IGNCR                = 0x00000080 // termios.h:97:1:
	IGNPAR               = 0x00000004 // termios.h:92:1:
	IMAXBEL              = 0x00002000 // termios.h:104:1:
	INLCR                = 0x00000040 // termios.h:96:1:
	INPCK                = 0x00000010 // termios.h:94:1:
	IOCPARM_MASK         = 0x1fff     // ioccom.h:43:1:
	ISIG                 = 0x00000080 // termios.h:169:1:
	ISTRIP               = 0x00000020 // termios.h:95:1:
	IUCLC                = 0x00001000 // termios.h:103:1:
	IXANY                = 0x00000800 // termios.h:102:1:
	IXOFF                = 0x00000400 // termios.h:100:1:
	IXON                 = 0x00000200 // termios.h:99:1:
	MDMBUF               = 0x00100000 // termios.h:146:1:
	MSTSDISC             = 8          // ttycom.h:148:1:
	NCCS                 = 20         // termios.h:79:1:
	NMEADISC             = 7          // ttycom.h:147:1:
	NOFLSH               = 0x80000000 // termios.h:183:1:
	NOKERNINFO           = 0x02000000 // termios.h:180:1:
	OCRNL                = 0x00000010 // termios.h:119:1:
	OLCUC                = 0x00000020 // termios.h:120:1:
	ONLCR                = 0x00000002 // termios.h:112:1:
	ONLRET               = 0x00000080 // termios.h:122:1:
	ONOCR                = 0x00000040 // termios.h:121:1:
	ONOEOT               = 0x00000008 // termios.h:116:1:
	OPOST                = 0x00000001 // termios.h:110:1:
	OXTABS               = 0x00000004 // termios.h:115:1:
	PARENB               = 0x00001000 // termios.h:138:1:
	PARMRK               = 0x00000008 // termios.h:93:1:
	PARODD               = 0x00002000 // termios.h:139:1:
	PENDIN               = 0x20000000 // termios.h:181:1:
	PPPDISC              = 5          // ttycom.h:145:1:
	SLIPDISC             = 4          // ttycom.h:144:1:
	STRIPDISC            = 6          // ttycom.h:146:1:
	TABLDISC             = 3          // ttycom.h:143:1:
	TCIFLUSH             = 1          // termios.h:242:1:
	TCIOFF               = 3          // termios.h:247:1:
	TCIOFLUSH            = 3          // termios.h:244:1:
	TCION                = 4          // termios.h:248:1:
	TCOFLUSH             = 2          // termios.h:243:1:
	TCOOFF               = 1          // termios.h:245:1:
	TCOON                = 2          // termios.h:246:1:
	TCSADRAIN            = 1          // termios.h:203:1:
	TCSAFLUSH            = 2          // termios.h:204:1:
	TCSANOW              = 0          // termios.h:202:1:
	TCSASOFT             = 0x10       // termios.h:206:1:
	TIOCFLAG_CLOCAL      = 0x02       // ttycom.h:131:1:
	TIOCFLAG_CRTSCTS     = 0x04       // ttycom.h:132:1:
	TIOCFLAG_MDMBUF      = 0x08       // ttycom.h:133:1:
	TIOCFLAG_PPS         = 0x10       // ttycom.h:134:1:
	TIOCFLAG_SOFTCAR     = 0x01       // ttycom.h:130:1:
	TIOCM_CAR            = 0100       // ttycom.h:69:1:
	TIOCM_CD             = 64         // ttycom.h:70:1:
	TIOCM_CTS            = 0040       // ttycom.h:68:1:
	TIOCM_DSR            = 0400       // ttycom.h:73:1:
	TIOCM_DTR            = 0002       // ttycom.h:64:1:
	TIOCM_LE             = 0001       // ttycom.h:63:1:
	TIOCM_RI             = 128        // ttycom.h:72:1:
	TIOCM_RNG            = 0200       // ttycom.h:71:1:
	TIOCM_RTS            = 0004       // ttycom.h:65:1:
	TIOCM_SR             = 0020       // ttycom.h:67:1:
	TIOCM_ST             = 0010       // ttycom.h:66:1:
	TIOCPKT_DATA         = 0x00       // ttycom.h:100:1:
	TIOCPKT_DOSTOP       = 0x20       // ttycom.h:106:1:
	TIOCPKT_FLUSHREAD    = 0x01       // ttycom.h:101:1:
	TIOCPKT_FLUSHWRITE   = 0x02       // ttycom.h:102:1:
	TIOCPKT_IOCTL        = 0x40       // ttycom.h:107:1:
	TIOCPKT_NOSTOP       = 0x10       // ttycom.h:105:1:
	TIOCPKT_START        = 0x08       // ttycom.h:104:1:
	TIOCPKT_STOP         = 0x04       // ttycom.h:103:1:
	TOSTOP               = 0x00400000 // termios.h:176:1:
	TTYDEF_CFLAG         = 19200      // ttydefaults.h:52:1:
	TTYDEF_IFLAG         = 11010      // ttydefaults.h:49:1:
	TTYDEF_LFLAG         = 1483       // ttydefaults.h:51:1:
	TTYDEF_OFLAG         = 3          // ttydefaults.h:50:1:
	TTYDEF_SPEED         = 9600       // ttydefaults.h:53:1:
	TTYDISC              = 0          // ttycom.h:142:1:
	VDISCARD             = 15         // termios.h:71:1:
	VDSUSP               = 11         // termios.h:65:1:
	VEOF                 = 0          // termios.h:47:1:
	VEOL                 = 1          // termios.h:48:1:
	VEOL2                = 2          // termios.h:50:1:
	VERASE               = 3          // termios.h:52:1:
	VINTR                = 8          // termios.h:61:1:
	VKILL                = 5          // termios.h:56:1:
	VLNEXT               = 14         // termios.h:70:1:
	VMIN                 = 16         // termios.h:73:1:
	VQUIT                = 9          // termios.h:62:1:
	VREPRINT             = 6          // termios.h:58:1:
	VSTART               = 12         // termios.h:67:1:
	VSTATUS              = 18         // termios.h:76:1:
	VSTOP                = 13         // termios.h:68:1:
	VSUSP                = 10         // termios.h:63:1:
	VTIME                = 17         // termios.h:74:1:
	VWERASE              = 4          // termios.h:54:1:
	XCASE                = 0x01000000 // termios.h:179:1:
	X_FILE_OFFSET_BITS   = 64         // <builtin>:25:1:
	X_LP64               = 1          // <predefined>:1:1:
	X_MACHINE_CDEFS_H_   = 0          // cdefs.h:9:1:
	X_MACHINE__TYPES_H_  = 0          // _types.h:36:1:
	X_MAX_PAGE_SHIFT     = 12         // _types.h:52:1:
	X_PID_T_DEFINED_     = 0          // termios.h:254:1:
	X_POSIX_VDISABLE     = 255        // termios.h:81:1:
	X_RET_PROTECTOR      = 1          // <predefined>:2:1:
	X_STACKALIGNBYTES    = 15         // _types.h:49:1:
	X_SYS_CDEFS_H_       = 0          // cdefs.h:39:1:
	X_SYS_IOCCOM_H_      = 0          // ioccom.h:36:1:
	X_SYS_TERMIOS_H_     = 0          // termios.h:36:1:
	X_SYS_TTYCOM_H_      = 0          // ttycom.h:41:1:
	X_SYS_TTYDEFAULTS_H_ = 0          // ttydefaults.h:44:1:
	X_SYS__TYPES_H_      = 0          // _types.h:35:1:
	Unix                 = 1          // <predefined>:340:1:
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

//	$OpenBSD: termios.h,v 1.13 2016/09/20 21:10:22 fcambus Exp $
//	$NetBSD: termios.h,v 1.14 1996/04/09 20:55:41 cgd Exp $

// Copyright (c) 1988, 1989, 1993, 1994
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
//	@(#)termios.h	8.3 (Berkeley) 3/28/94

//	$OpenBSD: cdefs.h,v 1.43 2018/10/29 17:10:40 guenther Exp $
//	$NetBSD: cdefs.h,v 1.16 1996/04/03 20:46:39 christos Exp $

// Copyright (c) 1991, 1993
//	The Regents of the University of California.  All rights reserved.
//
// This code is derived from software contributed to Berkeley by
// Berkeley Software Design, Inc.
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
//	@(#)cdefs.h	8.7 (Berkeley) 1/21/94

//	$OpenBSD: cdefs.h,v 1.3 2013/03/28 17:30:45 martynas Exp $

// Written by J.T. Conklin <jtc@wimsey.com> 01/17/95.
// Public domain.

// Macro to test if we're using a specific version of gcc or later.

// The __CONCAT macro is used to concatenate parts of symbol names, e.g.
// with "#define OLD(foo) __CONCAT(old,foo)", OLD(foo) produces oldfoo.
// The __CONCAT macro is a bit tricky -- make sure you don't put spaces
// in between its arguments.  Do not use __CONCAT on double-quoted strings,
// such as those from the __STRING macro: to concatenate strings just put
// them next to each other.

// GCC1 and some versions of GCC2 declare dead (non-returning) and
// pure (no side effects) functions using "volatile" and "const";
// unfortunately, these then cause warnings under "-ansi -pedantic".
// GCC >= 2.5 uses the __attribute__((attrs)) style.  All of these
// work for GNU C++ (modulo a slight glitch in the C++ grammar in
// the distribution version of 2.5.5).

// __returns_twice makes the compiler not assume the function
// only returns once.  This affects registerisation of variables:
// even local variables need to be in memory across such a call.
// Example: setjmp()

// __only_inline makes the compiler only use this function definition
// for inlining; references that can't be inlined will be left as
// external references instead of generating a local copy.  The
// matching library should include a simple extern definition for
// the function to handle those references.  c.f. ctype.h

// GNU C version 2.96 adds explicit branch prediction so that
// the CPU back-end can hint the processor and also so that
// code blocks can be reordered such that the predicted path
// sees a more linear flow, thus improving cache behavior, etc.
//
// The following two macros provide us with a way to utilize this
// compiler feature.  Use __predict_true() if you expect the expression
// to evaluate to true, and __predict_false() if you expect the
// expression to evaluate to false.
//
// A few notes about usage:
//
//	* Generally, __predict_false() error condition checks (unless
//	  you have some _strong_ reason to do otherwise, in which case
//	  document it), and/or __predict_true() `no-error' condition
//	  checks, assuming you want to optimize for the no-error case.
//
//	* Other than that, if you don't know the likelihood of a test
//	  succeeding from empirical or other `hard' evidence, don't
//	  make predictions.
//
//	* These are meant to be used in places that are run `a lot'.
//	  It is wasteful to make predictions in code that is run
//	  seldomly (e.g. at subsystem initialization time) as the
//	  basic block reordering that this affects can often generate
//	  larger code.

// Delete pseudo-keywords wherever they are not available or needed.

// The __packed macro indicates that a variable or structure members
// should have the smallest possible alignment, despite any host CPU
// alignment requirements.
//
// The __aligned(x) macro specifies the minimum alignment of a
// variable or structure.
//
// These macros together are useful for describing the layout and
// alignment of messages exchanged with hardware or other systems.

// "The nice thing about standards is that there are so many to choose from."
// There are a number of "feature test macros" specified by (different)
// standards that determine which interfaces and types the header files
// should expose.
//
// Because of inconsistencies in these macros, we define our own
// set in the private name space that end in _VISIBLE.  These are
// always defined and so headers can test their values easily.
// Things can get tricky when multiple feature macros are defined.
// We try to take the union of all the features requested.
//
// The following macros are guaranteed to have a value after cdefs.h
// has been included:
//	__POSIX_VISIBLE
//	__XPG_VISIBLE
//	__ISO_C_VISIBLE
//	__BSD_VISIBLE

// X/Open Portability Guides and Single Unix Specifications.
// _XOPEN_SOURCE				XPG3
// _XOPEN_SOURCE && _XOPEN_VERSION = 4		XPG4
// _XOPEN_SOURCE && _XOPEN_SOURCE_EXTENDED = 1	XPG4v2
// _XOPEN_SOURCE == 500				XPG5
// _XOPEN_SOURCE == 520				XPG5v2
// _XOPEN_SOURCE == 600				POSIX 1003.1-2001 with XSI
// _XOPEN_SOURCE == 700				POSIX 1003.1-2008 with XSI
//
// The XPG spec implies a specific value for _POSIX_C_SOURCE.

// POSIX macros, these checks must follow the XOPEN ones above.
//
// _POSIX_SOURCE == 1		1003.1-1988 (superseded by _POSIX_C_SOURCE)
// _POSIX_C_SOURCE == 1		1003.1-1990
// _POSIX_C_SOURCE == 2		1003.2-1992
// _POSIX_C_SOURCE == 199309L	1003.1b-1993
// _POSIX_C_SOURCE == 199506L   1003.1c-1995, 1003.1i-1995,
//				and the omnibus ISO/IEC 9945-1:1996
// _POSIX_C_SOURCE == 200112L   1003.1-2001
// _POSIX_C_SOURCE == 200809L   1003.1-2008
//
// The POSIX spec implies a specific value for __ISO_C_VISIBLE, though
// this may be overridden by the _ISOC99_SOURCE macro later.

// _ANSI_SOURCE means to expose ANSI C89 interfaces only.
// If the user defines it in addition to one of the POSIX or XOPEN
// macros, assume the POSIX/XOPEN macro(s) should take precedence.

// _ISOC99_SOURCE, _ISOC11_SOURCE, __STDC_VERSION__, and __cplusplus
// override any of the other macros since they are non-exclusive.

// Finally deal with BSD-specific interfaces that are not covered
// by any standards.  We expose these when none of the POSIX or XPG
// macros is defined or if the user explicitly asks for them.

// Default values.

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

type Tcflag_t = uint32 /* termios.h:185:22 */
type Cc_t = uint8      /* termios.h:186:23 */
type Speed_t = uint32  /* termios.h:187:22 */

type Termios = struct {
	Fc_iflag  Tcflag_t
	Fc_oflag  Tcflag_t
	Fc_cflag  Tcflag_t
	Fc_lflag  Tcflag_t
	Fc_cc     [20]Cc_t
	Fc_ispeed int32
	Fc_ospeed int32
} /* termios.h:189:1 */

// Commands passed to tcsetattr() for setting the termios structure.

// Standard speeds

//	$OpenBSD: _types.h,v 1.9 2014/08/22 23:05:15 krw Exp $

// -
// Copyright (c) 1990, 1993
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
//	@(#)types.h	8.3 (Berkeley) 1/5/94

//	$OpenBSD: _types.h,v 1.17 2018/03/05 01:15:25 deraadt Exp $

// -
// Copyright (c) 1990, 1993
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
//	@(#)types.h	8.3 (Berkeley) 1/5/94
//	@(#)ansi.h	8.2 (Berkeley) 1/4/94

// _ALIGN(p) rounds p (pointer or byte index) up to a correctly-aligned
// value for all data types (int, long, ...).   The result is an
// unsigned long and must be cast to any desired pointer type.
//
// _ALIGNED_POINTER is a boolean macro that checks whether an address
// is valid to fetch data elements of type t from on this architecture.
// This does not reflect the optimal alignment, just the possibility
// (within reasonable limits).

// 7.18.1.1 Exact-width integer types
type X__int8_t = int8     /* _types.h:61:22 */
type X__uint8_t = uint8   /* _types.h:62:24 */
type X__int16_t = int16   /* _types.h:63:17 */
type X__uint16_t = uint16 /* _types.h:64:25 */
type X__int32_t = int32   /* _types.h:65:15 */
type X__uint32_t = uint32 /* _types.h:66:23 */
type X__int64_t = int64   /* _types.h:67:20 */
type X__uint64_t = uint64 /* _types.h:68:28 */

// 7.18.1.2 Minimum-width integer types
type X__int_least8_t = X__int8_t     /* _types.h:71:19 */
type X__uint_least8_t = X__uint8_t   /* _types.h:72:20 */
type X__int_least16_t = X__int16_t   /* _types.h:73:20 */
type X__uint_least16_t = X__uint16_t /* _types.h:74:21 */
type X__int_least32_t = X__int32_t   /* _types.h:75:20 */
type X__uint_least32_t = X__uint32_t /* _types.h:76:21 */
type X__int_least64_t = X__int64_t   /* _types.h:77:20 */
type X__uint_least64_t = X__uint64_t /* _types.h:78:21 */

// 7.18.1.3 Fastest minimum-width integer types
type X__int_fast8_t = X__int32_t    /* _types.h:81:20 */
type X__uint_fast8_t = X__uint32_t  /* _types.h:82:21 */
type X__int_fast16_t = X__int32_t   /* _types.h:83:20 */
type X__uint_fast16_t = X__uint32_t /* _types.h:84:21 */
type X__int_fast32_t = X__int32_t   /* _types.h:85:20 */
type X__uint_fast32_t = X__uint32_t /* _types.h:86:21 */
type X__int_fast64_t = X__int64_t   /* _types.h:87:20 */
type X__uint_fast64_t = X__uint64_t /* _types.h:88:21 */

// 7.18.1.4 Integer types capable of holding object pointers
type X__intptr_t = int64   /* _types.h:103:16 */
type X__uintptr_t = uint64 /* _types.h:104:24 */

// 7.18.1.5 Greatest-width integer types
type X__intmax_t = X__int64_t   /* _types.h:107:20 */
type X__uintmax_t = X__uint64_t /* _types.h:108:21 */

// Register size
type X__register_t = int64 /* _types.h:111:16 */

// VM system types
type X__vaddr_t = uint64 /* _types.h:114:24 */
type X__paddr_t = uint64 /* _types.h:115:24 */
type X__vsize_t = uint64 /* _types.h:116:24 */
type X__psize_t = uint64 /* _types.h:117:24 */

// Standard system types
type X__double_t = float64           /* _types.h:120:18 */
type X__float_t = float32            /* _types.h:121:17 */
type X__ptrdiff_t = int64            /* _types.h:122:16 */
type X__size_t = uint64              /* _types.h:123:24 */
type X__ssize_t = int64              /* _types.h:124:16 */
type X__va_list = X__builtin_va_list /* _types.h:126:27 */

// Wide character support types
type X__wchar_t = int32     /* _types.h:133:15 */
type X__wint_t = int32      /* _types.h:135:15 */
type X__rune_t = int32      /* _types.h:136:15 */
type X__wctrans_t = uintptr /* _types.h:137:14 */
type X__wctype_t = uintptr  /* _types.h:138:14 */

type X__blkcnt_t = X__int64_t    /* _types.h:39:19 */ // blocks allocated for file
type X__blksize_t = X__int32_t   /* _types.h:40:19 */ // optimal blocksize for I/O
type X__clock_t = X__int64_t     /* _types.h:41:19 */ // ticks in CLOCKS_PER_SEC
type X__clockid_t = X__int32_t   /* _types.h:42:19 */ // CLOCK_* identifiers
type X__cpuid_t = uint64         /* _types.h:43:23 */ // CPU id
type X__dev_t = X__int32_t       /* _types.h:44:19 */ // device number
type X__fixpt_t = X__uint32_t    /* _types.h:45:20 */ // fixed point number
type X__fsblkcnt_t = X__uint64_t /* _types.h:46:20 */ // file system block count
type X__fsfilcnt_t = X__uint64_t /* _types.h:47:20 */ // file system file count
type X__gid_t = X__uint32_t      /* _types.h:48:20 */ // group id
type X__id_t = X__uint32_t       /* _types.h:49:20 */ // may contain pid, uid or gid
type X__in_addr_t = X__uint32_t  /* _types.h:50:20 */ // base type for internet address
type X__in_port_t = X__uint16_t  /* _types.h:51:20 */ // IP port type
type X__ino_t = X__uint64_t      /* _types.h:52:20 */ // inode number
type X__key_t = int64            /* _types.h:53:15 */ // IPC key (for Sys V IPC)
type X__mode_t = X__uint32_t     /* _types.h:54:20 */ // permissions
type X__nlink_t = X__uint32_t    /* _types.h:55:20 */ // link count
type X__off_t = X__int64_t       /* _types.h:56:19 */ // file offset or size
type X__pid_t = X__int32_t       /* _types.h:57:19 */ // process id
type X__rlim_t = X__uint64_t     /* _types.h:58:20 */ // resource limit
type X__sa_family_t = X__uint8_t /* _types.h:59:19 */ // sockaddr address family type
type X__segsz_t = X__int32_t     /* _types.h:60:19 */ // segment size
type X__socklen_t = X__uint32_t  /* _types.h:61:20 */ // length type for network syscalls
type X__suseconds_t = int64      /* _types.h:62:15 */ // microseconds (signed)
type X__swblk_t = X__int32_t     /* _types.h:63:19 */ // swap offset
type X__time_t = X__int64_t      /* _types.h:64:19 */ // epoch time
type X__timer_t = X__int32_t     /* _types.h:65:19 */ // POSIX timer identifiers
type X__uid_t = X__uint32_t      /* _types.h:66:20 */ // user id
type X__useconds_t = X__uint32_t /* _types.h:67:20 */ // microseconds

// mbstate_t is an opaque object to keep conversion state, during multibyte
// stream conversions. The content must not be referenced by user programs.
type X__mbstate_t = struct {
	F__ccgo_pad1 [0]uint64
	F__mbstate8  [128]int8
} /* _types.h:76:3 */

type Pid_t = X__pid_t /* termios.h:255:18 */

// Include tty ioctl's that aren't just for backwards compatibility
// with the old tty driver.  These ioctl definitions were previously
// in <sys/ioctl.h>.
//	$OpenBSD: ttycom.h,v 1.17 2018/06/16 13:55:03 deraadt Exp $
//	$NetBSD: ttycom.h,v 1.4 1996/05/19 17:17:53 jonathan Exp $

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

//	$OpenBSD: ioccom.h,v 1.5 2013/03/22 21:22:05 deraadt Exp $
//	$NetBSD: ioccom.h,v 1.4 1994/10/30 21:49:56 cgd Exp $

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
//	@(#)ioccom.h	8.2 (Berkeley) 3/28/94

// Ioctl's have the command encoded in the lower word, and the size of
// any in or out parameters in the upper word.  The high 3 bits of the
// upper word are used to encode the in/out status of the parameter.

// no parameters
// copy parameters out
// copy parameters in
// copy parameters in and out
// mask for IN/OUT/VOID

// this should be _IORW, but stdio got there first

// Tty ioctl's.

// Window/terminal size structure.  This information is stored by the kernel
// in order to provide a consistent interface, but is not used by the kernel.
type Winsize = struct {
	Fws_row    uint16
	Fws_col    uint16
	Fws_xpixel uint16
	Fws_ypixel uint16
} /* ttycom.h:51:1 */

type Tstamps = struct {
	Fts_set int32
	Fts_clr int32
} /* ttycom.h:58:1 */

// 8-10 compat
// 15 unused
// 17-18 compat
// 127-124 compat
// 117-116 compat

// Backwards compatibility

// END OF PROTECTED INCLUDE.

//	$OpenBSD: ttydefaults.h,v 1.7 2019/03/12 11:01:25 nicm Exp $
//	$NetBSD: ttydefaults.h,v 1.8 1996/04/09 20:55:45 cgd Exp $

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
