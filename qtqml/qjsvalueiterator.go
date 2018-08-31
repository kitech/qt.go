package qtqml

// /usr/include/qt/QtQml/qjsvalueiterator.h
// #include <qjsvalueiterator.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QJSValueIterator struct {
	*qtrt.CObject
}
type QJSValueIterator_ITF interface {
	QJSValueIterator_PTR() *QJSValueIterator
}

func (ptr *QJSValueIterator) QJSValueIterator_PTR() *QJSValueIterator { return ptr }

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

/*
Constructs an iterator for traversing object. The iterator is set to be at the front of the sequence of properties (before the first property).
*/
func (*QJSValueIterator) NewForInherit(value QJSValue_ITF) *QJSValueIterator {
	return NewQJSValueIterator(value)
}
func NewQJSValueIterator(value QJSValue_ITF) *QJSValueIterator {
	var convArg0 unsafe.Pointer
	if value != nil && value.QJSValue_PTR() != nil {
		convArg0 = value.QJSValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QJSValueIteratorC2ERK8QJSValue", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQJSValueIteratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQJSValueIterator)
	return gothis
}

// /usr/include/qt/QtQml/qjsvalueiterator.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QJSValueIterator()

/*

 */
func DeleteQJSValueIterator(this *QJSValueIterator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QJSValueIteratorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qjsvalueiterator.h:59
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasNext() const

/*
Returns true if there is at least one item ahead of the iterator (i.e. the iterator is not at the back of the property sequence); otherwise returns false.

See also next().
*/
func (this *QJSValueIterator) HasNext() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QJSValueIterator7hasNextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalueiterator.h:60
// index:0
// Public Visibility=Default Availability=Available
// [1] bool next()

/*
Advances the iterator by one position. Returns true if there was at least one item ahead of the iterator (i.e. the iterator was not already at the back of the property sequence); otherwise returns false.

See also hasNext() and name().
*/
func (this *QJSValueIterator) Next() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QJSValueIterator4nextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qjsvalueiterator.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name() const

/*
Returns the name of the last property that was jumped over using next().

See also value().
*/
func (this *QJSValueIterator) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QJSValueIterator4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtQml/qjsvalueiterator.h:64
// index:0
// Public Visibility=Default Availability=Available
// [8] QJSValue value() const

/*
Returns the value of the last property that was jumped over using next().

See also name().
*/
func (this *QJSValueIterator) Value() *QJSValue /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QJSValueIterator5valueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValue)
	return rv2
}

// /usr/include/qt/QtQml/qjsvalueiterator.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QJSValueIterator & operator=(QJSValue &)

/*

 */
func (this *QJSValueIterator) Operator_equal(value QJSValue_ITF) *QJSValueIterator {
	var convArg0 unsafe.Pointer
	if value != nil && value.QJSValue_PTR() != nil {
		convArg0 = value.QJSValue_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QJSValueIteratoraSER8QJSValue", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQJSValueIteratorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQJSValueIterator)
	return rv2
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
