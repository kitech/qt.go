package qtgui

// /usr/include/qt/QtGui/qoffscreensurface.h
// #include <qoffscreensurface.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 42
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QOffscreenSurface struct {
	*qtcore.QObject
	*QSurface
}
type QOffscreenSurface_ITF interface {
	qtcore.QObject_ITF
	QSurface_ITF
	QOffscreenSurface_PTR() *QOffscreenSurface
}

func (ptr *QOffscreenSurface) QOffscreenSurface_PTR() *QOffscreenSurface { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QOffscreenSurface) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QOffscreenSurface10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qoffscreensurface.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QOffscreenSurface(QScreen *, QObject *)
func NewQOffscreenSurface(screen QScreen_ITF /*777 QScreen **/, parent qtcore.QObject_ITF /*777 QObject **/) *QOffscreenSurface {
	var convArg0 = screen.QScreen_PTR().GetCthis()
	var convArg1 = parent.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QOffscreenSurfaceC2EP7QScreenP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQOffscreenSurfaceFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qoffscreensurface.h:62
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QOffscreenSurface(QScreen *)
func NewQOffscreenSurface_1(screen QScreen_ITF /*777 QScreen **/) *QOffscreenSurface {
	var convArg0 = screen.QScreen_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QOffscreenSurfaceC2EP7QScreen", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQOffscreenSurfaceFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qoffscreensurface.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QOffscreenSurface()
func DeleteQOffscreenSurface(this *QOffscreenSurface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QOffscreenSurfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QSurface::SurfaceType surfaceType()
func (this *QOffscreenSurface) SurfaceType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QOffscreenSurface11surfaceTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void create()
func (this *QOffscreenSurface) Create() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QOffscreenSurface6createEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void destroy()
func (this *QOffscreenSurface) Destroy() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QOffscreenSurface7destroyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QOffscreenSurface) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QOffscreenSurface7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qoffscreensurface.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const QSurfaceFormat &)
func (this *QOffscreenSurface) SetFormat(format QSurfaceFormat_ITF) {
	var convArg0 = format.QSurfaceFormat_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QOffscreenSurface9setFormatERK14QSurfaceFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSurfaceFormat format()
func (this *QOffscreenSurface) Format() *QSurfaceFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QOffscreenSurface6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSurfaceFormat)
	return rv2
}

// /usr/include/qt/QtGui/qoffscreensurface.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QSurfaceFormat requestedFormat()
func (this *QOffscreenSurface) RequestedFormat() *QSurfaceFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QOffscreenSurface15requestedFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSurfaceFormat)
	return rv2
}

// /usr/include/qt/QtGui/qoffscreensurface.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize size()
func (this *QOffscreenSurface) Size() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QOffscreenSurface4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qoffscreensurface.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QScreen * screen()
func (this *QOffscreenSurface) Screen() *QScreen /*777 QScreen **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QOffscreenSurface6screenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQScreenFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qoffscreensurface.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScreen(QScreen *)
func (this *QOffscreenSurface) SetScreen(screen QScreen_ITF /*777 QScreen **/) {
	var convArg0 = screen.QScreen_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QOffscreenSurface9setScreenEP7QScreen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:83
// index:0
// Public Visibility=Default Availability=Available
// [8] void * nativeHandle()
func (this *QOffscreenSurface) NativeHandle() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QOffscreenSurface12nativeHandleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qoffscreensurface.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNativeHandle(void *)
func (this *QOffscreenSurface) SetNativeHandle(handle unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QOffscreenSurface15setNativeHandleEPv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), handle)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void screenChanged(QScreen *)
func (this *QOffscreenSurface) ScreenChanged(screen QScreen_ITF /*777 QScreen **/) {
	var convArg0 = screen.QScreen_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QOffscreenSurface13screenChangedEP7QScreen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

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
		qtcore.KeepMe()
	}
}

//  keep block end
