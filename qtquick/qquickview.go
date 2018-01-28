package qtquick

// /usr/include/qt/QtQuick/qquickview.h
// #include <qquickview.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 67
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
type QQuickView struct {
	*QQuickWindow
}

func (this *QQuickView) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QQuickWindow.GetCthis()
	}
}
func (this *QQuickView) SetCthis(cthis unsafe.Pointer) {
	this.QQuickWindow = NewQQuickWindowFromPointer(cthis)
}
func NewQQuickViewFromPointer(cthis unsafe.Pointer) *QQuickView {
	bcthis0 := NewQQuickWindowFromPointer(cthis)
	return &QQuickView{bcthis0}
}
func (*QQuickView) NewFromPointer(cthis unsafe.Pointer) *QQuickView {
	return NewQQuickViewFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qquickview.h:58
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QQuickView) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickView10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickview.h:63
// index:0
// Public
// void QQuickView(QWindow *)
func NewQQuickView(parent *qtgui.QWindow /*777 QWindow **/) *QQuickView {
	cthis := qtrt.Calloc(1, 256) // 40
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickViewC2EP7QWindow", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuickViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qquickview.h:64
// index:1
// Public
// void QQuickView(QQmlEngine *, QWindow *)
func NewQQuickView_1(engine *qtqml.QQmlEngine /*777 QQmlEngine **/, parent *qtgui.QWindow /*777 QWindow **/) *QQuickView {
	cthis := qtrt.Calloc(1, 256) // 40
	var convArg0 = engine.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickViewC2EP10QQmlEngineP7QWindow", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuickViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qquickview.h:65
// index:2
// Public
// void QQuickView(const QUrl &, QWindow *)
func NewQQuickView_2(source *qtcore.QUrl, parent *qtgui.QWindow /*777 QWindow **/) *QQuickView {
	cthis := qtrt.Calloc(1, 256) // 40
	var convArg0 = source.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickViewC2ERK4QUrlP7QWindow", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuickViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qquickview.h:66
// index:0
// Public virtual
// void ~QQuickView()
func DeleteQQuickView(*QQuickView) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickViewD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:68
// index:0
// Public
// QUrl source()
func (this *QQuickView) Source() *qtcore.QUrl /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickView6sourceEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickview.h:70
// index:0
// Public
// QQmlEngine * engine()
func (this *QQuickView) Engine() *qtqml.QQmlEngine /*777 QQmlEngine **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickView6engineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtqml.NewQQmlEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickview.h:71
// index:0
// Public
// QQmlContext * rootContext()
func (this *QQuickView) RootContext() *qtqml.QQmlContext /*777 QQmlContext **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickView11rootContextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtqml.NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickview.h:73
// index:0
// Public
// QQuickItem * rootObject()
func (this *QQuickView) RootObject() *QQuickItem /*777 QQuickItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickView10rootObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickview.h:77
// index:0
// Public
// QQuickView::ResizeMode resizeMode()
func (this *QQuickView) ResizeMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickView10resizeModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickview.h:78
// index:0
// Public
// void setResizeMode(enum QQuickView::ResizeMode)
func (this *QQuickView) SetResizeMode(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickView13setResizeModeENS_10ResizeModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:82
// index:0
// Public
// QQuickView::Status status()
func (this *QQuickView) Status() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QQuickView6statusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickview.h:86
// index:0
// Public
// QSize sizeHint()
func (this *QQuickView) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc7("_ZNK10QQuickView8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickview.h:87
// index:0
// Public
// QSize initialSize()
func (this *QQuickView) InitialSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc7("_ZNK10QQuickView11initialSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickview.h:90
// index:0
// Public
// void setSource(const QUrl &)
func (this *QQuickView) SetSource(arg0 *qtcore.QUrl) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickView9setSourceERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:91
// index:0
// Public
// void setContent(const QUrl &, QQmlComponent *, QObject *)
func (this *QQuickView) SetContent(url *qtcore.QUrl, component *qtqml.QQmlComponent /*777 QQmlComponent **/, item *qtcore.QObject /*777 QObject **/) {
	var convArg0 = url.GetCthis()
	var convArg1 = component.GetCthis()
	var convArg2 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickView10setContentERK4QUrlP13QQmlComponentP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:94
// index:0
// Public
// void statusChanged(QQuickView::Status)
func (this *QQuickView) StatusChanged(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickView13statusChangedENS_6StatusE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:100
// index:0
// Protected virtual
// void resizeEvent(QResizeEvent *)
func (this *QQuickView) ResizeEvent(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickView11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:101
// index:0
// Protected virtual
// void timerEvent(QTimerEvent *)
func (this *QQuickView) TimerEvent(arg0 *qtcore.QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickView10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:103
// index:0
// Protected virtual
// void keyPressEvent(QKeyEvent *)
func (this *QQuickView) KeyPressEvent(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickView13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:104
// index:0
// Protected virtual
// void keyReleaseEvent(QKeyEvent *)
func (this *QQuickView) KeyReleaseEvent(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickView15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:105
// index:0
// Protected virtual
// void mousePressEvent(QMouseEvent *)
func (this *QQuickView) MousePressEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickView15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:106
// index:0
// Protected virtual
// void mouseReleaseEvent(QMouseEvent *)
func (this *QQuickView) MouseReleaseEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickView17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:107
// index:0
// Protected virtual
// void mouseMoveEvent(QMouseEvent *)
func (this *QQuickView) MouseMoveEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QQuickView14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QQuickView__ResizeMode = int

const QQuickView__SizeViewToRootObject QQuickView__ResizeMode = 0
const QQuickView__SizeRootObjectToView QQuickView__ResizeMode = 1

type QQuickView__Status = int

const QQuickView__Null QQuickView__Status = 0
const QQuickView__Ready QQuickView__Status = 1
const QQuickView__Loading QQuickView__Status = 2
const QQuickView__Error QQuickView__Status = 3

//  body block end
