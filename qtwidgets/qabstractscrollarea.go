package qtwidgets

// /usr/include/qt/QtWidgets/qabstractscrollarea.h
// #include <qabstractscrollarea.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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
// void setViewportMargins(int, int, int, int)
func (this *QAbstractScrollArea) InheritSetViewportMargins(f func(left int, top int, right int, bottom int)) {
	ffiqt.SetAllInheritCallback(this, "setViewportMargins", f)
}

// QMargins viewportMargins()
func (this *QAbstractScrollArea) InheritViewportMargins(f func() unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "viewportMargins", f)
}

// bool eventFilter(class QObject *, class QEvent *)
func (this *QAbstractScrollArea) InheritEventFilter(f func(arg0 *qtcore.QObject /*777 QObject **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool) {
	ffiqt.SetAllInheritCallback(this, "eventFilter", f)
}

// bool event(class QEvent *)
func (this *QAbstractScrollArea) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	ffiqt.SetAllInheritCallback(this, "event", f)
}

// bool viewportEvent(class QEvent *)
func (this *QAbstractScrollArea) InheritViewportEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	ffiqt.SetAllInheritCallback(this, "viewportEvent", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QAbstractScrollArea) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QAbstractScrollArea) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "paintEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QAbstractScrollArea) InheritMousePressEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QAbstractScrollArea) InheritMouseReleaseEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QAbstractScrollArea) InheritMouseDoubleClickEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QAbstractScrollArea) InheritMouseMoveEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void wheelEvent(class QWheelEvent *)
func (this *QAbstractScrollArea) InheritWheelEvent(f func(arg0 *qtgui.QWheelEvent /*777 QWheelEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void contextMenuEvent(class QContextMenuEvent *)
func (this *QAbstractScrollArea) InheritContextMenuEvent(f func(arg0 *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void dragEnterEvent(class QDragEnterEvent *)
func (this *QAbstractScrollArea) InheritDragEnterEvent(f func(arg0 *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragMoveEvent(class QDragMoveEvent *)
func (this *QAbstractScrollArea) InheritDragMoveEvent(f func(arg0 *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dragLeaveEvent(class QDragLeaveEvent *)
func (this *QAbstractScrollArea) InheritDragLeaveEvent(f func(arg0 *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dropEvent(class QDropEvent *)
func (this *QAbstractScrollArea) InheritDropEvent(f func(arg0 *qtgui.QDropEvent /*777 QDropEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "dropEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QAbstractScrollArea) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void scrollContentsBy(int, int)
func (this *QAbstractScrollArea) InheritScrollContentsBy(f func(dx int, dy int)) {
	ffiqt.SetAllInheritCallback(this, "scrollContentsBy", f)
}

// QSize viewportSizeHint()
func (this *QAbstractScrollArea) InheritViewportSizeHint(f func() unsafe.Pointer) {
	ffiqt.SetAllInheritCallback(this, "viewportSizeHint", f)
}

type QAbstractScrollArea struct {
	*QFrame
}

func (this *QAbstractScrollArea) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QFrame.GetCthis()
	}
}
func (this *QAbstractScrollArea) SetCthis(cthis unsafe.Pointer) {
	this.QFrame = NewQFrameFromPointer(cthis)
}
func NewQAbstractScrollAreaFromPointer(cthis unsafe.Pointer) *QAbstractScrollArea {
	bcthis0 := NewQFrameFromPointer(cthis)
	return &QAbstractScrollArea{bcthis0}
}
func (*QAbstractScrollArea) NewFromPointer(cthis unsafe.Pointer) *QAbstractScrollArea {
	return NewQAbstractScrollAreaFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QAbstractScrollArea) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractScrollArea(QWidget *)
func NewQAbstractScrollArea(parent *QWidget /*777 QWidget **/) *QAbstractScrollArea {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollAreaC2EP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractScrollAreaFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractScrollArea()
func DeleteQAbstractScrollArea(this *QAbstractScrollArea) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollAreaD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:74
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ScrollBarPolicy verticalScrollBarPolicy()
func (this *QAbstractScrollArea) VerticalScrollBarPolicy() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea23verticalScrollBarPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalScrollBarPolicy(Qt::ScrollBarPolicy)
func (this *QAbstractScrollArea) SetVerticalScrollBarPolicy(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea26setVerticalScrollBarPolicyEN2Qt15ScrollBarPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QScrollBar * verticalScrollBar()
func (this *QAbstractScrollArea) VerticalScrollBar() *QScrollBar /*777 QScrollBar **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea17verticalScrollBarEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQScrollBarFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalScrollBar(QScrollBar *)
func (this *QAbstractScrollArea) SetVerticalScrollBar(scrollbar *QScrollBar /*777 QScrollBar **/) {
	var convArg0 = scrollbar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea20setVerticalScrollBarEP10QScrollBar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ScrollBarPolicy horizontalScrollBarPolicy()
func (this *QAbstractScrollArea) HorizontalScrollBarPolicy() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea25horizontalScrollBarPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalScrollBarPolicy(Qt::ScrollBarPolicy)
func (this *QAbstractScrollArea) SetHorizontalScrollBarPolicy(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea28setHorizontalScrollBarPolicyEN2Qt15ScrollBarPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QScrollBar * horizontalScrollBar()
func (this *QAbstractScrollArea) HorizontalScrollBar() *QScrollBar /*777 QScrollBar **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea19horizontalScrollBarEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQScrollBarFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalScrollBar(QScrollBar *)
func (this *QAbstractScrollArea) SetHorizontalScrollBar(scrollbar *QScrollBar /*777 QScrollBar **/) {
	var convArg0 = scrollbar.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea22setHorizontalScrollBarEP10QScrollBar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * cornerWidget()
func (this *QAbstractScrollArea) CornerWidget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea12cornerWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCornerWidget(QWidget *)
func (this *QAbstractScrollArea) SetCornerWidget(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea15setCornerWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addScrollBarWidget(QWidget *, Qt::Alignment)
func (this *QAbstractScrollArea) AddScrollBarWidget(widget *QWidget /*777 QWidget **/, alignment int) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea18addScrollBarWidgetEP7QWidget6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * viewport()
func (this *QAbstractScrollArea) Viewport() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea8viewportEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewport(QWidget *)
func (this *QAbstractScrollArea) SetViewport(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea11setViewportEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize maximumViewportSize()
func (this *QAbstractScrollArea) MaximumViewportSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea19maximumViewportSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:94
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QAbstractScrollArea) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QAbstractScrollArea) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:98
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setupViewport(QWidget *)
func (this *QAbstractScrollArea) SetupViewport(viewport *QWidget /*777 QWidget **/) {
	var convArg0 = viewport.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea13setupViewportEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractScrollArea::SizeAdjustPolicy sizeAdjustPolicy()
func (this *QAbstractScrollArea) SizeAdjustPolicy() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea16sizeAdjustPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSizeAdjustPolicy(enum QAbstractScrollArea::SizeAdjustPolicy)
func (this *QAbstractScrollArea) SetSizeAdjustPolicy(policy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea19setSizeAdjustPolicyENS_16SizeAdjustPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:105
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setViewportMargins(int, int, int, int)
func (this *QAbstractScrollArea) SetViewportMargins(left int, top int, right int, bottom int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea18setViewportMarginsEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:106
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void setViewportMargins(const QMargins &)
func (this *QAbstractScrollArea) SetViewportMargins_1(margins *qtcore.QMargins) {
	var convArg0 = margins.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea18setViewportMarginsERK8QMargins", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:107
// index:0
// Protected Visibility=Default Availability=Available
// [16] QMargins viewportMargins()
func (this *QAbstractScrollArea) ViewportMargins() *qtcore.QMargins /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea15viewportMarginsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMarginsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQMargins)
	return rv2
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:109
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)
func (this *QAbstractScrollArea) EventFilter(arg0 *qtcore.QObject /*777 QObject **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:110
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QAbstractScrollArea) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:111
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool viewportEvent(QEvent *)
func (this *QAbstractScrollArea) ViewportEvent(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea13viewportEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:113
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QAbstractScrollArea) ResizeEvent(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:114
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QAbstractScrollArea) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:115
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QAbstractScrollArea) MousePressEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:116
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QAbstractScrollArea) MouseReleaseEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:117
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)
func (this *QAbstractScrollArea) MouseDoubleClickEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:118
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QAbstractScrollArea) MouseMoveEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:120
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)
func (this *QAbstractScrollArea) WheelEvent(arg0 *qtgui.QWheelEvent /*777 QWheelEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:123
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)
func (this *QAbstractScrollArea) ContextMenuEvent(arg0 *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:126
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QDragEnterEvent *)
func (this *QAbstractScrollArea) DragEnterEvent(arg0 *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea14dragEnterEventEP15QDragEnterEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:127
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)
func (this *QAbstractScrollArea) DragMoveEvent(arg0 *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea13dragMoveEventEP14QDragMoveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:128
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)
func (this *QAbstractScrollArea) DragLeaveEvent(arg0 *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea14dragLeaveEventEP15QDragLeaveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:129
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)
func (this *QAbstractScrollArea) DropEvent(arg0 *qtgui.QDropEvent /*777 QDropEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:132
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QAbstractScrollArea) KeyPressEvent(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:134
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void scrollContentsBy(int, int)
func (this *QAbstractScrollArea) ScrollContentsBy(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractScrollArea16scrollContentsByEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractscrollarea.h:136
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSize viewportSizeHint()
func (this *QAbstractScrollArea) ViewportSizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractScrollArea16viewportSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

type QAbstractScrollArea__SizeAdjustPolicy = int

const QAbstractScrollArea__AdjustIgnored QAbstractScrollArea__SizeAdjustPolicy = 0
const QAbstractScrollArea__AdjustToContentsOnFirstShow QAbstractScrollArea__SizeAdjustPolicy = 1
const QAbstractScrollArea__AdjustToContents QAbstractScrollArea__SizeAdjustPolicy = 2

//  body block end
