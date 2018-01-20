//  header block begin
// /usr/include/qt/QtWidgets/qabstractscrollarea.h
// #include <qabstractscrollarea.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
}

//  ext block end

//  body block begin
type QAbstractScrollArea struct {
	*QFrame
}

func (this *QAbstractScrollArea) GetCthis() unsafe.Pointer {
	return this.QFrame.GetCthis()
}
func NewQAbstractScrollAreaFromPointer(cthis unsafe.Pointer) *QAbstractScrollArea {
	bcthis0 := NewQFrameFromPointer(cthis)
	return &QAbstractScrollArea{bcthis0}
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:57
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QAbstractScrollArea) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:64
// index:0
// Public
// void QAbstractScrollArea(class QWidget *)
func NewQAbstractScrollArea(parent unsafe.Pointer) *QAbstractScrollArea {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollAreaC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractScrollAreaFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:65
// index:0
// Public virtual
// void ~QAbstractScrollArea()
func DeleteQAbstractScrollArea(*QAbstractScrollArea) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollAreaD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:74
// index:0
// Public
// Qt::ScrollBarPolicy verticalScrollBarPolicy()
func (this *QAbstractScrollArea) VerticalScrollBarPolicy() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea23verticalScrollBarPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:75
// index:0
// Public
// void setVerticalScrollBarPolicy(Qt::ScrollBarPolicy)
func (this *QAbstractScrollArea) SetVerticalScrollBarPolicy(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea26setVerticalScrollBarPolicyEN2Qt15ScrollBarPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:76
// index:0
// Public
// QScrollBar * verticalScrollBar()
func (this *QAbstractScrollArea) VerticalScrollBar() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea17verticalScrollBarEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:77
// index:0
// Public
// void setVerticalScrollBar(class QScrollBar *)
func (this *QAbstractScrollArea) SetVerticalScrollBar(scrollbar unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea20setVerticalScrollBarEP10QScrollBar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), scrollbar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:79
// index:0
// Public
// Qt::ScrollBarPolicy horizontalScrollBarPolicy()
func (this *QAbstractScrollArea) HorizontalScrollBarPolicy() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea25horizontalScrollBarPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:80
// index:0
// Public
// void setHorizontalScrollBarPolicy(Qt::ScrollBarPolicy)
func (this *QAbstractScrollArea) SetHorizontalScrollBarPolicy(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea28setHorizontalScrollBarPolicyEN2Qt15ScrollBarPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:81
// index:0
// Public
// QScrollBar * horizontalScrollBar()
func (this *QAbstractScrollArea) HorizontalScrollBar() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea19horizontalScrollBarEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:82
// index:0
// Public
// void setHorizontalScrollBar(class QScrollBar *)
func (this *QAbstractScrollArea) SetHorizontalScrollBar(scrollbar unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea22setHorizontalScrollBarEP10QScrollBar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), scrollbar)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:84
// index:0
// Public
// QWidget * cornerWidget()
func (this *QAbstractScrollArea) CornerWidget() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea12cornerWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:85
// index:0
// Public
// void setCornerWidget(class QWidget *)
func (this *QAbstractScrollArea) SetCornerWidget(widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea15setCornerWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:90
// index:0
// Public
// QWidget * viewport()
func (this *QAbstractScrollArea) Viewport() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea8viewportEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:91
// index:0
// Public
// void setViewport(class QWidget *)
func (this *QAbstractScrollArea) SetViewport(widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea11setViewportEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:92
// index:0
// Public
// QSize maximumViewportSize()
func (this *QAbstractScrollArea) MaximumViewportSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea19maximumViewportSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:94
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QAbstractScrollArea) MinimumSizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:96
// index:0
// Public virtual
// QSize sizeHint()
func (this *QAbstractScrollArea) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:98
// index:0
// Public virtual
// void setupViewport(class QWidget *)
func (this *QAbstractScrollArea) SetupViewport(viewport unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea13setupViewportEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), viewport)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:100
// index:0
// Public
// QAbstractScrollArea::SizeAdjustPolicy sizeAdjustPolicy()
func (this *QAbstractScrollArea) SizeAdjustPolicy() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea16sizeAdjustPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:101
// index:0
// Public
// void setSizeAdjustPolicy(enum QAbstractScrollArea::SizeAdjustPolicy)
func (this *QAbstractScrollArea) SetSizeAdjustPolicy(policy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea19setSizeAdjustPolicyENS_16SizeAdjustPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:105
// index:0
// Protected
// void setViewportMargins(int, int, int, int)
func (this *QAbstractScrollArea) SetViewportMargins(left int, top int, right int, bottom int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea18setViewportMarginsEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &left, &top, &right, &bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:106
// index:1
// Protected
// void setViewportMargins(const class QMargins &)
func (this *QAbstractScrollArea) SetViewportMargins_1(margins *qtcore.QMargins) {
	var convArg0 = margins.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea18setViewportMarginsERK8QMargins", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:107
// index:0
// Protected
// QMargins viewportMargins()
func (this *QAbstractScrollArea) ViewportMargins() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea15viewportMarginsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:109
// index:0
// Protected virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QAbstractScrollArea) EventFilter(arg0 unsafe.Pointer, arg1 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:110
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QAbstractScrollArea) Event(arg0 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:111
// index:0
// Protected virtual
// bool viewportEvent(class QEvent *)
func (this *QAbstractScrollArea) ViewportEvent(arg0 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea13viewportEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:113
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QAbstractScrollArea) ResizeEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:114
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QAbstractScrollArea) PaintEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:115
// index:0
// Protected virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QAbstractScrollArea) MousePressEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:116
// index:0
// Protected virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QAbstractScrollArea) MouseReleaseEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:117
// index:0
// Protected virtual
// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QAbstractScrollArea) MouseDoubleClickEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:118
// index:0
// Protected virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QAbstractScrollArea) MouseMoveEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:120
// index:0
// Protected virtual
// void wheelEvent(class QWheelEvent *)
func (this *QAbstractScrollArea) WheelEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:123
// index:0
// Protected virtual
// void contextMenuEvent(class QContextMenuEvent *)
func (this *QAbstractScrollArea) ContextMenuEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:126
// index:0
// Protected virtual
// void dragEnterEvent(class QDragEnterEvent *)
func (this *QAbstractScrollArea) DragEnterEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea14dragEnterEventEP15QDragEnterEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:127
// index:0
// Protected virtual
// void dragMoveEvent(class QDragMoveEvent *)
func (this *QAbstractScrollArea) DragMoveEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea13dragMoveEventEP14QDragMoveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:128
// index:0
// Protected virtual
// void dragLeaveEvent(class QDragLeaveEvent *)
func (this *QAbstractScrollArea) DragLeaveEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea14dragLeaveEventEP15QDragLeaveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:129
// index:0
// Protected virtual
// void dropEvent(class QDropEvent *)
func (this *QAbstractScrollArea) DropEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:132
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QAbstractScrollArea) KeyPressEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:134
// index:0
// Protected virtual
// void scrollContentsBy(int, int)
func (this *QAbstractScrollArea) ScrollContentsBy(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea16scrollContentsByEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:136
// index:0
// Protected virtual
// QSize viewportSizeHint()
func (this *QAbstractScrollArea) ViewportSizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea16viewportSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
