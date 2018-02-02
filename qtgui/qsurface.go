package qtgui

// /usr/include/qt/QtGui/qsurface.h
// #include <qsurface.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 39
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

type QSurface struct {
	*qtrt.CObject
}

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
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSurfaceD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qsurface.h:74
// index:0
// Public Visibility=Default Availability=Available
// [4] QSurface::SurfaceClass surfaceClass()
func (this *QSurface) SurfaceClass() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSurface12surfaceClassEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qsurface.h:76
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSurfaceFormat format()
func (this *QSurface) Format() *QSurfaceFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSurface6formatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSurfaceFormat)
	return rv2
}

// /usr/include/qt/QtGui/qsurface.h:79
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QSurface::SurfaceType surfaceType()
func (this *QSurface) SurfaceType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSurface11surfaceTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qsurface.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool supportsOpenGL()
func (this *QSurface) SupportsOpenGL() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSurface14supportsOpenGLEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qsurface.h:82
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSize size()
func (this *QSurface) Size() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QSurface4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qsurface.h:85
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QSurface(enum QSurface::SurfaceClass)
func NewQSurface(type_ int) *QSurface {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QSurfaceC1ENS_12SurfaceClassE", ffiqt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
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
