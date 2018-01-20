//  header block begin
// /usr/include/qt/QtWidgets/qscrollarea.h
// #include <qscrollarea.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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
type QScrollArea struct {
	*QAbstractScrollArea
}

func (this *QScrollArea) GetCthis() unsafe.Pointer {
	return this.QAbstractScrollArea.GetCthis()
}

// /usr/include/qt/QtWidgets/qscrollarea.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QScrollArea) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QScrollArea10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:59
// index:0
// void QScrollArea(class QWidget *)
func NewQScrollArea(parent unsafe.Pointer) *QScrollArea {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollAreaC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQScrollAreaFromPointer(cthis)
	return gothis
}
func NewQScrollAreaFromPointer(cthis unsafe.Pointer) *QScrollArea {
	bcthis0 := NewQAbstractScrollAreaFromPointer(cthis)
	return &QScrollArea{bcthis0}
}

// /usr/include/qt/QtWidgets/qscrollarea.h:80
// index:1
// void QScrollArea(class QScrollAreaPrivate &, class QWidget *)
func NewQScrollArea_1(dd unsafe.Pointer, parent unsafe.Pointer) *QScrollArea {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollAreaC2ER18QScrollAreaPrivateP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQScrollAreaFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qscrollarea.h:60
// index:0
// virtual
// void ~QScrollArea()
func DeleteQScrollArea(*QScrollArea) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollAreaD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:62
// index:0
// QWidget * widget()
func (this *QScrollArea) Widget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QScrollArea6widgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:63
// index:0
// void setWidget(class QWidget *)
func (this *QScrollArea) SetWidget(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea9setWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:64
// index:0
// QWidget * takeWidget()
func (this *QScrollArea) TakeWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea10takeWidgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:66
// index:0
// bool widgetResizable()
func (this *QScrollArea) WidgetResizable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QScrollArea15widgetResizableEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:67
// index:0
// void setWidgetResizable(_Bool)
func (this *QScrollArea) SetWidgetResizable(resizable bool) {
	// 0: (, resizable bool), (&resizable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea18setWidgetResizableEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &resizable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:69
// index:0
// virtual
// QSize sizeHint()
func (this *QScrollArea) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QScrollArea8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:71
// index:0
// virtual
// bool focusNextPrevChild(_Bool)
func (this *QScrollArea) FocusNextPrevChild(next bool) {
	// 0: (, next bool), (&next)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea18focusNextPrevChildEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &next)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:73
// index:0
// Qt::Alignment alignment()
func (this *QScrollArea) Alignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QScrollArea9alignmentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:74
// index:0
// void setAlignment(Qt::Alignment)
func (this *QScrollArea) SetAlignment(arg0 int) {
	// 0: (, Qt::Alignment arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:76
// index:0
// void ensureVisible(int, int, int, int)
func (this *QScrollArea) EnsureVisible(x int, y int, xmargin int, ymargin int) {
	// 0: (, x int, y int, xmargin int, ymargin int), (&x, &y, &xmargin, &ymargin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea13ensureVisibleEiiii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &xmargin, &ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:77
// index:0
// void ensureWidgetVisible(class QWidget *, int, int)
func (this *QScrollArea) EnsureWidgetVisible(childWidget unsafe.Pointer, xmargin int, ymargin int) {
	// 0: (, childWidget QWidget *, xmargin int, ymargin int), (childWidget, &xmargin, &ymargin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea19ensureWidgetVisibleEP7QWidgetii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), childWidget, &xmargin, &ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:81
// index:0
// virtual
// bool event(class QEvent *)
func (this *QScrollArea) Event(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:82
// index:0
// virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QScrollArea) EventFilter(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	// 0: (, QObject * arg0, QEvent * arg1), (arg0, arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:83
// index:0
// virtual
// void resizeEvent(class QResizeEvent *)
func (this *QScrollArea) ResizeEvent(arg0 unsafe.Pointer) {
	// 0: (, QResizeEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:84
// index:0
// virtual
// void scrollContentsBy(int, int)
func (this *QScrollArea) ScrollContentsBy(dx int, dy int) {
	// 0: (, dx int, dy int), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea16scrollContentsByEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:86
// index:0
// virtual
// QSize viewportSizeHint()
func (this *QScrollArea) ViewportSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QScrollArea16viewportSizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
