package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
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

type QAccessibleValueInterface struct {
	*qtrt.CObject
}
type QAccessibleValueInterface_ITF interface {
	QAccessibleValueInterface_PTR() *QAccessibleValueInterface
}

func (ptr *QAccessibleValueInterface) QAccessibleValueInterface_PTR() *QAccessibleValueInterface {
	return ptr
}

func (this *QAccessibleValueInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAccessibleValueInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAccessibleValueInterfaceFromPointer(cthis unsafe.Pointer) *QAccessibleValueInterface {
	return &QAccessibleValueInterface{&qtrt.CObject{cthis}}
}
func (*QAccessibleValueInterface) NewFromPointer(cthis unsafe.Pointer) *QAccessibleValueInterface {
	return NewQAccessibleValueInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:566
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleValueInterface()
func DeleteQAccessibleValueInterface(this *QAccessibleValueInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QAccessibleValueInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:568
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QVariant currentValue()
func (this *QAccessibleValueInterface) CurrentValue() *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QAccessibleValueInterface12currentValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:569
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setCurrentValue(const QVariant &)
func (this *QAccessibleValueInterface) SetCurrentValue(value qtcore.QVariant_ITF) {
	var convArg0 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg0 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN25QAccessibleValueInterface15setCurrentValueERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:570
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QVariant maximumValue()
func (this *QAccessibleValueInterface) MaximumValue() *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QAccessibleValueInterface12maximumValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:571
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QVariant minimumValue()
func (this *QAccessibleValueInterface) MinimumValue() *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QAccessibleValueInterface12minimumValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:572
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QVariant minimumStepSize()
func (this *QAccessibleValueInterface) MinimumStepSize() *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QAccessibleValueInterface15minimumStepSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
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
}

//  keep block end
