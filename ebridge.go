package ebridge

/*
#cgo CFLAGS: -I./include -I/opt/egcs/epics/base/include -I/opt/egcs/epics/base/include/os/Linux
#cgo LDFLAGS: -L./lib/linux-x86_64 -lezca
#include <stdio.h>
#include <tsDefs.h>
#include <cadef.h>
#include "ezca.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	d := new(int)
	C.ezcaGet(C.CString("epicsHost:ai1"), C.ezcaLong, 1, unsafe.Pointer(d))
	fmt.Printf("d=[%d]\n", *d)
}
