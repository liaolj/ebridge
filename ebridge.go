package ebridge

/*
#cgo CFLAGS: -I/opt/egcs/epics/extensions/include -I/opt/egcs/epics/base/include -I/opt/egcs/epics/base/include/os/Linux
#cgo LDFLAGS: -L/opt/egcs/epics/extensions/lib/linux-x86_64 -lezca
#include <stdio.h>
#include <tsDefs.h>
#include <cadef.h>
#include "ezca.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

func ezcaInit() {
	C.ezcaSetTimeout(0.5)
	C.ezcaSetRetryCount(3)
}

func LongGet(pv string) (int, error) {
	ezcaInit()
	result := new(int)
	ezcaReturn := C.ezcaGet(C.CString(pv), C.ezcaLong, 1, unsafe.Pointer(result))
	if ezcaReturn != C.EZCA_OK {
		return -1, errors.New("long PV获取失败")
	}
	return *result, nil
}

func Edemo() int {
	d := new(int)
	C.ezcaGet(C.CString("epicsHost:ai1"), C.ezcaLong, 1, unsafe.Pointer(d))
	fmt.Printf("d=[%d]\n", *d)
	return *d
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
