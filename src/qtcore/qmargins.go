//  header block begin
// /usr/include/qt/QtCore/qmargins.h
// #include <qmargins.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
type QMargins struct {
	*qtrt.CObject
}

func (this *QMargins) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qmargins.h:54
// index:0
// inline
// void QMargins()
func NewQMargins() *QMargins {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMarginsC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQMarginsFromPointer(cthis)
	return gothis
}
func NewQMarginsFromPointer(cthis unsafe.Pointer) *QMargins {
	return &QMargins{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qmargins.h:55
// index:1
// inline
// void QMargins(int, int, int, int)
func NewQMargins_1(left int, top int, right int, bottom int) *QMargins {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMarginsC2Eiiii", ffiqt.FFI_TYPE_VOID, cthis, &left, &top, &right, &bottom)
	gopp.ErrPrint(err, rv)
	gothis := NewQMarginsFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qmargins.h:57
// index:0
// inline
// bool isNull()
func (this *QMargins) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMargins6isNullEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:59
// index:0
// inline
// int left()
func (this *QMargins) Left() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMargins4leftEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:60
// index:0
// inline
// int top()
func (this *QMargins) Top() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMargins3topEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:61
// index:0
// inline
// int right()
func (this *QMargins) Right() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMargins5rightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:62
// index:0
// inline
// int bottom()
func (this *QMargins) Bottom() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMargins6bottomEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:64
// index:0
// void setLeft(int)
func (this *QMargins) SetLeft(left int) {
	// 0: (, left int), (&left)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMargins7setLeftEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &left)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:65
// index:0
// void setTop(int)
func (this *QMargins) SetTop(top int) {
	// 0: (, top int), (&top)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMargins6setTopEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &top)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:66
// index:0
// void setRight(int)
func (this *QMargins) SetRight(right int) {
	// 0: (, right int), (&right)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMargins8setRightEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &right)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:67
// index:0
// void setBottom(int)
func (this *QMargins) SetBottom(bottom int) {
	// 0: (, bottom int), (&bottom)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMargins9setBottomEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &bottom)
	gopp.ErrPrint(err, rv)
}

//  body block end
