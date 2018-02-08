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
func (this *QRegularExpressionMatchIterator) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQRegularExpressionMatchIteratorFromPointer(cthis unsafe.Pointer) *QRegularExpressionMatchIterator {
	return &QRegularExpressionMatchIterator{&qtrt.CObject{cthis}}
}
func (*QRegularExpressionMatchIterator) NewFromPointer(cthis unsafe.Pointer) *QRegularExpressionMatchIterator {
	return NewQRegularExpressionMatchIteratorFromPointer(cthis)
}

// /usr/include/qt/QtCore/qregularexpression.h:249
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRegularExpressionMatchIterator()
func NewQRegularExpressionMatchIterator() *QRegularExpressionMatchIterator {
	rv, err := qtrt.InvokeQtFunc6("_ZN31QRegularExpressionMatchIteratorC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegularExpressionMatchIteratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegularExpressionMatchIterator)
	return gothis
}

// /usr/include/qt/QtCore/qregularexpression.h:250
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QRegularExpressionMatchIterator()
func DeleteQRegularExpressionMatchIterator(this *QRegularExpressionMatchIterator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN31QRegularExpressionMatchIteratorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qregularexpression.h:257
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QRegularExpressionMatchIterator &)
func (this *QRegularExpressionMatchIterator) Swap(other *QRegularExpressionMatchIterator) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN31QRegularExpressionMatchIterator4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:259
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QRegularExpressionMatchIterator) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregularexpression.h:261
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasNext()
func (this *QRegularExpressionMatchIterator) HasNext() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator7hasNextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregularexpression.h:262
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatch next()
func (this *QRegularExpressionMatchIterator) Next() *QRegularExpressionMatch /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN31QRegularExpressionMatchIterator4nextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatch)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:263
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpressionMatch peekNext()
func (this *QRegularExpressionMatchIterator) PeekNext() *QRegularExpressionMatch /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator8peekNextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionMatchFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpressionMatch)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:265
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpression regularExpression()
func (this *QRegularExpressionMatchIterator) RegularExpression() *QRegularExpression /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator17regularExpressionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpression)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:266
// index:0
// Public Visibility=Default Availability=Available
// [4] QRegularExpression::MatchType matchType()
func (this *QRegularExpressionMatchIterator) MatchType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator9matchTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:267
// index:0
// Public Visibility=Default Availability=Available
// [4] QRegularExpression::MatchOptions matchOptions()
func (this *QRegularExpressionMatchIterator) MatchOptions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK31QRegularExpressionMatchIterator12matchOptionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

//  body block end
