package qtgui

// /usr/include/qt/QtGui/qoffscreensurface.h
// #include <qoffscreensurface.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 42
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin
type QOffscreenSurface struct {
	*qtcore.QObject
	*QSurface
}

func (this *QOffscreenSurface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QOffscreenSurface) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
	this.QSurface = NewQSurfaceFromPointer(cthis)
}
func NewQOffscreenSurfaceFromPointer(cthis unsafe.Pointer) *QOffscreenSurface {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := NewQSurfaceFromPointer(cthis)
	return &QOffscreenSurface{bcthis0, bcthis1}
}
func (*QOffscreenSurface) NewFromPointer(cthis unsafe.Pointer) *QOffscreenSurface {
	return NewQOffscreenSurfaceFromPointer(cthis)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:56
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QOffscreenSurface) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QOffscreenSurface10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qoffscreensurface.h:61
// index:0
// Public
// void QOffscreenSurface(QScreen *, QObject *)
func NewQOffscreenSurface(screen *QScreen /*777 QScreen **/, parent *qtcore.QObject /*777 QObject **/) *QOffscreenSurface {
	cthis := qtrt.Calloc(1, 256) // 40
	var convArg0 = screen.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QOffscreenSurfaceC2EP7QScreenP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQOffscreenSurfaceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qoffscreensurface.h:62
// index:1
// Public
// void QOffscreenSurface(QScreen *)
func NewQOffscreenSurface_1(screen *QScreen /*777 QScreen **/) *QOffscreenSurface {
	cthis := qtrt.Calloc(1, 256) // 40
	var convArg0 = screen.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QOffscreenSurfaceC2EP7QScreen", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQOffscreenSurfaceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qoffscreensurface.h:63
// index:0
// Public virtual
// void ~QOffscreenSurface()
func DeleteQOffscreenSurface(*QOffscreenSurface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QOffscreenSurfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:65
// index:0
// Public virtual
// QSurface::SurfaceType surfaceType()
func (this *QOffscreenSurface) SurfaceType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QOffscreenSurface11surfaceTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:67
// index:0
// Public
// void create()
func (this *QOffscreenSurface) Create() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QOffscreenSurface6createEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:68
// index:0
// Public
// void destroy()
func (this *QOffscreenSurface) Destroy() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QOffscreenSurface7destroyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:70
// index:0
// Public
// bool isValid()
func (this *QOffscreenSurface) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QOffscreenSurface7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qoffscreensurface.h:72
// index:0
// Public
// void setFormat(const QSurfaceFormat &)
func (this *QOffscreenSurface) SetFormat(format *QSurfaceFormat) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QOffscreenSurface9setFormatERK14QSurfaceFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:73
// index:0
// Public virtual
// QSurfaceFormat format()
func (this *QOffscreenSurface) Format() *QSurfaceFormat /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QOffscreenSurface6formatEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qoffscreensurface.h:74
// index:0
// Public
// QSurfaceFormat requestedFormat()
func (this *QOffscreenSurface) RequestedFormat() *QSurfaceFormat /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QOffscreenSurface15requestedFormatEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qoffscreensurface.h:76
// index:0
// Public virtual
// QSize size()
func (this *QOffscreenSurface) Size() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc7("_ZNK17QOffscreenSurface4sizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qoffscreensurface.h:78
// index:0
// Public
// QScreen * screen()
func (this *QOffscreenSurface) Screen() *QScreen /*777 QScreen **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QOffscreenSurface6screenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQScreenFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qoffscreensurface.h:79
// index:0
// Public
// void setScreen(QScreen *)
func (this *QOffscreenSurface) SetScreen(screen *QScreen /*777 QScreen **/) {
	var convArg0 = screen.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QOffscreenSurface9setScreenEP7QScreen", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:83
// index:0
// Public
// void * nativeHandle()
func (this *QOffscreenSurface) NativeHandle() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QOffscreenSurface12nativeHandleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qoffscreensurface.h:84
// index:0
// Public
// void setNativeHandle(void *)
func (this *QOffscreenSurface) SetNativeHandle(handle unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QOffscreenSurface15setNativeHandleEPv", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), handle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:87
// index:0
// Public
// void screenChanged(QScreen *)
func (this *QOffscreenSurface) ScreenChanged(screen *QScreen /*777 QScreen **/) {
	var convArg0 = screen.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QOffscreenSurface13screenChangedEP7QScreen", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
