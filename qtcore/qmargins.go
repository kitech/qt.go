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
import "gopp"
import "qt.go/qtrt"

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
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMarginsFromPointer(cthis unsafe.Pointer) *QMargins {
	return &QMargins{&qtrt.CObject{cthis}}
}
func (*QMargins) NewFromPointer(cthis unsafe.Pointer) *QMargins {
	return NewQMarginsFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmargins.h:54
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QMargins()
func NewQMargins() *QMargins {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMarginsC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQMarginsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMargins)
	return gothis
}

// /usr/include/qt/QtCore/qmargins.h:55
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QMargins(int, int, int, int)
func NewQMargins_1(left int, top int, right int, bottom int) *QMargins {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMarginsC2Eiiii", qtrt.FFI_TYPE_POINTER, left, top, right, bottom)
	gopp.ErrPrint(err, rv)
	gothis := NewQMarginsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMargins)
	return gothis
}

// /usr/include/qt/QtCore/qmargins.h:57
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QMargins) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMargins6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmargins.h:59
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int left()
func (this *QMargins) Left() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMargins4leftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmargins.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int top()
func (this *QMargins) Top() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMargins3topEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmargins.h:61
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int right()
func (this *QMargins) Right() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMargins5rightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmargins.h:62
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int bottom()
func (this *QMargins) Bottom() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QMargins6bottomEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmargins.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLeft(int)
func (this *QMargins) SetLeft(left int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMargins7setLeftEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTop(int)
func (this *QMargins) SetTop(top int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMargins6setTopEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), top)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRight(int)
func (this *QMargins) SetRight(right int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMargins8setRightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), right)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBottom(int)
func (this *QMargins) SetBottom(bottom int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMargins9setBottomEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bottom)
	gopp.ErrPrint(err, rv)
}

func DeleteQMargins(this *QMargins) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QMarginsD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
