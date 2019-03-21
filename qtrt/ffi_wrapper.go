package qtrt

import (
	"unsafe"

	"github.com/kitech/dl"
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

func (lib FFILibrary) Symbol(fctname string) (unsafe.Pointer, error) {
	//println("Fct(",fctname,")...")
	sym, err := lib.handle.Symbol(fctname)
	if err != nil {
		return nil, err
	}

	addr := unsafe.Pointer(sym)
	return addr, nil
}
