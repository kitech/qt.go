package qtcore

// /usr/include/qt/QtCore/qsharedpointer_impl.h
// #include <qsharedpointer_impl.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type ExternalRefCountData struct {
	*qtrt.CObject
}
type ExternalRefCountData_ITF interface {
	ExternalRefCountData_PTR() *ExternalRefCountData
}

func (ptr *ExternalRefCountData) ExternalRefCountData_PTR() *ExternalRefCountData { return ptr }

func (this *ExternalRefCountData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *ExternalRefCountData) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewExternalRefCountDataFromPointer(cthis unsafe.Pointer) *ExternalRefCountData {
	return &ExternalRefCountData{&qtrt.CObject{cthis}}
}
func (*ExternalRefCountData) NewFromPointer(cthis unsafe.Pointer) *ExternalRefCountData {
	return NewExternalRefCountDataFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsharedpointer_impl.h:154
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ExternalRefCountData(Qt::Initialization)

/*

 */
func (*ExternalRefCountData) NewForInherit(arg0 int) *ExternalRefCountData {
	return NewExternalRefCountData(arg0)
}
func NewExternalRefCountData(arg0 int) *ExternalRefCountData {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QtSharedPointer20ExternalRefCountDataC2EN2Qt14InitializationE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewExternalRefCountDataFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteExternalRefCountData)
	return gothis
}

// /usr/include/qt/QtCore/qsharedpointer_impl.h:155
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~ExternalRefCountData()

/*

 */
func DeleteExternalRefCountData(this *ExternalRefCountData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QtSharedPointer20ExternalRefCountDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsharedpointer_impl.h:157
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void destroy()

/*

 */
func (this *ExternalRefCountData) Destroy() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QtSharedPointer20ExternalRefCountData7destroyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedpointer_impl.h:160
// index:0
// Public static Visibility=Default Availability=Available
// [8] QtSharedPointer::ExternalRefCountData * getAndRef(const QObject *)

/*

 */
func (this *ExternalRefCountData) GetAndRef(arg0 QObject_ITF /*777 const QObject **/) unsafe.Pointer /*666*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QtSharedPointer20ExternalRefCountData9getAndRefEPK7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}
func ExternalRefCountData_GetAndRef(arg0 QObject_ITF /*777 const QObject **/) unsafe.Pointer /*666*/ {
	var nilthis *ExternalRefCountData
	rv := nilthis.GetAndRef(arg0)
	return rv
}

// /usr/include/qt/QtCore/qsharedpointer_impl.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setQObjectShared(const QObject *, bool)

/*

 */
func (this *ExternalRefCountData) SetQObjectShared(arg0 QObject_ITF /*777 const QObject **/, enable bool) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QtSharedPointer20ExternalRefCountData16setQObjectSharedEPK7QObjectb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedpointer_impl.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void checkQObjectShared(const QObject *)

/*

 */
func (this *ExternalRefCountData) CheckQObjectShared(arg0 QObject_ITF /*777 const QObject **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QtSharedPointer20ExternalRefCountData18checkQObjectSharedEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedpointer_impl.h:167
// index:0
// Public static inline Visibility=Default Availability=Available
// [-2] void operator delete(void *)

/*

 */
func (this *ExternalRefCountData) Operator_delete(ptr unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QtSharedPointer20ExternalRefCountDatadlEPv", qtrt.FFI_TYPE_POINTER, ptr)
	qtrt.ErrPrint(err, rv)
}
func ExternalRefCountData_Operator_delete(ptr unsafe.Pointer /*666*/) {
	var nilthis *ExternalRefCountData
	nilthis.Operator_delete(ptr)
}

// /usr/include/qt/QtCore/qsharedpointer_impl.h:168
// index:1
// Public static inline Visibility=Default Availability=Available
// [-2] void operator delete(void *, void *)

/*

 */
func (this *ExternalRefCountData) Operator_delete1(arg0 unsafe.Pointer /*666*/, arg1 unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QtSharedPointer20ExternalRefCountDatadlEPvS1_", qtrt.FFI_TYPE_POINTER, arg0, arg1)
	qtrt.ErrPrint(err, rv)
}
func ExternalRefCountData_Operator_delete1(arg0 unsafe.Pointer /*666*/, arg1 unsafe.Pointer /*666*/) {
	var nilthis *ExternalRefCountData
	nilthis.Operator_delete1(arg0, arg1)
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
