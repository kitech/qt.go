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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qscrollarea.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QScrollArea) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QScrollArea10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:59
// index:0
// void QScrollArea(class QWidget *)
func NewQScrollArea(parent unsafe.Pointer) *QScrollArea {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollAreaC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QScrollArea{cthis}
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QScrollArea6widgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:63
// index:0
// void setWidget(class QWidget *)
func (this *QScrollArea) SetWidget(widget unsafe.Pointer) {
	// 0: (, QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea9setWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:64
// index:0
// QWidget * takeWidget()
func (this *QScrollArea) TakeWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea10takeWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:66
// index:0
// bool widgetResizable()
func (this *QScrollArea) WidgetResizable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QScrollArea15widgetResizableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:67
// index:0
// void setWidgetResizable(_Bool)
func (this *QScrollArea) SetWidgetResizable(resizable bool) {
	// 0: (, bool resizable), (&resizable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea18setWidgetResizableEb", ffiqt.FFI_TYPE_VOID, this.cthis, &resizable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:69
// index:0
// virtual
// QSize sizeHint()
func (this *QScrollArea) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QScrollArea8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:71
// index:0
// virtual
// bool focusNextPrevChild(_Bool)
func (this *QScrollArea) FocusNextPrevChild(next bool) {
	// 0: (, bool next), (&next)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea18focusNextPrevChildEb", ffiqt.FFI_TYPE_VOID, this.cthis, &next)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:73
// index:0
// Qt::Alignment alignment()
func (this *QScrollArea) Alignment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QScrollArea9alignmentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:76
// index:0
// void ensureVisible(int, int, int, int)
func (this *QScrollArea) EnsureVisible(x int, y int, xmargin int, ymargin int) {
	// 0: (, int x, int y, int xmargin, int ymargin), (&x, &y, &xmargin, &ymargin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea13ensureVisibleEiiii", ffiqt.FFI_TYPE_VOID, this.cthis, &x, &y, &xmargin, &ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:77
// index:0
// void ensureWidgetVisible(class QWidget *, int, int)
func (this *QScrollArea) EnsureWidgetVisible(childWidget unsafe.Pointer, xmargin int, ymargin int) {
	// 0: (, QWidget * childWidget, int xmargin, int ymargin), (childWidget, &xmargin, &ymargin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea19ensureWidgetVisibleEP7QWidgetii", ffiqt.FFI_TYPE_VOID, this.cthis, childWidget, &xmargin, &ymargin)
	gopp.ErrPrint(err, rv)
}

//  body block end
