package hackrf

/*
#cgo linux LDFLAGS: -lhackrf -lusb-1.0 -L.
#cgo linux CFLAGS: -I/usr/local/include/libhackrf
#include "hackrf.h"
*/
import "C"
import "unsafe"

var localCallback CallbackFunc

//export go_callback
func go_callback(transfer *C.hackrf_transfer) C.int {
	buf := unsafe.Slice(
		(*uint8)(unsafe.Pointer(transfer.buffer)),
		transfer.valid_length,
	)
	return C.int(localCallback(buf))
}
