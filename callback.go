package hackrf

// #cgo windows CFLAGS: -I.
// #cgo windows LDFLAGS: -lhackrf -L.
// #cgo linux LDFLAGS: -lhackrf -lusb-1.0 -L.
// #cgo linux CFLAGS: -I.
// #include "hackrf.h"
import "C"
import "unsafe"

var localCallback CallbackFunc

//export go_callback
func go_callback(transfer C.hackrf_transfer) C.int {
	// I know this looks incredibly unsafe. We slice the array on
	// transfer.valid_length and set it's capacity to the same so the user
	// can't walk off the end of the array if they try to do anything like
	// appending to the slice.
	buf := (*[MaxArrayLen]int8)(unsafe.Pointer(transfer.buffer))
	return C.int(localCallback(buf[:transfer.valid_length:transfer.valid_length]))
}
