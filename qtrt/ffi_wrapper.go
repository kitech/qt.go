package qtrt

/*
#include <stdint.h>
*/
import "C"
import (
	dl "github.com/kitech/dl/dl2"
)

// Library is a dl-opened library holding the corresponding dl.Handle
type FFILibrary struct {
	name   string
	handle dl.Handle
}

// NewFFILibrary takes the library filename and returns a handle towards it.
func NewFFILibrary(libname string) (lib FFILibrary, err error) {
	//libname = get_lib_arch_name(libname)
	lib.name = libname
	lib.handle, err = dl.Open(libname, dl.Now)
	return
}

func (lib FFILibrary) Name() string { return lib.name }
func (lib FFILibrary) Close() error {
	return lib.handle.Close()
}

func (lib FFILibrary) Symbol(fctname string) (Voidptr, error) {
	//println("Fct(",fctname,")...")
	sym, err := lib.handle.Symbol(fctname)
	if err != nil {
		return nil, err
	}

	addr := Voidptr(sym)
	return addr, nil
}

///
