package qtrt

/*
#include <string.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func Malloc(sz int) unsafe.Pointer        { return C.calloc(1, C.size_t(sz)) }
func Calloc(n int, sz int) unsafe.Pointer { return C.calloc(C.size_t(n), C.size_t(sz)) }
func CFree(p unsafe.Pointer)              { C.free(p) }
func FreeMem(p unsafe.Pointer)            { C.free(p) }

func Cmemcmp(p1, p2 unsafe.Pointer, sz int) int {
	return int(C.memcmp(p1, p2, C.size_t(sz)))
}
