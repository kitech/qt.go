package qtquickwidgets

// /usr/include/qt/QtQuickWidgets/qquickwidget.h
// #include <qquickwidget.h>
// #include <QtQuickWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"
import "github.com/kitech/qt.go/qtqml"
import "github.com/kitech/qt.go/qtquick"

//  ext block end

//  body block begin

// void resizeEvent(QResizeEvent *)
func (this *QQuickWidget) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void timerEvent(QTimerEvent *)
func (this *QQuickWidget) InheritTimerEvent(f func(arg0 *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QQuickWidget) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(QKeyEvent *)
func (this *QQuickWidget) InheritKeyReleaseEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QQuickWidget) InheritMousePressEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QQuickWidget) InheritMouseReleaseEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QQuickWidget) InheritMouseMoveEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseDoubleClickEvent(QMouseEvent *)
func (this *QQuickWidget) InheritMouseDoubleClickEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void showEvent(QShowEvent *)
func (this *QQuickWidget) InheritShowEvent(f func(arg0 *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void hideEvent(QHideEvent *)
func (this *QQuickWidget) InheritHideEvent(f func(arg0 *qtgui.QHideEvent /*777 QHideEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hideEvent", f)
}

// void focusInEvent(QFocusEvent *)
func (this *QQuickWidget) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QQuickWidget) InheritFocusOutEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void wheelEvent(QWheelEvent *)
func (this *QQuickWidget) InheritWheelEvent(f func(arg0 *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void dragEnterEvent(QDragEnterEvent *)
func (this *QQuickWidget) InheritDragEnterEvent(f func(arg0 *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragMoveEvent(QDragMoveEvent *)
func (this *QQuickWidget) InheritDragMoveEvent(f func(arg0 *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dragLeaveEvent(QDragLeaveEvent *)
func (this *QQuickWidget) InheritDragLeaveEvent(f func(arg0 *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dropEvent(QDropEvent *)
func (this *QQuickWidget) InheritDropEvent(f func(arg0 *qtgui.QDropEvent /*777 QDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// bool event(QEvent *)
func (this *QQuickWidget) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void paintEvent(QPaintEvent *)
func (this *QQuickWidget) InheritPaintEvent(f func(event *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

/*

 */
type QQuickWidget struct {
	*qtwidgets.QWidget
}
type QQuickWidget_ITF interface {
	qtwidgets.QWidget_ITF
	QQuickWidget_PTR() *QQuickWidget
}

func (ptr *QQuickWidget) QQuickWidget_PTR() *QQuickWidget { return ptr }

func (this *QQuickWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QQuickWidget) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = qtwidgets.NewQWidgetFromPointer(cthis)
}
func NewQQuickWidgetFromPointer(cthis unsafe.Pointer) *QQuickWidget {
	bcthis0 := qtwidgets.NewQWidgetFromPointer(cthis)
	return &QQuickWidget{bcthis0}
}
func (*QQuickWidget) NewFromPointer(cthis unsafe.Pointer) *QQuickWidget {
	return NewQQuickWidgetFromPointer(cthis)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QQuickWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickWidget(QWidget *)

/*

 */
func (*QQuickWidget) NewForInherit(parent qtwidgets.QWidget_ITF /*777 QWidget **/) *QQuickWidget {
	return NewQQuickWidget(parent)
}
func NewQQuickWidget(parent qtwidgets.QWidget_ITF /*777 QWidget **/) *QQuickWidget {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickWidget")
	return gothis
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickWidget(QWidget *)

/*

 */
func (*QQuickWidget) NewForInheritp() *QQuickWidget {
	return NewQQuickWidgetp()
}
func NewQQuickWidgetp() *QQuickWidget {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickWidget")
	return gothis
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:68
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QQuickWidget(QQmlEngine *, QWidget *)

/*

 */
func (*QQuickWidget) NewForInherit1(engine qtqml.QQmlEngine_ITF /*777 QQmlEngine **/, parent qtwidgets.QWidget_ITF /*777 QWidget **/) *QQuickWidget {
	return NewQQuickWidget1(engine, parent)
}
func NewQQuickWidget1(engine qtqml.QQmlEngine_ITF /*777 QQmlEngine **/, parent qtwidgets.QWidget_ITF /*777 QWidget **/) *QQuickWidget {
	var convArg0 unsafe.Pointer
	if engine != nil && engine.QQmlEngine_PTR() != nil {
		convArg0 = engine.QQmlEngine_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidgetC2EP10QQmlEngineP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickWidget")
	return gothis
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:69
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QQuickWidget(const QUrl &, QWidget *)

/*

 */
func (*QQuickWidget) NewForInherit2(source qtcore.QUrl_ITF, parent qtwidgets.QWidget_ITF /*777 QWidget **/) *QQuickWidget {
	return NewQQuickWidget2(source, parent)
}
func NewQQuickWidget2(source qtcore.QUrl_ITF, parent qtwidgets.QWidget_ITF /*777 QWidget **/) *QQuickWidget {
	var convArg0 unsafe.Pointer
	if source != nil && source.QUrl_PTR() != nil {
		convArg0 = source.QUrl_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidgetC2ERK4QUrlP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickWidget")
	return gothis
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:69
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QQuickWidget(const QUrl &, QWidget *)

/*

 */
func (*QQuickWidget) NewForInherit2p(source qtcore.QUrl_ITF) *QQuickWidget {
	return NewQQuickWidget2p(source)
}
func NewQQuickWidget2p(source qtcore.QUrl_ITF) *QQuickWidget {
	var convArg0 unsafe.Pointer
	if source != nil && source.QUrl_PTR() != nil {
		convArg0 = source.QUrl_PTR().GetCthis()
	}
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidgetC2ERK4QUrlP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickWidget")
	return gothis
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQuickWidget()

/*

 */
func DeleteQQuickWidget(this *QQuickWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl source() const

/*

 */
func (this *QQuickWidget) Source() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWidget6sourceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlEngine * engine() const

/*

 */
func (this *QQuickWidget) Engine() *qtqml.QQmlEngine /*777 QQmlEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWidget6engineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtqml.NewQQmlEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlContext * rootContext() const

/*

 */
func (this *QQuickWidget) RootContext() *qtqml.QQmlContext /*777 QQmlContext **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWidget11rootContextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtqml.NewQQmlContextFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickItem * rootObject() const

/*

 */
func (this *QQuickWidget) RootObject() *qtquick.QQuickItem /*777 QQuickItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWidget10rootObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtquick.NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] QQuickWidget::ResizeMode resizeMode() const

/*

 */
func (this *QQuickWidget) ResizeMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWidget10resizeModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResizeMode(QQuickWidget::ResizeMode)

/*

 */
func (this *QQuickWidget) SetResizeMode(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget13setResizeModeENS_10ResizeModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] QQuickWidget::Status status() const

/*

 */
func (this *QQuickWidget) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWidget6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:90
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*

 */
func (this *QQuickWidget) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWidget8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize initialSize() const

/*

 */
func (this *QQuickWidget) InitialSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWidget11initialSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const QSurfaceFormat &)

/*

 */
func (this *QQuickWidget) SetFormat(format qtgui.QSurfaceFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QSurfaceFormat_PTR() != nil {
		convArg0 = format.QSurfaceFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget9setFormatERK14QSurfaceFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] QSurfaceFormat format() const

/*

 */
func (this *QQuickWidget) Format() *qtgui.QSurfaceFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWidget6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQSurfaceFormat)
	return rv2
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:96
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage grabFramebuffer() const

/*

 */
func (this *QQuickWidget) GrabFramebuffer() *qtgui.QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWidget15grabFramebufferEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQImage)
	return rv2
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClearColor(const QColor &)

/*

 */
func (this *QQuickWidget) SetClearColor(color qtgui.QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget13setClearColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickWindow * quickWindow() const

/*

 */
func (this *QQuickWidget) QuickWindow() *qtquick.QQuickWindow /*777 QQuickWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWidget11quickWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtquick.NewQQuickWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSource(const QUrl &)

/*

 */
func (this *QQuickWidget) SetSource(arg0 qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QUrl_PTR() != nil {
		convArg0 = arg0.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget9setSourceERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContent(const QUrl &, QQmlComponent *, QObject *)

/*

 */
func (this *QQuickWidget) SetContent(url qtcore.QUrl_ITF, component qtqml.QQmlComponent_ITF /*777 QQmlComponent **/, item qtcore.QObject_ITF /*777 QObject **/) {
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
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget10setContentERK4QUrlP13QQmlComponentP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void statusChanged(QQuickWidget::Status)

/*

 */
func (this *QQuickWidget) StatusChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget13statusChangedENS_6StatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sceneGraphError(QQuickWindow::SceneGraphError, const QString &)

/*

 */
func (this *QQuickWidget) SceneGraphError(error int, message string) {
	var tmpArg1 = qtcore.NewQString5(message)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget15sceneGraphErrorEN12QQuickWindow15SceneGraphErrorERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:118
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*

 */
func (this *QQuickWidget) ResizeEvent(arg0 qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QResizeEvent_PTR() != nil {
		convArg0 = arg0.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:119
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)

/*

 */
func (this *QQuickWidget) TimerEvent(arg0 qtcore.QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTimerEvent_PTR() != nil {
		convArg0 = arg0.QTimerEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:121
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*

 */
func (this *QQuickWidget) KeyPressEvent(arg0 qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QKeyEvent_PTR() != nil {
		convArg0 = arg0.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:122
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)

/*

 */
func (this *QQuickWidget) KeyReleaseEvent(arg0 qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QKeyEvent_PTR() != nil {
		convArg0 = arg0.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:123
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*

 */
func (this *QQuickWidget) MousePressEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:124
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*

 */
func (this *QQuickWidget) MouseReleaseEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:125
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*

 */
func (this *QQuickWidget) MouseMoveEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:126
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)

/*

 */
func (this *QQuickWidget) MouseDoubleClickEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:128
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)

/*

 */
func (this *QQuickWidget) ShowEvent(arg0 qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QShowEvent_PTR() != nil {
		convArg0 = arg0.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:129
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)

/*

 */
func (this *QQuickWidget) HideEvent(arg0 qtgui.QHideEvent_ITF /*777 QHideEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QHideEvent_PTR() != nil {
		convArg0 = arg0.QHideEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget9hideEventEP10QHideEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:131
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*

 */
func (this *QQuickWidget) FocusInEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:132
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*

 */
func (this *QQuickWidget) FocusOutEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:135
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*

 */
func (this *QQuickWidget) WheelEvent(arg0 qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWheelEvent_PTR() != nil {
		convArg0 = arg0.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:139
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QDragEnterEvent *)

/*

 */
func (this *QQuickWidget) DragEnterEvent(arg0 qtgui.QDragEnterEvent_ITF /*777 QDragEnterEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDragEnterEvent_PTR() != nil {
		convArg0 = arg0.QDragEnterEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget14dragEnterEventEP15QDragEnterEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:140
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)

/*

 */
func (this *QQuickWidget) DragMoveEvent(arg0 qtgui.QDragMoveEvent_ITF /*777 QDragMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDragMoveEvent_PTR() != nil {
		convArg0 = arg0.QDragMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:141
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)

/*

 */
func (this *QQuickWidget) DragLeaveEvent(arg0 qtgui.QDragLeaveEvent_ITF /*777 QDragLeaveEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDragLeaveEvent_PTR() != nil {
		convArg0 = arg0.QDragLeaveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget14dragLeaveEventEP15QDragLeaveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:142
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)

/*

 */
func (this *QQuickWidget) DropEvent(arg0 qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDropEvent_PTR() != nil {
		convArg0 = arg0.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:145
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*

 */
func (this *QQuickWidget) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuickWidgets/qquickwidget.h:146
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*

 */
func (this *QQuickWidget) PaintEvent(event qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QPaintEvent_PTR() != nil {
		convArg0 = event.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWidget10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QQuickWidget__ResizeMode = int

//
const QQuickWidget__SizeViewToRootObject QQuickWidget__ResizeMode = 0

//
const QQuickWidget__SizeRootObjectToView QQuickWidget__ResizeMode = 1

func (this *QQuickWidget) ResizeModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QQuickWidget_ResizeModeItemName(val int) string {
	var nilthis *QQuickWidget
	return nilthis.ResizeModeItemName(val)
}

/*


 */
type QQuickWidget__Status = int

//
const QQuickWidget__Null QQuickWidget__Status = 0

//
const QQuickWidget__Ready QQuickWidget__Status = 1

//
const QQuickWidget__Loading QQuickWidget__Status = 2

//
const QQuickWidget__Error QQuickWidget__Status = 3

func (this *QQuickWidget) StatusItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QQuickWidget_StatusItemName(val int) string {
	var nilthis *QQuickWidget
	return nilthis.StatusItemName(val)
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
		qtnetwork.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtwidgets.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
	if false {
		qtquick.KeepMe()
	}
}

//  keep block end
