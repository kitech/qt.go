package qtandroidextras

// /usr/include/qt/QtAndroidExtras/qandroidbinder.h
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

/*

 */
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

// /usr/include/qt/QtAndroidExtras/qandroidbinder.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAndroidBinder()

/*
Creates a new object which can be used to perform IPC.

See also onTransact and transact.
*/
func NewQAndroidBinder() *QAndroidBinder {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidBinderC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAndroidBinderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidBinder)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/qandroidbinder.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAndroidBinder(const QAndroidJniObject &)

/*
Creates a new object which can be used to perform IPC.

See also onTransact and transact.
*/
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

// /usr/include/qt/QtAndroidExtras/qandroidbinder.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAndroidBinder()

/*

 */
func DeleteQAndroidBinder(this *QAndroidBinder) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidBinderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtAndroidExtras/qandroidbinder.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool onTransact(int, const QAndroidParcel &, const QAndroidParcel &, QAndroidBinder::CallType)

/*
Default implementation is a stub that returns false. The user should override this method to get the transact data from the caller.

The code is the action to perform. The data is the marshaled data sent by the caller.
 The reply is the marshaled data to be sent to the caller.
 The flags are the additional operation flags.


Warning: This method is called from Binder's thread which is different from the thread that this object was created.

See also transact.
*/
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

// /usr/include/qt/QtAndroidExtras/qandroidbinder.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool transact(int, const QAndroidParcel &, QAndroidParcel *, QAndroidBinder::CallType) const

/*
Performs an IPC call

The code is the action to perform. Should be between FIRST_CALL_TRANSACTION and LAST_CALL_TRANSACTION.
 The data is the marshaled data to send to the target.
 The reply (if specified) is the marshaled data to be received from the target. May be nullptr if you are not interested in the return value.
 The flags are the additional operation flags.


Returns true on success
*/
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

// /usr/include/qt/QtAndroidExtras/qandroidbinder.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool transact(int, const QAndroidParcel &, QAndroidParcel *, QAndroidBinder::CallType) const

/*
Performs an IPC call

The code is the action to perform. Should be between FIRST_CALL_TRANSACTION and LAST_CALL_TRANSACTION.
 The data is the marshaled data to send to the target.
 The reply (if specified) is the marshaled data to be received from the target. May be nullptr if you are not interested in the return value.
 The flags are the additional operation flags.


Returns true on success
*/
func (this *QAndroidBinder) Transact__(code int, data QAndroidParcel_ITF) bool {
	var convArg1 unsafe.Pointer
	if data != nil && data.QAndroidParcel_PTR() != nil {
		convArg1 = data.QAndroidParcel_PTR().GetCthis()
	}
	// arg: 2, QAndroidParcel *=Pointer, QAndroidParcel=Record, , Invalid
	var convArg2 unsafe.Pointer
	// arg: 3, QAndroidBinder::CallType=Enum, QAndroidBinder::CallType=Enum, , Invalid
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidBinder8transactEiRK14QAndroidParcelPS0_NS_8CallTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), code, convArg1, convArg2, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtAndroidExtras/qandroidbinder.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool transact(int, const QAndroidParcel &, QAndroidParcel *, QAndroidBinder::CallType) const

/*
Performs an IPC call

The code is the action to perform. Should be between FIRST_CALL_TRANSACTION and LAST_CALL_TRANSACTION.
 The data is the marshaled data to send to the target.
 The reply (if specified) is the marshaled data to be received from the target. May be nullptr if you are not interested in the return value.
 The flags are the additional operation flags.


Returns true on success
*/
func (this *QAndroidBinder) Transact__1(code int, data QAndroidParcel_ITF, reply QAndroidParcel_ITF /*777 QAndroidParcel **/) bool {
	var convArg1 unsafe.Pointer
	if data != nil && data.QAndroidParcel_PTR() != nil {
		convArg1 = data.QAndroidParcel_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if reply != nil && reply.QAndroidParcel_PTR() != nil {
		convArg2 = reply.QAndroidParcel_PTR().GetCthis()
	}
	// arg: 3, QAndroidBinder::CallType=Enum, QAndroidBinder::CallType=Enum, , Invalid
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidBinder8transactEiRK14QAndroidParcelPS0_NS_8CallTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), code, convArg1, convArg2, flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtAndroidExtras/qandroidbinder.h:70
// index:0
// Public Visibility=Default Availability=Available
// [16] QAndroidJniObject handle() const

/*
The return value is useful to call other Java API which are not covered by this wrapper
*/
func (this *QAndroidBinder) Handle() *QAndroidJniObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidBinder6handleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAndroidJniObject)
	return rv2
}

/*
This enum is used with QAndroidBinder::transact() to describe the mode in which the IPC call is performed.


*/
type QAndroidBinder__CallType = int

// normal IPC, meaning that the caller waits the result from the callee
const QAndroidBinder__Normal QAndroidBinder__CallType = 0

// one-way IPC, meaning that the caller returns immediately, without waiting for a result from the callee
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
