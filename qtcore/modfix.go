package qtcore

import (
	"unsafe"

	"github.com/kitech/qt.go/qtrt"
)

func (this *QByteArray) Data_fix() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringIN(rv, this.Size())
}

// First parameter need keep alive until object destroyed
func NewQGenericReturnArgument_fix(aName string, aData unsafe.Pointer /*666*/) *QGenericReturnArgument {
	var convArg0 = qtrt.CString(aName)
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGenericReturnArgumentC2EPKcPv", qtrt.FFI_TYPE_POINTER, convArg0, aData)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGenericReturnArgumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, func(obj *QGenericReturnArgument) {
		DeleteQGenericReturnArgument(obj)
		qtrt.FreeMem(convArg0)
	})
	return gothis
}

// First parameter need keep alive until object destroyed
func NewQGenericArgument_fix(aName string, aData unsafe.Pointer /*666*/) *QGenericArgument {
	var convArg0 = qtrt.CString(aName)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGenericArgumentC2EPKcPKv", qtrt.FFI_TYPE_POINTER, convArg0, aData)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGenericArgumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, func(obj *QGenericArgument) {
		DeleteQGenericArgument(obj)
		qtrt.FreeMem(convArg0)
	})
	return gothis
}

func (this *QObject) FindChild(name string) *QObject {
	return this.FindChild_1(name, Qt__FindChildrenRecursively)
}

func (this *QObject) FindChild_1(name string, options int) *QObject {
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("C_QObject_findChild", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, options)
	qtrt.ErrPrint(err, rv)
	return NewQObjectFromPointer(unsafe.Pointer(uintptr(rv)))
}

func (this *QObject) FindChildren(name string) *QObjectListx {
	return this.FindChildren_1(name, Qt__FindChildrenRecursively)
}

func (this *QObject) FindChildren_1(name string, options int) *QObjectListx {
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("C_QObject_findChildren", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, options)
	qtrt.ErrPrint(err, rv)
	return NewQObjectListxFromPointer(unsafe.Pointer(uintptr(rv)))
}

/*
func QVersion() string {
	rv, err := qtrt.InvokeQtFunc6("C_qVersion", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}
*/

type QByteArrayList_ITF interface {
	QByteArrayList_PTR() qtrt.CObjectITF
}
