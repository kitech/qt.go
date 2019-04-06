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
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QOffscreenSurface) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QOffscreenSurface10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qoffscreensurface.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QOffscreenSurface(QScreen *, QObject *)

/*
Creates an offscreen surface for the targetScreen with the given parent.

The underlying platform surface is not created until create() is called.

This function was introduced in  Qt 5.10.

See also setScreen() and create().
*/
func (*QOffscreenSurface) NewForInherit(screen QScreen_ITF /*777 QScreen **/, parent qtcore.QObject_ITF /*777 QObject **/) *QOffscreenSurface {
	return NewQOffscreenSurface(screen, parent)
}
func NewQOffscreenSurface(screen QScreen_ITF /*777 QScreen **/, parent qtcore.QObject_ITF /*777 QObject **/) *QOffscreenSurface {
	var convArg0 unsafe.Pointer
	if screen != nil && screen.QScreen_PTR() != nil {
		convArg0 = screen.QScreen_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QOffscreenSurfaceC2EP7QScreenP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQOffscreenSurfaceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QOffscreenSurface")
	return gothis
}

// /usr/include/qt/QtGui/qoffscreensurface.h:62
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QOffscreenSurface(QScreen *)

/*
Creates an offscreen surface for the targetScreen with the given parent.

The underlying platform surface is not created until create() is called.

This function was introduced in  Qt 5.10.

See also setScreen() and create().
*/
func (*QOffscreenSurface) NewForInherit1(screen QScreen_ITF /*777 QScreen **/) *QOffscreenSurface {
	return NewQOffscreenSurface1(screen)
}
func NewQOffscreenSurface1(screen QScreen_ITF /*777 QScreen **/) *QOffscreenSurface {
	var convArg0 unsafe.Pointer
	if screen != nil && screen.QScreen_PTR() != nil {
		convArg0 = screen.QScreen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QOffscreenSurfaceC2EP7QScreen", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQOffscreenSurfaceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QOffscreenSurface")
	return gothis
}

// /usr/include/qt/QtGui/qoffscreensurface.h:62
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QOffscreenSurface(QScreen *)

/*
Creates an offscreen surface for the targetScreen with the given parent.

The underlying platform surface is not created until create() is called.

This function was introduced in  Qt 5.10.

See also setScreen() and create().
*/
func (*QOffscreenSurface) NewForInherit1p() *QOffscreenSurface {
	return NewQOffscreenSurface1p()
}
func NewQOffscreenSurface1p() *QOffscreenSurface {
	// arg: 0, QScreen *=Pointer, QScreen=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN17QOffscreenSurfaceC2EP7QScreen", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQOffscreenSurfaceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QOffscreenSurface")
	return gothis
}

// /usr/include/qt/QtGui/qoffscreensurface.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QOffscreenSurface()

/*

 */
func DeleteQOffscreenSurface(this *QOffscreenSurface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QOffscreenSurfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QSurface::SurfaceType surfaceType() const

/*
Reimplemented from QSurface::surfaceType().

Returns the surface type of the offscreen surface.

The surface type of an offscreen surface is always QSurface::OpenGLSurface.
*/
func (this *QOffscreenSurface) SurfaceType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QOffscreenSurface11surfaceTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void create()

/*
Allocates the platform resources associated with the offscreen surface.

It is at this point that the surface format set using setFormat() gets resolved into an actual native surface.

Call destroy() to free the platform resources if necessary.

Note: Some platforms require this function to be called on the main (GUI) thread.

See also destroy().
*/
func (this *QOffscreenSurface) Create() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QOffscreenSurface6createEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void destroy()

/*
Releases the native platform resources associated with this offscreen surface.

See also create().
*/
func (this *QOffscreenSurface) Destroy() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QOffscreenSurface7destroyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if this offscreen surface is valid; otherwise returns false.

The offscreen surface is valid if the platform resources have been successfuly allocated.

See also create().
*/
func (this *QOffscreenSurface) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QOffscreenSurface7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qoffscreensurface.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const QSurfaceFormat &)

/*
Sets the offscreen surface format.

The surface format will be resolved in the create() function. Calling this function after create() will not re-resolve the surface format of the native surface.

See also format(), create(), and destroy().
*/
func (this *QOffscreenSurface) SetFormat(format QSurfaceFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QSurfaceFormat_PTR() != nil {
		convArg0 = format.QSurfaceFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QOffscreenSurface9setFormatERK14QSurfaceFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSurfaceFormat format() const

/*
Reimplemented from QSurface::format().

Returns the actual format of this offscreen surface.

After the offscreen surface has been created, this function will return the actual surface format of the surface. It might differ from the requested format if the requested format could not be fulfilled by the platform.

See also setFormat(), create(), and requestedFormat().
*/
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
// [8] QSurfaceFormat requestedFormat() const

/*
Returns the requested surfaceformat of this offscreen surface.

If the requested format was not supported by the platform implementation, the requestedFormat will differ from the actual offscreen surface format.

This is the value set with setFormat().

See also setFormat() and format().
*/
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
// [8] QSize size() const

/*
Reimplemented from QSurface::size().

Returns the size of the offscreen surface.
*/
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
// [8] QScreen * screen() const

/*
Returns the screen to which the offscreen surface is connected.

See also setScreen().
*/
func (this *QOffscreenSurface) Screen() *QScreen /*777 QScreen **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QOffscreenSurface6screenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQScreenFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qoffscreensurface.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScreen(QScreen *)

/*
Sets the screen to which the offscreen surface is connected.

If the offscreen surface has been created, it will be recreated on the newScreen.

See also screen().
*/
func (this *QOffscreenSurface) SetScreen(screen QScreen_ITF /*777 QScreen **/) {
	var convArg0 unsafe.Pointer
	if screen != nil && screen.QScreen_PTR() != nil {
		convArg0 = screen.QScreen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QOffscreenSurface9setScreenEP7QScreen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:83
// index:0
// Public Visibility=Default Availability=Available
// [8] void * nativeHandle() const

/*
Returns an optional native handle to which the offscreen surface is connected.

This function was introduced in  Qt 5.9.

See also setNativeHandle().
*/
func (this *QOffscreenSurface) NativeHandle() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QOffscreenSurface12nativeHandleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qoffscreensurface.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNativeHandle(void *)

/*
Sets the native handle to which the offscreen surface is connected to handle.

The native handle will be resolved in the create() function. Calling this function after create() will not re-create a native surface.

Note: The interpretation of the native handle is platform specific. Only some platforms will support adopting native handles of offscreen surfaces and platforms that do not implement this support will ignore the handle.

This function was introduced in  Qt 5.9.

See also nativeHandle().
*/
func (this *QOffscreenSurface) SetNativeHandle(handle unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QOffscreenSurface15setNativeHandleEPv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), handle)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void screenChanged(QScreen *)

/*
This signal is emitted when an offscreen surface's screen changes, either by being set explicitly with setScreen(), or automatically when the window's screen is removed.
*/
func (this *QOffscreenSurface) ScreenChanged(screen QScreen_ITF /*777 QScreen **/) {
	var convArg0 unsafe.Pointer
	if screen != nil && screen.QScreen_PTR() != nil {
		convArg0 = screen.QScreen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QOffscreenSurface13screenChangedEP7QScreen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_10899() {
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
