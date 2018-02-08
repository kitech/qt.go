package qtgui

// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QTextLength struct {
	*qtrt.CObject
}

func (this *QTextLength) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextLength) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextLengthFromPointer(cthis unsafe.Pointer) *QTextLength {
	return &QTextLength{&qtrt.CObject{cthis}}
}
func (*QTextLength) NewFromPointer(cthis unsafe.Pointer) *QTextLength {
	return NewQTextLengthFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextformat.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QTextLength()
func NewQTextLength() *QTextLength {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLengthC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextLengthFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextLength)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:91
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QTextLength(enum QTextLength::Type, qreal)
func NewQTextLength_1(type_ int, value float64) *QTextLength {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLengthC2ENS_4TypeEd", qtrt.FFI_TYPE_POINTER, type_, value)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextLengthFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextLength)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QTextLength::Type type()
func (this *QTextLength) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLength4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:94
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal value(qreal)
func (this *QTextLength) Value(maximumLength float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLength5valueEd", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), maximumLength)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:104
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal rawValue()
func (this *QTextLength) RawValue() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextLength8rawValueEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

func DeleteQTextLength(this *QTextLength) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextLengthD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QTextLength__Type = int

const QTextLength__VariableLength QTextLength__Type = 0
const QTextLength__FixedLength QTextLength__Type = 1
const QTextLength__PercentageLength QTextLength__Type = 2

//  body block end
