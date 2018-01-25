package qtwidgets

// /usr/include/qt/QtWidgets/qscrollarea.h
// #include <qscrollarea.h>
// #include <QtWidgets>

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
	if this == nil {
		return nil
	} else {
		return this.QAbstractScrollArea.GetCthis()
	}
}
func (this *QScrollArea) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractScrollArea = NewQAbstractScrollAreaFromPointer(cthis)
}
func NewQScrollAreaFromPointer(cthis unsafe.Pointer) *QScrollArea {
	bcthis0 := NewQAbstractScrollAreaFromPointer(cthis)
	return &QScrollArea{bcthis0}
}
func (*QScrollArea) NewFromPointer(cthis unsafe.Pointer) *QScrollArea {
	return NewQScrollAreaFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QScrollArea) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QScrollArea10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qscrollarea.h:59
// index:0
// Public
// void QScrollArea(class QWidget *)
func NewQScrollArea(parent *QWidget /*444 QWidget **/) *QScrollArea {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollAreaC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQScrollAreaFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qscrollarea.h:60
// index:0
// Public virtual
// void ~QScrollArea()
func DeleteQScrollArea(*QScrollArea) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollAreaD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:62
// index:0
// Public
// QWidget * widget()
func (this *QScrollArea) Widget() *QWidget /*444 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QScrollArea6widgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qscrollarea.h:63
// index:0
// Public
// void setWidget(class QWidget *)
func (this *QScrollArea) SetWidget(widget *QWidget /*444 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea9setWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:64
// index:0
// Public
// QWidget * takeWidget()
func (this *QScrollArea) TakeWidget() *QWidget /*444 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea10takeWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qscrollarea.h:66
// index:0
// Public
// bool widgetResizable()
func (this *QScrollArea) WidgetResizable() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QScrollArea15widgetResizableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qscrollarea.h:67
// index:0
// Public
// void setWidgetResizable(_Bool)
func (this *QScrollArea) SetWidgetResizable(resizable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea18setWidgetResizableEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), resizable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:69
// index:0
// Public virtual
// QSize sizeHint()
func (this *QScrollArea) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QScrollArea8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qscrollarea.h:71
// index:0
// Public virtual
// bool focusNextPrevChild(_Bool)
func (this *QScrollArea) FocusNextPrevChild(next bool) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea18focusNextPrevChildEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), next)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qscrollarea.h:73
// index:0
// Public
// Qt::Alignment alignment()
func (this *QScrollArea) Alignment() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QScrollArea9alignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:76
// index:0
// Public
// void ensureVisible(int, int, int, int)
func (this *QScrollArea) EnsureVisible(x int, y int, xmargin int, ymargin int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea13ensureVisibleEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, xmargin, ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:77
// index:0
// Public
// void ensureWidgetVisible(class QWidget *, int, int)
func (this *QScrollArea) EnsureWidgetVisible(childWidget *QWidget /*444 QWidget **/, xmargin int, ymargin int) {
	var convArg0 = childWidget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea19ensureWidgetVisibleEP7QWidgetii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:81
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QScrollArea) Event(arg0 *qtcore.QEvent /*444 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qscrollarea.h:82
// index:0
// Protected virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QScrollArea) EventFilter(arg0 *qtcore.QObject /*444 QObject **/, arg1 *qtcore.QEvent /*444 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qscrollarea.h:83
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QScrollArea) ResizeEvent(arg0 *qtgui.QResizeEvent /*444 QResizeEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:84
// index:0
// Protected virtual
// void scrollContentsBy(int, int)
func (this *QScrollArea) ScrollContentsBy(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QScrollArea16scrollContentsByEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qscrollarea.h:86
// index:0
// Protected virtual
// QSize viewportSizeHint()
func (this *QScrollArea) ViewportSizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QScrollArea16viewportSizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
