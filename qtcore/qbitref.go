package qtcore

// /usr/include/qt/QtCore/qbitarray.h
// #include <qbitarray.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 33
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QBitRef struct {
	*qtrt.CObject
}
type QBitRef_ITF interface {
	QBitRef_PTR() *QBitRef
}

func (ptr *QBitRef) QBitRef_PTR() *QBitRef { return ptr }

func (this *QBitRef) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QBitRef) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQBitRefFromPointer(cthis unsafe.Pointer) *QBitRef {
	return &QBitRef{&qtrt.CObject{cthis}}
}
func (*QBitRef) NewFromPointer(cthis unsafe.Pointer) *QBitRef {
	return NewQBitRefFromPointer(cthis)
}

// /usr/include/qt/QtCore/qbitarray.h:152
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!()
func (this *QBitRef) Operator_not() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QBitRefntEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbitarray.h:153
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QBitRef & operator=(const QBitRef &)
func (this *QBitRef) Operator_equal(val QBitRef_ITF) *QBitRef {
	var convArg0 unsafe.Pointer
	if val != nil && val.QBitRef_PTR() != nil {
		convArg0 = val.QBitRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBitRefaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitRef)
	return rv2
}

// /usr/include/qt/QtCore/qbitarray.h:154
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QBitRef & operator=(_Bool)
func (this *QBitRef) Operator_equal_1(val bool) *QBitRef {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBitRefaSEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), val)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBitRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBitRef)
	return rv2
}

func DeleteQBitRef(this *QBitRef) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBitRefD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
