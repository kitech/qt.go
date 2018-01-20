//  header block begin
// /usr/include/qt/QtCore/qregularexpression.h
// #include <qregularexpression.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 28
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
type QRegularExpressionMatchIterator struct {
	*qtrt.CObject
}

func (this *QRegularExpressionMatchIterator) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qregularexpression.h:249
// index:0
// void QRegularExpressionMatchIterator()
func NewQRegularExpressionMatchIterator() *QRegularExpressionMatchIterator {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN31QRegularExpressionMatchIteratorC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegularExpressionMatchIteratorFromPointer(cthis)
	return gothis
}
func NewQRegularExpressionMatchIteratorFromPointer(cthis unsafe.Pointer) *QRegularExpressionMatchIterator {
	return &QRegularExpressionMatchIterator{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qregularexpression.h:250
// index:0
// void ~QRegularExpressionMatchIterator()
func DeleteQRegularExpressionMatchIterator(*QRegularExpressionMatchIterator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN31QRegularExpressionMatchIteratorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:257
// index:0
// inline
// void swap(class QRegularExpressionMatchIterator &)
func (this *QRegularExpressionMatchIterator) Swap(other unsafe.Pointer) {
	// 0: (, other QRegularExpressionMatchIterator &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN31QRegularExpressionMatchIterator4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:259
// index:0
// bool isValid()
func (this *QRegularExpressionMatchIterator) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator7isValidEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:261
// index:0
// bool hasNext()
func (this *QRegularExpressionMatchIterator) HasNext() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator7hasNextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:262
// index:0
// QRegularExpressionMatch next()
func (this *QRegularExpressionMatchIterator) Next() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN31QRegularExpressionMatchIterator4nextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:263
// index:0
// QRegularExpressionMatch peekNext()
func (this *QRegularExpressionMatchIterator) PeekNext() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator8peekNextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:265
// index:0
// QRegularExpression regularExpression()
func (this *QRegularExpressionMatchIterator) RegularExpression() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator17regularExpressionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:266
// index:0
// QRegularExpression::MatchType matchType()
func (this *QRegularExpressionMatchIterator) MatchType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator9matchTypeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:267
// index:0
// QRegularExpression::MatchOptions matchOptions()
func (this *QRegularExpressionMatchIterator) MatchOptions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator12matchOptionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
