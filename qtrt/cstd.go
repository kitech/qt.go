package qtrt

/*
#include <string.h>
#include <stdlib.h>
*/
import "C"

func Malloc(sz int) Voidptr        { return C.calloc(1, C.size_t(sz)) }
func Calloc(n int, sz int) Voidptr { return C.calloc(C.size_t(n), C.size_t(sz)) }
func CFree(p Voidptr)              { C.free(p) }
func FreeMem(p Voidptr)            { C.free(p) }
func FreeMemI(p uint64)            { C.free(Voidptr(uintptr(p))) }

func Cmemcmp(p1, p2 Voidptr, sz int) int {
	return int(C.memcmp(p1, p2, C.size_t(sz)))
}
