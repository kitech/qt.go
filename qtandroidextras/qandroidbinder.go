package qtandroidextras

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidbinder.h
// #include <qandroidbinder.h>
// #include <QtAndroidExtras>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QAndroidBinder struct {
	*qtrt.CObject
}
type QAndroidBinder_ITF interface {
	QAndroidBinder_PTR() *QAndroidBinder
}

func (ptr *QAndroidBinder) QAndroidBinder_PTR() *QAndroidBinder { return ptr }

func (this *QAndroidBinder) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAndroidBinder) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAndroidBinderFromPointer(cthis unsafe.Pointer) *QAndroidBinder {
	return &QAndroidBinder{&qtrt.CObject{cthis}}
}
func (*QAndroidBinder) NewFromPointer(cthis unsafe.Pointer) *QAndroidBinder {
	return NewQAndroidBinderFromPointer(cthis)
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidbinder.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAndroidBinder()
func NewQAndroidBinder() *QAndroidBinder {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidBinderC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAndroidBinderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidBinder)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidbinder.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAndroidBinder(const QAndroidJniObject &)
func NewQAndroidBinder_1(binder QAndroidJniObject_ITF) *QAndroidBinder {
	var convArg0 unsafe.Pointer
	if binder != nil && binder.QAndroidJniObject_PTR() != nil {
		convArg0 = binder.QAndroidJniObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidBinderC2ERK17QAndroidJniObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAndroidBinderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidBinder)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidbinder.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAndroidBinder()
func DeleteQAndroidBinder(this *QAndroidBinder) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidBinderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 1)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidbinder.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool onTransact(int, const QAndroidParcel &, const QAndroidParcel &, enum QAndroidBinder::CallType)
func (this *QAndroidBinder) OnTransact(code int, data QAndroidParcel_ITF, reply QAndroidParcel_ITF, flags int) bool {
	var convArg1 unsafe.Pointer
	if data != nil && data.QAndroidParcel_PTR() != nil {
		convArg1 = data.QAndroidParcel_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if reply != nil && reply.QAndroidParcel_PTR() != nil {
		convArg2 = reply.QAndroidParcel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidBinder10onTransactEiRK14QAndroidParcelS2_NS_8CallTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), code, convArg1, convArg2, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidbinder.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool transact(int, const QAndroidParcel &, QAndroidParcel *, enum QAndroidBinder::CallType) const
func (this *QAndroidBinder) Transact(code int, data QAndroidParcel_ITF, reply QAndroidParcel_ITF /*777 QAndroidParcel **/, flags int) bool {
	var convArg1 unsafe.Pointer
	if data != nil && data.QAndroidParcel_PTR() != nil {
		convArg1 = data.QAndroidParcel_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if reply != nil && reply.QAndroidParcel_PTR() != nil {
		convArg2 = reply.QAndroidParcel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidBinder8transactEiRK14QAndroidParcelPS0_NS_8CallTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), code, convArg1, convArg2, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidbinder.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool transact(int, const QAndroidParcel &, QAndroidParcel *, enum QAndroidBinder::CallType) const
func (this *QAndroidBinder) Transact__(code int, data QAndroidParcel_ITF) bool {
	var convArg1 unsafe.Pointer
	if data != nil && data.QAndroidParcel_PTR() != nil {
		convArg1 = data.QAndroidParcel_PTR().GetCthis()
	}
	// arg: 2, QAndroidParcel *=Pointer, QAndroidParcel=Record,
	var convArg2 unsafe.Pointer
	// arg: 3, QAndroidBinder::CallType=Enum, QAndroidBinder::CallType=Enum,
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidBinder8transactEiRK14QAndroidParcelPS0_NS_8CallTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), code, convArg1, convArg2, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidbinder.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool transact(int, const QAndroidParcel &, QAndroidParcel *, enum QAndroidBinder::CallType) const
func (this *QAndroidBinder) Transact__1(code int, data QAndroidParcel_ITF, reply QAndroidParcel_ITF /*777 QAndroidParcel **/) bool {
	var convArg1 unsafe.Pointer
	if data != nil && data.QAndroidParcel_PTR() != nil {
		convArg1 = data.QAndroidParcel_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if reply != nil && reply.QAndroidParcel_PTR() != nil {
		convArg2 = reply.QAndroidParcel_PTR().GetCthis()
	}
	// arg: 3, QAndroidBinder::CallType=Enum, QAndroidBinder::CallType=Enum,
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidBinder8transactEiRK14QAndroidParcelPS0_NS_8CallTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), code, convArg1, convArg2, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidbinder.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] QAndroidJniObject handle() const
func (this *QAndroidBinder) Handle() *QAndroidJniObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidBinder6handleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAndroidJniObject)
	return rv2
}

type QAndroidBinder__CallType = int

const QAndroidBinder__Normal QAndroidBinder__CallType = 0
const QAndroidBinder__OneWay QAndroidBinder__CallType = 1

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
}

//  keep block end
