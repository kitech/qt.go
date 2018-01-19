//  header block begin
// /usr/include/qt/QtWidgets/qsplitter.h
// #include <qsplitter.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 30
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qsplitter.h:138
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QSplitterHandle) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSplitterHandle10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:140
// index:0
// void QSplitterHandle(Qt::Orientation, class QSplitter *)
func NewQSplitterHandle(o int, parent unsafe.Pointer) *QSplitterHandle {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandleC2EN2Qt11OrientationEP9QSplitter", ffiqt.FFI_TYPE_VOID, cthis, &o, parent)
	gopp.ErrPrint(err, rv)
	return &QSplitterHandle{cthis}
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
	// 0: (, Qt::Orientation o), (&o)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSplitterHandle14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_VOID, this.cthis, &o)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:144
// index:0
// Qt::Orientation orientation()
func (this *QSplitterHandle) Orientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSplitterHandle11orientationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:145
// index:0
// bool opaqueResize()
func (this *QSplitterHandle) OpaqueResize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSplitterHandle12opaqueResizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:146
// index:0
// QSplitter * splitter()
func (this *QSplitterHandle) Splitter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSplitterHandle8splitterEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:148
// index:0
// virtual
// QSize sizeHint()
func (this *QSplitterHandle) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSplitterHandle8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
