package qtgui

// /usr/include/qt/QtGui/qsurface.h
// #include <qsurface.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 43
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QSurface struct {
	*qtrt.CObject
}
type QSurface_ITF interface {
	QSurface_PTR() *QSurface
}

func (ptr *QSurface) QSurface_PTR() *QSurface { return ptr }

func (this *QSurface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSurface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSurfaceFromPointer(cthis unsafe.Pointer) *QSurface {
	return &QSurface{&qtrt.CObject{cthis}}
}
func (*QSurface) NewFromPointer(cthis unsafe.Pointer) *QSurface {
	return NewQSurfaceFromPointer(cthis)
}

// /usr/include/qt/QtGui/qsurface.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSurface()
func DeleteQSurface(this *QSurface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSurfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qsurface.h:74
// index:0
// Public Visibility=Default Availability=Available
// [4] QSurface::SurfaceClass surfaceClass()
func (this *QSurface) SurfaceClass() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSurface12surfaceClassEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qsurface.h:76
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSurfaceFormat format()
func (this *QSurface) Format() *QSurfaceFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSurface6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSurfaceFormat)
	return rv2
}

// /usr/include/qt/QtGui/qsurface.h:79
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QSurface::SurfaceType surfaceType()
func (this *QSurface) SurfaceType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSurface11surfaceTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qsurface.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool supportsOpenGL()
func (this *QSurface) SupportsOpenGL() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSurface14supportsOpenGLEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qsurface.h:82
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSize size()
func (this *QSurface) Size() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSurface4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qsurface.h:85
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QSurface(enum QSurface::SurfaceClass)
func NewQSurface(type_ int) *QSurface {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSurfaceC2ENS_12SurfaceClassE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSurfaceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSurface)
	return gothis
}

type QSurface__SurfaceClass = int

const QSurface__Window QSurface__SurfaceClass = 0
const QSurface__Offscreen QSurface__SurfaceClass = 1

type QSurface__SurfaceType = int

const QSurface__RasterSurface QSurface__SurfaceType = 0
const QSurface__OpenGLSurface QSurface__SurfaceType = 1
const QSurface__RasterGLSurface QSurface__SurfaceType = 2
const QSurface__OpenVGSurface QSurface__SurfaceType = 3
const QSurface__VulkanSurface QSurface__SurfaceType = 4

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
