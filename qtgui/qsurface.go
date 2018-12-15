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
// extern C begin: 44
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
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

// /usr/include/qt/QtGui/qsurface.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSurface()

/*

 */
func DeleteQSurface(this *QSurface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSurfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qsurface.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] QSurface::SurfaceClass surfaceClass() const

/*
Returns the surface class of this surface.
*/
func (this *QSurface) SurfaceClass() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSurface12surfaceClassEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qsurface.h:80
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSurfaceFormat format() const

/*
Returns the format of the surface.
*/
func (this *QSurface) Format() *QSurfaceFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSurface6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSurfaceFormat)
	return rv2
}

// /usr/include/qt/QtGui/qsurface.h:83
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QSurface::SurfaceType surfaceType() const

/*
Returns the type of the surface.
*/
func (this *QSurface) SurfaceType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSurface11surfaceTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qsurface.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool supportsOpenGL() const

/*
Returns true if the surface is OpenGL compatible and can be used in conjunction with QOpenGLContext; otherwise returns false.

This function was introduced in  Qt 5.3.
*/
func (this *QSurface) SupportsOpenGL() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSurface14supportsOpenGLEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qsurface.h:86
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSize size() const

/*
Returns the size of the surface in pixels.
*/
func (this *QSurface) Size() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QSurface4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qsurface.h:89
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QSurface(QSurface::SurfaceClass)

/*
Creates a surface with the given type.
*/
func (*QSurface) NewForInherit(type_ int) *QSurface {
	return NewQSurface(type_)
}
func NewQSurface(type_ int) *QSurface {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QSurfaceC2ENS_12SurfaceClassE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSurfaceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSurface)
	return gothis
}

/*
The SurfaceClass enum describes the actual subclass of the surface.


*/
type QSurface__SurfaceClass = int

// The surface is an instance of QWindow.
const QSurface__Window QSurface__SurfaceClass = 0

// The surface is an instance of QOffscreenSurface.
const QSurface__Offscreen QSurface__SurfaceClass = 1

func (this *QSurface) SurfaceClassItemName(val int) string {
	switch val {
	case QSurface__Window: // 0
		return "Window"
	case QSurface__Offscreen: // 1
		return "Offscreen"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSurface_SurfaceClassItemName(val int) string {
	var nilthis *QSurface
	return nilthis.SurfaceClassItemName(val)
}

/*
The SurfaceType enum describes what type of surface this is.


*/
type QSurface__SurfaceType = int

// The surface is is composed of pixels and can be rendered to using a software rasterizer like Qt's raster paint engine.
const QSurface__RasterSurface QSurface__SurfaceType = 0

// The surface is an OpenGL compatible surface and can be used in conjunction with QOpenGLContext.
const QSurface__OpenGLSurface QSurface__SurfaceType = 1

// The surface can be rendered to using a software rasterizer, and also supports OpenGL. This surface type is intended for internal Qt use, and requires the use of private API.
const QSurface__RasterGLSurface QSurface__SurfaceType = 2

// The surface is an OpenVG compatible surface and can be used in conjunction with OpenVG contexts.
const QSurface__OpenVGSurface QSurface__SurfaceType = 3

// The surface is a Vulkan compatible surface and can be used in conjunction with the Vulkan graphics API.
const QSurface__VulkanSurface QSurface__SurfaceType = 4

// The surface is a Metal compatible surface and can be used in conjunction with Apple's Metal graphics API. This surface type is supported on macOS only.
const QSurface__MetalSurface QSurface__SurfaceType = 5

func (this *QSurface) SurfaceTypeItemName(val int) string {
	switch val {
	case QSurface__RasterSurface: // 0
		return "RasterSurface"
	case QSurface__OpenGLSurface: // 1
		return "OpenGLSurface"
	case QSurface__RasterGLSurface: // 2
		return "RasterGLSurface"
	case QSurface__OpenVGSurface: // 3
		return "OpenVGSurface"
	case QSurface__VulkanSurface: // 4
		return "VulkanSurface"
	case QSurface__MetalSurface: // 5
		return "MetalSurface"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSurface_SurfaceTypeItemName(val int) string {
	var nilthis *QSurface
	return nilthis.SurfaceTypeItemName(val)
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
