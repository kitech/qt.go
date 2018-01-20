//  header block begin
// /usr/include/qt/QtWidgets/qsplitter.h
// #include <qsplitter.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 38
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
type QSplitterHandle struct {
	*QWidget
}

func (this *QSplitterHandle) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}

// /usr/include/qt/QtWidgets/qsplitter.h:138
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QSplitterHandle) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSplitterHandle10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:140
// index:0
// void QSplitterHandle(Qt::Orientation, class QSplitter *)
func NewQSplitterHandle(o int, parent unsafe.Pointer) *QSplitterHandle {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandleC2EN2Qt11OrientationEP9QSplitter", ffiqt.FFI_TYPE_VOID, cthis, &o, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSplitterHandleFromPointer(cthis)
	return gothis
}
func NewQSplitterHandleFromPointer(cthis unsafe.Pointer) *QSplitterHandle {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QSplitterHandle{bcthis0}
}

// /usr/include/qt/QtWidgets/qsplitter.h:141
// index:0
// virtual
// void ~QSplitterHandle()
func DeleteQSplitterHandle(*QSplitterHandle) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandleD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:143
// index:0
// void setOrientation(Qt::Orientation)
func (this *QSplitterHandle) SetOrientation(o int) {
	// 0: (, o Qt::Orientation), (&o)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &o)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:144
// index:0
// Qt::Orientation orientation()
func (this *QSplitterHandle) Orientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSplitterHandle11orientationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:145
// index:0
// bool opaqueResize()
func (this *QSplitterHandle) OpaqueResize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSplitterHandle12opaqueResizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:146
// index:0
// QSplitter * splitter()
func (this *QSplitterHandle) Splitter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSplitterHandle8splitterEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:148
// index:0
// virtual
// QSize sizeHint()
func (this *QSplitterHandle) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSplitterHandle8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:151
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QSplitterHandle) PaintEvent(arg0 unsafe.Pointer) {
	// 0: (, QPaintEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:152
// index:0
// virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QSplitterHandle) MouseMoveEvent(arg0 unsafe.Pointer) {
	// 0: (, QMouseEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:153
// index:0
// virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QSplitterHandle) MousePressEvent(arg0 unsafe.Pointer) {
	// 0: (, QMouseEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:154
// index:0
// virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QSplitterHandle) MouseReleaseEvent(arg0 unsafe.Pointer) {
	// 0: (, QMouseEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:155
// index:0
// virtual
// void resizeEvent(class QResizeEvent *)
func (this *QSplitterHandle) ResizeEvent(arg0 unsafe.Pointer) {
	// 0: (, QResizeEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:156
// index:0
// virtual
// bool event(class QEvent *)
func (this *QSplitterHandle) Event(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:158
// index:0
// void moveSplitter(int)
func (this *QSplitterHandle) MoveSplitter(p int) {
	// 0: (, p int), (&p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle12moveSplitterEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:159
// index:0
// int closestLegalPosition(int)
func (this *QSplitterHandle) ClosestLegalPosition(p int) {
	// 0: (, p int), (&p)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle20closestLegalPositionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &p)
	gopp.ErrPrint(err, rv)
}

//  body block end
