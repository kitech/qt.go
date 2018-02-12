package qtcore

// /usr/include/qt/QtCore/qmap.h
// #include <qmap.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QMapNodeBase struct {
	*qtrt.CObject
}

func (this *QMapNodeBase) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMapNodeBase) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMapNodeBaseFromPointer(cthis unsafe.Pointer) *QMapNodeBase {
	return &QMapNodeBase{&qtrt.CObject{cthis}}
}
func (*QMapNodeBase) NewFromPointer(cthis unsafe.Pointer) *QMapNodeBase {
	return NewQMapNodeBaseFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmap.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] const QMapNodeBase * nextNode()
func (this *QMapNodeBase) NextNode() *QMapNodeBase /*777 const QMapNodeBase **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMapNodeBase8nextNodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMapNodeBaseFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qmap.h:92
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QMapNodeBase * nextNode()
func (this *QMapNodeBase) NextNode_1() *QMapNodeBase /*777 QMapNodeBase **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMapNodeBase8nextNodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMapNodeBaseFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qmap.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] const QMapNodeBase * previousNode()
func (this *QMapNodeBase) PreviousNode() *QMapNodeBase /*777 const QMapNodeBase **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMapNodeBase12previousNodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMapNodeBaseFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qmap.h:94
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QMapNodeBase * previousNode()
func (this *QMapNodeBase) PreviousNode_1() *QMapNodeBase /*777 QMapNodeBase **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMapNodeBase12previousNodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMapNodeBaseFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qmap.h:96
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QMapNodeBase::Color color()
func (this *QMapNodeBase) Color() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMapNodeBase5colorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qmap.h:97
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setColor(enum QMapNodeBase::Color)
func (this *QMapNodeBase) SetColor(c int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMapNodeBase8setColorENS_5ColorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmap.h:98
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QMapNodeBase * parent()
func (this *QMapNodeBase) Parent() *QMapNodeBase /*777 QMapNodeBase **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMapNodeBase6parentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMapNodeBaseFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

func DeleteQMapNodeBase(this *QMapNodeBase) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMapNodeBaseD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QMapNodeBase__Color = int

const QMapNodeBase__Red QMapNodeBase__Color = 0
const QMapNodeBase__Black QMapNodeBase__Color = 1

type QMapNodeBase__ = int

const QMapNodeBase__Mask QMapNodeBase__ = 3

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
		qtrt.KeepMe()
	}
}

//  keep block end
