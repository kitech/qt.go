//  header block begin

// +build !minimal

package qtquick

// /usr/include/qt/QtQuick/qquickwindow.h
// #include <qquickwindow.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

// /usr/include/qt/QtQuick/qquickwindow.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetOpenGLState()

/*
Call this function to reset the OpenGL context its default state.

The scene graph uses the OpenGL context and will both rely on and clobber its state. When mixing raw OpenGL commands with scene graph rendering, this function provides a convenient way of resetting the OpenGL context state back to its default values.

This function does not touch state in the fixed-function pipeline.

This function does not clear the color, depth and stencil buffers. Use QQuickWindow::setClearBeforeRendering to control clearing of the color buffer. The depth and stencil buffer might be clobbered by the scene graph renderer. Clear these manually on demand.

Note: This function only has an effect when using the default OpenGL scene graph adaptation.

This function was introduced in  Qt 5.2.

See also QQuickWindow::beforeRendering().
*/
func (this *QQuickWindow) ResetOpenGLState() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow16resetOpenGLStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:141
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * accessibleRoot() const

/*
Returns an accessibility interface for this window, or 0 if such an interface cannot be created.
*/
func (this *QQuickWindow) AccessibleRoot() *qtgui.QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow14accessibleRootEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:228
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
Reimplemented from QWindow::wheelEvent().
*/
func (this *QQuickWindow) WheelEvent(arg0 qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWheelEvent_PTR() != nil {
		convArg0 = arg0.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11588() {
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
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  keep block end
