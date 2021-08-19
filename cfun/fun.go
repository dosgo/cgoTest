package cfun

//#cgo CFLAGS: -I./cfun
//#cgo 386,windows LDFLAGS: -L${SRCDIR}/cfun -lcfun_windows_386
//#include "cfun.h"
import "C"

func Number_add_mod(one C.int, two C.int, three C.int) C.int {
	return C.number_add_mod(one, two, three)
}
