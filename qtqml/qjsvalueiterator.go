package qtqml

// /usr/include/qt/QtQml/qjsvalueiterator.h
// #include <qjsvalueiterator.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

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
	if false {
		qtnetwork.KeepMe()
	}
}

//  ext block end

//  body block begin

type QJSValueIterator struct {
	*qtrt.CObject
}

func (this *QJSValueIterator) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QJSValueIterator) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQJSValueIteratorFromPointer(cthis unsafe.Pointer) *QJSValueIterator {
	return &QJSValueIterator{&qtrt.CObject{cthis}}
}
func (*QJSValueIterator) NewFromPointer(cthis unsafe.Pointer) *QJSValueIterator {
	return NewQJSValueIteratorFromPointer(cthis)
}

// /usr/include/qt/QtQml/qjsvalueiterator.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QJSValueIterator(const QJSValue &)
func NewQJSValueIterator(value *QJSValue) *QJSValueIterator {
	var convArg0 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QJSValueIteratorC2ERK8QJSValue", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQJSValueIteratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJSValueIterator)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalueiterator.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QJSValueIterator()
func DeleteQJSValueIterator(this *QJSValueIterator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QJSValueIteratorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qjsvalueiterator.h:59
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasNext()
func (this *QJSValueIterator) HasNext() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QJSValueIterator7hasNextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalueiterator.h:60
// index:0
// Public Visibility=Default Availability=Available
// [1] bool next()
func (this *QJSValueIterator) Next() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QJSValueIterator4nextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalueiterator.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QJSValueIterator) Name() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QJSValueIterator4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtQml/qjsvalueiterator.h:64
// index:0
// Public Visibility=Default Availability=Available
// [8] QJSValue value()
func (this *QJSValueIterator) Value() *QJSValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QJSValueIterator5valueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValue)
	return rv2
}

//  body block end
