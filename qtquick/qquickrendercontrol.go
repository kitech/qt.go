package qtquick

// /usr/include/qt/QtQuick/qquickrendercontrol.h
// #include <qquickrendercontrol.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 33
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"
import "qt.go/qtnetwork"
import "qt.go/qtqml"

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
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  ext block end

//  body block begin
type QQuickRenderControl struct {
	*qtcore.QObject
}

func (this *QQuickRenderControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QQuickRenderControl) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQQuickRenderControlFromPointer(cthis unsafe.Pointer) *QQuickRenderControl {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QQuickRenderControl{bcthis0}
}
func (*QQuickRenderControl) NewFromPointer(cthis unsafe.Pointer) *QQuickRenderControl {
	return NewQQuickRenderControlFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QQuickRenderControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QQuickRenderControl10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickRenderControl(QObject *)
func NewQQuickRenderControl(parent *qtcore.QObject /*777 QObject **/) *QQuickRenderControl {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QQuickRenderControlC2EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuickRenderControlFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQuickRenderControl()
func DeleteQQuickRenderControl(*QQuickRenderControl) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QQuickRenderControlD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void prepareThread(QThread *)
func (this *QQuickRenderControl) PrepareThread(targetThread *qtcore.QThread /*777 QThread **/) {
	var convArg0 = targetThread.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QQuickRenderControl13prepareThreadEP7QThread", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void invalidate()
func (this *QQuickRenderControl) Invalidate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QQuickRenderControl10invalidateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void polishItems()
func (this *QQuickRenderControl) PolishItems() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QQuickRenderControl11polishItemsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void render()
func (this *QQuickRenderControl) Render() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QQuickRenderControl6renderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:67
// index:0
// Public Visibility=Default Availability=Available
// [1] bool sync()
func (this *QQuickRenderControl) Sync() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QQuickRenderControl4syncEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:69
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage grab()
func (this *QQuickRenderControl) Grab() *qtgui.QImage /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QQuickRenderControl4grabEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:71
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWindow * renderWindowFor(QQuickWindow *, QPoint *)
func (this *QQuickRenderControl) RenderWindowFor(win *QQuickWindow /*777 QQuickWindow **/, offset *qtcore.QPoint /*777 QPoint **/) *qtgui.QWindow /*777 QWindow **/ {
	var convArg0 = win.GetCthis()
	var convArg1 = offset.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QQuickRenderControl15renderWindowForEP12QQuickWindowP6QPoint", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtgui.NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QQuickRenderControl_RenderWindowFor(win *QQuickWindow /*777 QQuickWindow **/, offset *qtcore.QPoint /*777 QPoint **/) *qtgui.QWindow /*777 QWindow **/ {
	var nilthis *QQuickRenderControl
	rv := nilthis.RenderWindowFor(win, offset)
	return rv
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:72
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [8] QWindow * renderWindow(QPoint *)
func (this *QQuickRenderControl) RenderWindow(offset *qtcore.QPoint /*777 QPoint **/) *qtgui.QWindow /*777 QWindow **/ {
	var convArg0 = offset.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QQuickRenderControl12renderWindowEP6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void renderRequested()
func (this *QQuickRenderControl) RenderRequested() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QQuickRenderControl15renderRequestedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sceneChanged()
func (this *QQuickRenderControl) SceneChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QQuickRenderControl12sceneChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
