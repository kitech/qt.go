package qtcore

// /usr/include/qt/QtCore/qobjectdefs.h
// #include <qobjectdefs.h>
// #include <QtCore>

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
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QGenericReturnArgument struct {
	*QGenericArgument
}

func (this *QGenericReturnArgument) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGenericArgument.GetCthis()
	}
}
func (this *QGenericReturnArgument) SetCthis(cthis unsafe.Pointer) {
	this.QGenericArgument = NewQGenericArgumentFromPointer(cthis)
}
func NewQGenericReturnArgumentFromPointer(cthis unsafe.Pointer) *QGenericReturnArgument {
	bcthis0 := NewQGenericArgumentFromPointer(cthis)
	return &QGenericReturnArgument{bcthis0}
}
func (*QGenericReturnArgument) NewFromPointer(cthis unsafe.Pointer) *QGenericReturnArgument {
	return NewQGenericReturnArgumentFromPointer(cthis)
}

// /usr/include/qt/QtCore/qobjectdefs.h:310
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QGenericReturnArgument(const char *, void *)
func NewQGenericReturnArgument(aName string, aData unsafe.Pointer /*666*/) *QGenericReturnArgument {
	var convArg0 = qtrt.CString(aName)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGenericReturnArgumentC2EPKcPv", qtrt.FFI_TYPE_POINTER, convArg0, aData)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGenericReturnArgumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGenericReturnArgument)
	return gothis
}

func DeleteQGenericReturnArgument(this *QGenericReturnArgument) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGenericReturnArgumentD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
		qtrt.KeepMe()
	}
}

//  keep block end
