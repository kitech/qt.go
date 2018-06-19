package qtquick

// /usr/include/qt/QtQuick/qquickview.h
// #include <qquickview.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 69
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

// void resizeEvent(QResizeEvent *)
func (this *QQuickView) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void timerEvent(QTimerEvent *)
func (this *QQuickView) InheritTimerEvent(f func(arg0 *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QQuickView) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(QKeyEvent *)
func (this *QQuickView) InheritKeyReleaseEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QQuickView) InheritMousePressEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QQuickView) InheritMouseReleaseEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QQuickView) InheritMouseMoveEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

/*

 */
type QQuickView struct {
	*QQuickWindow
}
type QQuickView_ITF interface {
	QQuickWindow_ITF
	QQuickView_PTR() *QQuickView
}

func (ptr *QQuickView) QQuickView_PTR() *QQuickView { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QQuickView) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickView10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickview.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickView(QWindow *)

/*
Constructs a QQuickView with the given parent. The default value of parent is 0.
*/
func NewQQuickView(parent qtgui.QWindow_ITF /*777 QWindow **/) *QQuickView {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWindow_PTR() != nil {
		convArg0 = parent.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickViewC2EP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickView")
	return gothis
}

// /usr/include/qt/QtQuick/qquickview.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickView(QWindow *)

/*
Constructs a QQuickView with the given parent. The default value of parent is 0.
*/
func NewQQuickView__() *QQuickView {
	// arg: 0, QWindow *=Pointer, QWindow=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickViewC2EP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickView")
	return gothis
}

// /usr/include/qt/QtQuick/qquickview.h:64
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QQuickView(QQmlEngine *, QWindow *)

/*
Constructs a QQuickView with the given parent. The default value of parent is 0.
*/
func NewQQuickView_1(engine qtqml.QQmlEngine_ITF /*777 QQmlEngine **/, parent qtgui.QWindow_ITF /*777 QWindow **/) *QQuickView {
	var convArg0 unsafe.Pointer
	if engine != nil && engine.QQmlEngine_PTR() != nil {
		convArg0 = engine.QQmlEngine_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWindow_PTR() != nil {
		convArg1 = parent.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickViewC2EP10QQmlEngineP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickView")
	return gothis
}

// /usr/include/qt/QtQuick/qquickview.h:65
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QQuickView(const QUrl &, QWindow *)

/*
Constructs a QQuickView with the given parent. The default value of parent is 0.
*/
func NewQQuickView_2(source qtcore.QUrl_ITF, parent qtgui.QWindow_ITF /*777 QWindow **/) *QQuickView {
	var convArg0 unsafe.Pointer
	if source != nil && source.QUrl_PTR() != nil {
		convArg0 = source.QUrl_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWindow_PTR() != nil {
		convArg1 = parent.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickViewC2ERK4QUrlP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickView")
	return gothis
}

// /usr/include/qt/QtQuick/qquickview.h:65
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QQuickView(const QUrl &, QWindow *)

/*
Constructs a QQuickView with the given parent. The default value of parent is 0.
*/
func NewQQuickView_2_(source qtcore.QUrl_ITF) *QQuickView {
	var convArg0 unsafe.Pointer
	if source != nil && source.QUrl_PTR() != nil {
		convArg0 = source.QUrl_PTR().GetCthis()
	}
	// arg: 1, QWindow *=Pointer, QWindow=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickViewC2ERK4QUrlP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickView")
	return gothis
}

// /usr/include/qt/QtQuick/qquickview.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQuickView()

/*

 */
func DeleteQQuickView(this *QQuickView) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickViewD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qquickview.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl source() const

/*
Returns the source URL, if set.

Note: Getter function for property source.

See also setSource().
*/
func (this *QQuickView) Source() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickView6sourceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtQuick/qquickview.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlEngine * engine() const

/*
Returns a pointer to the QQmlEngine used for instantiating QML Components.
*/
func (this *QQuickView) Engine() *qtqml.QQmlEngine /*777 QQmlEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickView6engineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtqml.NewQQmlEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickview.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlContext * rootContext() const

/*
This function returns the root of the context hierarchy. Each QML component is instantiated in a QQmlContext. QQmlContext's are essential for passing data to QML components. In QML, contexts are arranged hierarchically and this hierarchy is managed by the QQmlEngine.
*/
func (this *QQuickView) RootContext() *qtqml.QQmlContext /*777 QQmlContext **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickView11rootContextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtqml.NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickview.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickItem * rootObject() const

/*
Returns the view's root item.
*/
func (this *QQuickView) RootObject() *QQuickItem /*777 QQuickItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickView10rootObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickview.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] QQuickView::ResizeMode resizeMode() const

/*

 */
func (this *QQuickView) ResizeMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickView10resizeModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickview.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResizeMode(QQuickView::ResizeMode)

/*

 */
func (this *QQuickView) SetResizeMode(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickView13setResizeModeENS_10ResizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] QQuickView::Status status() const

/*

 */
func (this *QQuickView) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickView6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickview.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*

 */
func (this *QQuickView) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickView8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtQuick/qquickview.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize initialSize() const

/*
Returns the initial size of the root object.

If resizeMode is QQuickItem::SizeRootObjectToView the root object will be resized to the size of the view. initialSize contains the size of the root object before it was resized.
*/
func (this *QQuickView) InitialSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickView11initialSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtQuick/qquickview.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSource(const QUrl &)

/*
Sets the source to the url, loads the QML component and instantiates it.

Ensure that the URL provided is full and correct, in particular, use QUrl::fromLocalFile() when loading a file from the local filesystem.

Calling this method multiple times with the same url will result in the QML component being reinstantiated.

Note: Setter function for property source.

See also source().
*/
func (this *QQuickView) SetSource(arg0 qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QUrl_PTR() != nil {
		convArg0 = arg0.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickView9setSourceERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContent(const QUrl &, QQmlComponent *, QObject *)

/*

 */
func (this *QQuickView) SetContent(url qtcore.QUrl_ITF, component qtqml.QQmlComponent_ITF /*777 QQmlComponent **/, item qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if component != nil && component.QQmlComponent_PTR() != nil {
		convArg1 = component.QQmlComponent_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if item != nil && item.QObject_PTR() != nil {
		convArg2 = item.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickView10setContentERK4QUrlP13QQmlComponentP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void statusChanged(QQuickView::Status)

/*
This signal is emitted when the component's current status changes.

Note: Notifier signal for property status.
*/
func (this *QQuickView) StatusChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickView13statusChangedENS_6StatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:100
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*

 */
func (this *QQuickView) ResizeEvent(arg0 qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QResizeEvent_PTR() != nil {
		convArg0 = arg0.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickView11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:101
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)

/*

 */
func (this *QQuickView) TimerEvent(arg0 qtcore.QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTimerEvent_PTR() != nil {
		convArg0 = arg0.QTimerEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickView10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:103
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QWindow::keyPressEvent().
*/
func (this *QQuickView) KeyPressEvent(arg0 qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QKeyEvent_PTR() != nil {
		convArg0 = arg0.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickView13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:104
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)

/*
Reimplemented from QWindow::keyReleaseEvent().
*/
func (this *QQuickView) KeyReleaseEvent(arg0 qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QKeyEvent_PTR() != nil {
		convArg0 = arg0.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickView15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:105
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
Reimplemented from QWindow::mousePressEvent().
*/
func (this *QQuickView) MousePressEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickView15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:106
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
Reimplemented from QWindow::mouseReleaseEvent().
*/
func (this *QQuickView) MouseReleaseEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickView17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickview.h:107
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
Reimplemented from QWindow::mouseMoveEvent().
*/
func (this *QQuickView) MouseMoveEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickView14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
This enum specifies how to resize the view.


*/
type QQuickView__ResizeMode = int

// The view resizes with the root item in the QML.
const QQuickView__SizeViewToRootObject QQuickView__ResizeMode = 0

// The view will automatically resize the root item to the size of the view.
const QQuickView__SizeRootObjectToView QQuickView__ResizeMode = 1

/*
Specifies the loading status of the QQuickView.


*/
type QQuickView__Status = int

// This QQuickView has no source set.
const QQuickView__Null QQuickView__Status = 0

// This QQuickView has loaded and created the QML component.
const QQuickView__Ready QQuickView__Status = 1

// This QQuickView is loading network data.
const QQuickView__Loading QQuickView__Status = 2

// One or more errors has occurred. Call errors() to retrieve a list of errors.
const QQuickView__Error QQuickView__Status = 3

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
