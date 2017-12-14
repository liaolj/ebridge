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
	"fmt"
	"unsafe"
)

func main() {
	d := new(int)
	C.ezcaGet(C.CString("epicsHost:ai1"), C.ezcaLong, 1, unsafe.Pointer(d))
	fmt.Printf("d=[%d]\n", *d)
}
