package qtcore

// /usr/include/qt/QtCore/qmargins.h
// #include <qmargins.h>
// #include <QtCore>

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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMargins) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQMarginsFromPointer(cthis unsafe.Pointer) *QMargins {
	return &QMargins{&qtrt.CObject{cthis}}
}
func (*QMargins) NewFromPointer(cthis unsafe.Pointer) *QMargins {
	return NewQMarginsFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmargins.h:54
// index:0
// Public inline
// void QMargins()
func NewQMargins() *QMargins {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMarginsC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQMarginsFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qmargins.h:55
// index:1
// Public inline
// void QMargins(int, int, int, int)
func NewQMargins_1(left int, top int, right int, bottom int) *QMargins {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMarginsC2Eiiii", ffiqt.FFI_TYPE_VOID, cthis, left, top, right, bottom)
	gopp.ErrPrint(err, rv)
	gothis := NewQMarginsFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qmargins.h:57
// index:0
// Public inline
// bool isNull()
func (this *QMargins) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMargins6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmargins.h:59
// index:0
// Public inline
// int left()
func (this *QMargins) Left() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMargins4leftEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qmargins.h:60
// index:0
// Public inline
// int top()
func (this *QMargins) Top() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMargins3topEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qmargins.h:61
// index:0
// Public inline
// int right()
func (this *QMargins) Right() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMargins5rightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qmargins.h:62
// index:0
// Public inline
// int bottom()
func (this *QMargins) Bottom() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QMargins6bottomEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qmargins.h:64
// index:0
// Public
// void setLeft(int)
func (this *QMargins) SetLeft(left int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMargins7setLeftEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), left)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:65
// index:0
// Public
// void setTop(int)
func (this *QMargins) SetTop(top int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMargins6setTopEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), top)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:66
// index:0
// Public
// void setRight(int)
func (this *QMargins) SetRight(right int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMargins8setRightEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), right)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:67
// index:0
// Public
// void setBottom(int)
func (this *QMargins) SetBottom(bottom int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QMargins9setBottomEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), bottom)
	gopp.ErrPrint(err, rv)
}

//  body block end
