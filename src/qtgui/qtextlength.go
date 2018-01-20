//  header block begin
// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	return this.Cthis
}

// /usr/include/qt/QtGui/qtextformat.h:89
// index:0
// inline
// void QTextLength()
func NewQTextLength() *QTextLength {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLengthC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextLengthFromPointer(cthis)
	return gothis
}
func NewQTextLengthFromPointer(cthis unsafe.Pointer) *QTextLength {
	return &QTextLength{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qtextformat.h:91
// index:1
// inline
// void QTextLength(enum QTextLength::Type, qreal)
func NewQTextLength_1(type_ int, value float64) *QTextLength {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextLengthC2ENS_4TypeEd", ffiqt.FFI_TYPE_VOID, cthis, &type_, &value)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextLengthFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:93
// index:0
// inline
// QTextLength::Type type()
func (this *QTextLength) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLength4typeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:94
// index:0
// inline
// qreal value(qreal)
func (this *QTextLength) Value(maximumLength float64) {
	// 0: (, maximumLength qreal), (&maximumLength)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLength5valueEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &maximumLength)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:104
// index:0
// inline
// qreal rawValue()
func (this *QTextLength) RawValue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextLength8rawValueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end