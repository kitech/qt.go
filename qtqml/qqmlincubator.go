package qtqml

// /usr/include/qt/QtQml/qqmlincubator.h
// #include <qqmlincubator.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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

// void statusChanged(QQmlIncubator::Status)
func (this *QQmlIncubator) InheritStatusChanged(f func(arg0 int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "statusChanged", f)
}

// void setInitialState(QObject *)
func (this *QQmlIncubator) InheritSetInitialState(f func(arg0 *qtcore.QObject /*777 QObject **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setInitialState", f)
}

/*

 */
type QQmlIncubator struct {
	*qtrt.CObject
}
type QQmlIncubator_ITF interface {
	QQmlIncubator_PTR() *QQmlIncubator
}

func (ptr *QQmlIncubator) QQmlIncubator_PTR() *QQmlIncubator { return ptr }

func (this *QQmlIncubator) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlIncubator) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQmlIncubatorFromPointer(cthis unsafe.Pointer) *QQmlIncubator {
	return &QQmlIncubator{&qtrt.CObject{cthis}}
}
func (*QQmlIncubator) NewFromPointer(cthis unsafe.Pointer) *QQmlIncubator {
	return NewQQmlIncubatorFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlincubator.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlIncubator(QQmlIncubator::IncubationMode)

/*
Create a new incubator with the specified mode
*/
func NewQQmlIncubator(arg0 int) *QQmlIncubator {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlIncubatorC2ENS_14IncubationModeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlIncubatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlIncubator)
	return gothis
}

// /usr/include/qt/QtQml/qqmlincubator.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlIncubator(QQmlIncubator::IncubationMode)

/*
Create a new incubator with the specified mode
*/
func NewQQmlIncubator__() *QQmlIncubator {
	// arg: 0, QQmlIncubator::IncubationMode=Enum, QQmlIncubator::IncubationMode=Enum, , Invalid
	arg0 := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlIncubatorC2ENS_14IncubationModeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlIncubatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlIncubator)
	return gothis
}

// /usr/include/qt/QtQml/qqmlincubator.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlIncubator()

/*

 */
func DeleteQQmlIncubator(this *QQmlIncubator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlIncubatorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlincubator.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Clears the incubator. Any in-progress incubation is aborted. If the incubator is in the Ready state, the created object is not deleted.
*/
func (this *QQmlIncubator) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlIncubator5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlincubator.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void forceCompletion()

/*
Force any in-progress incubation to finish synchronously. Once this call returns, the incubator will not be in the Loading state.
*/
func (this *QQmlIncubator) ForceCompletion() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlIncubator15forceCompletionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlincubator.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*
Returns true if the incubator's status() is Null.
*/
func (this *QQmlIncubator) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlIncubator6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlincubator.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReady() const

/*
Returns true if the incubator's status() is Ready.
*/
func (this *QQmlIncubator) IsReady() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlIncubator7isReadyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlincubator.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isError() const

/*
Returns true if the incubator's status() is Error.
*/
func (this *QQmlIncubator) IsError() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlIncubator7isErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlincubator.h:77
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLoading() const

/*
Returns true if the incubator's status() is Loading.
*/
func (this *QQmlIncubator) IsLoading() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlIncubator9isLoadingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlincubator.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] QQmlIncubator::IncubationMode incubationMode() const

/*
Return the incubation mode passed to the QQmlIncubator constructor.
*/
func (this *QQmlIncubator) IncubationMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlIncubator14incubationModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQml/qqmlincubator.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] QQmlIncubator::Status status() const

/*
Return the current status of the incubator.
*/
func (this *QQmlIncubator) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlIncubator6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQml/qqmlincubator.h:85
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * object() const

/*
Return the incubated object if the status is Ready, otherwise 0.
*/
func (this *QQmlIncubator) Object() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QQmlIncubator6objectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlincubator.h:88
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void statusChanged(QQmlIncubator::Status)

/*
Called when the status of the incubator changes. status is the new status.

The default implementation does nothing.
*/
func (this *QQmlIncubator) StatusChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlIncubator13statusChangedENS_6StatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlincubator.h:89
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setInitialState(QObject *)

/*
Called after the object is first created, but before property bindings are evaluated and, if applicable, QQmlParserStatus::componentComplete() is called. This is equivalent to the point between QQmlComponent::beginCreate() and QQmlComponent::completeCreate(), and can be used to assign initial values to the object's properties.

The default implementation does nothing.
*/
func (this *QQmlIncubator) SetInitialState(arg0 qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QQmlIncubator15setInitialStateEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
Specifies the mode the incubator operates in. Regardless of the incubation mode, a QQmlIncubator will behave synchronously if the QQmlEngine does not have a QQmlIncubationController set.


*/
type QQmlIncubator__IncubationMode = int

// The object will be created asynchronously.
const QQmlIncubator__Asynchronous QQmlIncubator__IncubationMode = 0

// If the object is being created in a context that is already part of an asynchronous creation, this incubator will join that existing incubation and execute asynchronously. The existing incubation will not become Ready until both it and this incubation have completed. Otherwise, the incubation will execute synchronously.
const QQmlIncubator__AsynchronousIfNested QQmlIncubator__IncubationMode = 1

// The object will be created synchronously.
const QQmlIncubator__Synchronous QQmlIncubator__IncubationMode = 2

/*
Specifies the status of the QQmlIncubator.


*/
type QQmlIncubator__Status = int

// Incubation is not in progress. Call QQmlComponent::create() to begin incubating.
const QQmlIncubator__Null QQmlIncubator__Status = 0

// The object is fully created and can be accessed by calling object().
const QQmlIncubator__Ready QQmlIncubator__Status = 1

// The object is in the process of being created.
const QQmlIncubator__Loading QQmlIncubator__Status = 2

// An error occurred. The errors can be access by calling errors().
const QQmlIncubator__Error QQmlIncubator__Status = 3

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
