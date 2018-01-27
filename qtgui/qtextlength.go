package qtgui

// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
		ffiqt.KeepMe()
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
	this.CObject = &qtrt.CObject{cthis}
}
func NewQTextLengthFromPointer(cthis unsafe.Pointer) *QTextLength {
	return &QTextLength{&qtrt.CObject{cthis}}
}
func (*QTextLength) NewFromPointer(cthis unsafe.Pointer) *QTextLength {
	return NewQTextLengthFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextformat.h:89
// index:0
// Public inline
// void QTextLength()
func NewQTextLength() *QTextLength {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLengthC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextLengthFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:91
// index:1
// Public inline
// void QTextLength(QTextLength::Type, qreal)
func NewQTextLength_1(type_ int, value float64) *QTextLength {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLengthC2ENS_4TypeEd", ffiqt.FFI_TYPE_VOID, cthis, type_, value)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextLengthFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:93
// index:0
// Public inline
// QTextLength::Type type()
func (this *QTextLength) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLength4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:94
// index:0
// Public inline
// qreal value(qreal)
func (this *QTextLength) Value(maximumLength float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLength5valueEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), maximumLength)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextformat.h:104
// index:0
// Public inline
// qreal rawValue()
func (this *QTextLength) RawValue() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLength8rawValueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

type QTextLength__Type = int

const QTextLength__VariableLength QTextLength__Type = 0
const QTextLength__FixedLength QTextLength__Type = 1
const QTextLength__PercentageLength QTextLength__Type = 2

//  body block end
