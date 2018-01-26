package qtwidgets

// /usr/include/qt/QtWidgets/qsplitter.h
// #include <qsplitter.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 37
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
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
type QSplitterHandle struct {
	*QWidget
}

func (this *QSplitterHandle) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QSplitterHandle) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQSplitterHandleFromPointer(cthis unsafe.Pointer) *QSplitterHandle {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QSplitterHandle{bcthis0}
}
func (*QSplitterHandle) NewFromPointer(cthis unsafe.Pointer) *QSplitterHandle {
	return NewQSplitterHandleFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qsplitter.h:138
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QSplitterHandle) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSplitterHandle10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qsplitter.h:140
// index:0
// Public
// void QSplitterHandle(Qt::Orientation, class QSplitter *)
func NewQSplitterHandle(o int, parent *QSplitter /*777 QSplitter **/) *QSplitterHandle {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandleC2EN2Qt11OrientationEP9QSplitter", ffiqt.FFI_TYPE_VOID, cthis, o, convArg1)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), o)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:144
// index:0
// Public
// Qt::Orientation orientation()
func (this *QSplitterHandle) Orientation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSplitterHandle11orientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:145
// index:0
// Public
// bool opaqueResize()
func (this *QSplitterHandle) OpaqueResize() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSplitterHandle12opaqueResizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsplitter.h:146
// index:0
// Public
// QSplitter * splitter()
func (this *QSplitterHandle) Splitter() *QSplitter /*777 QSplitter **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSplitterHandle8splitterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSplitterFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qsplitter.h:148
// index:0
// Public virtual
// QSize sizeHint()
func (this *QSplitterHandle) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSplitterHandle8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qsplitter.h:151
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QSplitterHandle) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:152
// index:0
// Protected virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QSplitterHandle) MouseMoveEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:153
// index:0
// Protected virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QSplitterHandle) MousePressEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:154
// index:0
// Protected virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QSplitterHandle) MouseReleaseEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:155
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QSplitterHandle) ResizeEvent(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:156
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QSplitterHandle) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsplitter.h:158
// index:0
// Protected
// void moveSplitter(int)
func (this *QSplitterHandle) MoveSplitter(p int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle12moveSplitterEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:159
// index:0
// Protected
// int closestLegalPosition(int)
func (this *QSplitterHandle) ClosestLegalPosition(p int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle20closestLegalPositionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), p)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

//  body block end
