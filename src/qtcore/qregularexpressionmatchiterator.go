package qtcore

// /usr/include/qt/QtCore/qregularexpression.h
// #include <qregularexpression.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQRegularExpressionMatchIteratorFromPointer(cthis unsafe.Pointer) *QRegularExpressionMatchIterator {
	return &QRegularExpressionMatchIterator{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qregularexpression.h:249
// index:0
// Public
// void QRegularExpressionMatchIterator()
func NewQRegularExpressionMatchIterator() *QRegularExpressionMatchIterator {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN31QRegularExpressionMatchIteratorC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegularExpressionMatchIteratorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qregularexpression.h:250
// index:0
// Public
// void ~QRegularExpressionMatchIterator()
func DeleteQRegularExpressionMatchIterator(*QRegularExpressionMatchIterator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN31QRegularExpressionMatchIteratorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:257
// index:0
// Public inline
// void swap(class QRegularExpressionMatchIterator &)
func (this *QRegularExpressionMatchIterator) Swap(other *QRegularExpressionMatchIterator) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN31QRegularExpressionMatchIterator4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:259
// index:0
// Public
// bool isValid()
func (this *QRegularExpressionMatchIterator) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qregularexpression.h:261
// index:0
// Public
// bool hasNext()
func (this *QRegularExpressionMatchIterator) HasNext() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator7hasNextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qregularexpression.h:262
// index:0
// Public
// QRegularExpressionMatch next()
func (this *QRegularExpressionMatchIterator) Next() *QRegularExpressionMatch /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN31QRegularExpressionMatchIterator4nextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQRegularExpressionMatchFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:263
// index:0
// Public
// QRegularExpressionMatch peekNext()
func (this *QRegularExpressionMatchIterator) PeekNext() *QRegularExpressionMatch /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator8peekNextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQRegularExpressionMatchFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:265
// index:0
// Public
// QRegularExpression regularExpression()
func (this *QRegularExpressionMatchIterator) RegularExpression() *QRegularExpression /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator17regularExpressionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:266
// index:0
// Public
// QRegularExpression::MatchType matchType()
func (this *QRegularExpressionMatchIterator) MatchType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator9matchTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:267
// index:0
// Public
// QRegularExpression::MatchOptions matchOptions()
func (this *QRegularExpressionMatchIterator) MatchOptions() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator12matchOptionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

//  body block end
