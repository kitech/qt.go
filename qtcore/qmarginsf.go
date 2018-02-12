package qtcore

// /usr/include/qt/QtCore/qmargins.h
// #include <qmargins.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QMarginsF struct {
	*qtrt.CObject
}

func (this *QMarginsF) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMarginsF) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMarginsFFromPointer(cthis unsafe.Pointer) *QMarginsF {
	return &QMarginsF{&qtrt.CObject{cthis}}
}
func (*QMarginsF) NewFromPointer(cthis unsafe.Pointer) *QMarginsF {
	return NewQMarginsFFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmargins.h:288
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QMarginsF()
func NewQMarginsF() *QMarginsF {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMarginsFC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMarginsFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMarginsF)
	return gothis
}

// /usr/include/qt/QtCore/qmargins.h:289
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QMarginsF(qreal, qreal, qreal, qreal)
func NewQMarginsF_1(left float64, top float64, right float64, bottom float64) *QMarginsF {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMarginsFC2Edddd", qtrt.FFI_TYPE_POINTER, left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMarginsFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMarginsF)
	return gothis
}

// /usr/include/qt/QtCore/qmargins.h:290
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QMarginsF(const QMargins &)
func NewQMarginsF_2(margins *QMargins) *QMarginsF {
	var convArg0 = margins.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMarginsFC2ERK8QMargins", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMarginsFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMarginsF)
	return gothis
}

// /usr/include/qt/QtCore/qmargins.h:292
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QMarginsF) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMarginsF6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmargins.h:294
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal left()
func (this *QMarginsF) Left() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMarginsF4leftEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmargins.h:295
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal top()
func (this *QMarginsF) Top() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMarginsF3topEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmargins.h:296
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal right()
func (this *QMarginsF) Right() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMarginsF5rightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmargins.h:297
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal bottom()
func (this *QMarginsF) Bottom() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMarginsF6bottomEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmargins.h:299
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLeft(qreal)
func (this *QMarginsF) SetLeft(left float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMarginsF7setLeftEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:300
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTop(qreal)
func (this *QMarginsF) SetTop(top float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMarginsF6setTopEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), top)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:301
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRight(qreal)
func (this *QMarginsF) SetRight(right float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMarginsF8setRightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), right)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:302
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBottom(qreal)
func (this *QMarginsF) SetBottom(bottom float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMarginsF9setBottomEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmargins.h:311
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QMargins toMargins()
func (this *QMarginsF) ToMargins() *QMargins /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QMarginsF9toMarginsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMarginsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMargins)
	return rv2
}

func DeleteQMarginsF(this *QMarginsF) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QMarginsFD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

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
}

//  keep block end
