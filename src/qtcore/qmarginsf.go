//  header block begin
// /usr/include/qt/QtCore/qmargins.h
// #include <qmargins.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QMarginsF struct {
	*qtrt.CObject
}

func (this *QMarginsF) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQMarginsFFromPointer(cthis unsafe.Pointer) *QMarginsF {
	return &QMarginsF{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qmargins.h:288
// index:0
// Public inline
// void QMarginsF()
func NewQMarginsF() *QMarginsF {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMarginsFC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQMarginsFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qmargins.h:289
// index:1
// Public inline
// void QMarginsF(qreal, qreal, qreal, qreal)
func NewQMarginsF_1(left float64, top float64, right float64, bottom float64) *QMarginsF {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMarginsFC2Edddd", ffiqt.FFI_TYPE_VOID, cthis, &left, &top, &right, &bottom)
	gopp.ErrPrint(err, rv)
	gothis := NewQMarginsFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qmargins.h:290
// index:2
// Public inline
// void QMarginsF(const class QMargins &)
func NewQMarginsF_2(margins *QMargins) *QMarginsF {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = margins.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMarginsFC2ERK8QMargins", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQMarginsFFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qmargins.h:292
// index:0
// Public inline
// bool isNull()
func (this *QMarginsF) IsNull() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMarginsF6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmargins.h:294
// index:0
// Public inline
// qreal left()
func (this *QMarginsF) Left() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMarginsF4leftEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmargins.h:295
// index:0
// Public inline
// qreal top()
func (this *QMarginsF) Top() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMarginsF3topEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmargins.h:296
// index:0
// Public inline
// qreal right()
func (this *QMarginsF) Right() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMarginsF5rightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmargins.h:297
// index:0
// Public inline
// qreal bottom()
func (this *QMarginsF) Bottom() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMarginsF6bottomEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmargins.h:299
// index:0
// Public
// void setLeft(qreal)
func (this *QMarginsF) SetLeft(left float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMarginsF7setLeftEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &left)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:300
// index:0
// Public
// void setTop(qreal)
func (this *QMarginsF) SetTop(top float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMarginsF6setTopEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &top)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:301
// index:0
// Public
// void setRight(qreal)
func (this *QMarginsF) SetRight(right float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMarginsF8setRightEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &right)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:302
// index:0
// Public
// void setBottom(qreal)
func (this *QMarginsF) SetBottom(bottom float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QMarginsF9setBottomEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:311
// index:0
// Public inline
// QMargins toMargins()
func (this *QMarginsF) ToMargins() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QMarginsF9toMarginsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
