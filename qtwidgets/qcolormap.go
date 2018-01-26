package qtwidgets

// /usr/include/qt/QtWidgets/qcolormap.h
// #include <qcolormap.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QColormap struct {
	*qtrt.CObject
}

func (this *QColormap) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QColormap) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQColormapFromPointer(cthis unsafe.Pointer) *QColormap {
	return &QColormap{&qtrt.CObject{cthis}}
}
func (*QColormap) NewFromPointer(cthis unsafe.Pointer) *QColormap {
	return NewQColormapFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qcolormap.h:60
// index:0
// Public static
// void initialize()
func (this *QColormap) Initialize() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QColormap10initializeEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QColormap_Initialize() {
	var nilthis *QColormap
	nilthis.Initialize()
}

// /usr/include/qt/QtWidgets/qcolormap.h:61
// index:0
// Public static
// void cleanup()
func (this *QColormap) Cleanup() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QColormap7cleanupEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QColormap_Cleanup() {
	var nilthis *QColormap
	nilthis.Cleanup()
}

// /usr/include/qt/QtWidgets/qcolormap.h:63
// index:0
// Public static
// QColormap instance(int)
func (this *QColormap) Instance(screen int) *QColormap /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QColormap8instanceEi", ffiqt.FFI_TYPE_POINTER, screen)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQColormapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QColormap_Instance(screen int) *QColormap /*123*/ {
	var nilthis *QColormap
	rv := nilthis.Instance(screen)
	return rv
}

// /usr/include/qt/QtWidgets/qcolormap.h:66
// index:0
// Public
// void ~QColormap()
func DeleteQColormap(*QColormap) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QColormapD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolormap.h:70
// index:0
// Public
// QColormap::Mode mode()
func (this *QColormap) Mode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QColormap4modeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcolormap.h:72
// index:0
// Public
// int depth()
func (this *QColormap) Depth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QColormap5depthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qcolormap.h:73
// index:0
// Public
// int size()
func (this *QColormap) Size() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QColormap4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qcolormap.h:75
// index:0
// Public
// uint pixel(const class QColor &)
func (this *QColormap) Pixel(color *qtgui.QColor) uint {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QColormap5pixelERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtWidgets/qcolormap.h:76
// index:0
// Public
// const QColor colorAt(uint)
func (this *QColormap) ColorAt(pixel uint) *qtgui.QColor /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QColormap7colorAtEj", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), pixel)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

type QColormap__Mode = int

const QColormap__Direct QColormap__Mode = 0
const QColormap__Indexed QColormap__Mode = 1
const QColormap__Gray QColormap__Mode = 2

//  body block end
