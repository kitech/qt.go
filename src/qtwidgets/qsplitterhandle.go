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
func NewQSplitterHandleFromPointer(cthis unsafe.Pointer) *QSplitterHandle {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QSplitterHandle{bcthis0}
}

// /usr/include/qt/QtWidgets/qsplitter.h:138
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QSplitterHandle) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSplitterHandle10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:140
// index:0
// Public
// void QSplitterHandle(Qt::Orientation, class QSplitter *)
func NewQSplitterHandle(o int, parent unsafe.Pointer) *QSplitterHandle {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandleC2EN2Qt11OrientationEP9QSplitter", ffiqt.FFI_TYPE_VOID, cthis, &o, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSplitterHandleFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qsplitter.h:141
// index:0
// Public virtual
// void ~QSplitterHandle()
func DeleteQSplitterHandle(*QSplitterHandle) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandleD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:143
// index:0
// Public
// void setOrientation(Qt::Orientation)
func (this *QSplitterHandle) SetOrientation(o int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &o)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:144
// index:0
// Public
// Qt::Orientation orientation()
func (this *QSplitterHandle) Orientation() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSplitterHandle11orientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:145
// index:0
// Public
// bool opaqueResize()
func (this *QSplitterHandle) OpaqueResize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSplitterHandle12opaqueResizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:146
// index:0
// Public
// QSplitter * splitter()
func (this *QSplitterHandle) Splitter() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSplitterHandle8splitterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:148
// index:0
// Public virtual
// QSize sizeHint()
func (this *QSplitterHandle) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSplitterHandle8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:151
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QSplitterHandle) PaintEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:152
// index:0
// Protected virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QSplitterHandle) MouseMoveEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:153
// index:0
// Protected virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QSplitterHandle) MousePressEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:154
// index:0
// Protected virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QSplitterHandle) MouseReleaseEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:155
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QSplitterHandle) ResizeEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:156
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QSplitterHandle) Event(arg0 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:158
// index:0
// Protected
// void moveSplitter(int)
func (this *QSplitterHandle) MoveSplitter(p int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle12moveSplitterEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:159
// index:0
// Protected
// int closestLegalPosition(int)
func (this *QSplitterHandle) ClosestLegalPosition(p int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle20closestLegalPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &p)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
