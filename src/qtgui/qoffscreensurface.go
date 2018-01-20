//  header block begin
// /usr/include/qt/QtGui/qoffscreensurface.h
// #include <qoffscreensurface.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 43
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtGui/qoffscreensurface.h:56
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QOffscreenSurface) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QOffscreenSurface10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:61
// index:0
// void QOffscreenSurface(class QScreen *, class QObject *)
func NewQOffscreenSurface(screen unsafe.Pointer, parent unsafe.Pointer) *QOffscreenSurface {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QOffscreenSurfaceC2EP7QScreenP7QObject", ffiqt.FFI_TYPE_VOID, cthis, screen, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQOffscreenSurfaceFromPointer(cthis)
	return gothis
}
func NewQOffscreenSurfaceFromPointer(cthis unsafe.Pointer) *QOffscreenSurface {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := NewQSurfaceFromPointer(cthis)
	return &QOffscreenSurface{bcthis0, bcthis1}
}

// /usr/include/qt/QtGui/qoffscreensurface.h:62
// index:1
// void QOffscreenSurface(class QScreen *)
func NewQOffscreenSurface_1(screen unsafe.Pointer) *QOffscreenSurface {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QOffscreenSurfaceC2EP7QScreen", ffiqt.FFI_TYPE_VOID, cthis, screen)
	gopp.ErrPrint(err, rv)
	gothis := NewQOffscreenSurfaceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qoffscreensurface.h:63
// index:0
// virtual
// void ~QOffscreenSurface()
func DeleteQOffscreenSurface(*QOffscreenSurface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QOffscreenSurfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:65
// index:0
// virtual
// QSurface::SurfaceType surfaceType()
func (this *QOffscreenSurface) SurfaceType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QOffscreenSurface11surfaceTypeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:67
// index:0
// void create()
func (this *QOffscreenSurface) Create() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QOffscreenSurface6createEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:68
// index:0
// void destroy()
func (this *QOffscreenSurface) Destroy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QOffscreenSurface7destroyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:70
// index:0
// bool isValid()
func (this *QOffscreenSurface) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QOffscreenSurface7isValidEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:72
// index:0
// void setFormat(const class QSurfaceFormat &)
func (this *QOffscreenSurface) SetFormat(format unsafe.Pointer) {
	// 0: (, format const QSurfaceFormat &), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QOffscreenSurface9setFormatERK14QSurfaceFormat", ffiqt.FFI_TYPE_VOID, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:73
// index:0
// virtual
// QSurfaceFormat format()
func (this *QOffscreenSurface) Format() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QOffscreenSurface6formatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:74
// index:0
// QSurfaceFormat requestedFormat()
func (this *QOffscreenSurface) RequestedFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QOffscreenSurface15requestedFormatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:76
// index:0
// virtual
// QSize size()
func (this *QOffscreenSurface) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QOffscreenSurface4sizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:78
// index:0
// QScreen * screen()
func (this *QOffscreenSurface) Screen() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QOffscreenSurface6screenEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:79
// index:0
// void setScreen(class QScreen *)
func (this *QOffscreenSurface) SetScreen(screen unsafe.Pointer) {
	// 0: (, screen QScreen *), (screen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QOffscreenSurface9setScreenEP7QScreen", ffiqt.FFI_TYPE_VOID, this.GetCthis(), screen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:81
// index:0
// QPlatformOffscreenSurface * handle()
func (this *QOffscreenSurface) Handle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QOffscreenSurface6handleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:83
// index:0
// void * nativeHandle()
func (this *QOffscreenSurface) NativeHandle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QOffscreenSurface12nativeHandleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:84
// index:0
// void setNativeHandle(void *)
func (this *QOffscreenSurface) SetNativeHandle(handle unsafe.Pointer) {
	// 0: (, handle void *), (handle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QOffscreenSurface15setNativeHandleEPv", ffiqt.FFI_TYPE_VOID, this.GetCthis(), handle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qoffscreensurface.h:87
// index:0
// void screenChanged(class QScreen *)
func (this *QOffscreenSurface) ScreenChanged(screen unsafe.Pointer) {
	// 0: (, screen QScreen *), (screen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QOffscreenSurface13screenChangedEP7QScreen", ffiqt.FFI_TYPE_VOID, this.GetCthis(), screen)
	gopp.ErrPrint(err, rv)
}

//  body block end
