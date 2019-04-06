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
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QGenericReturnArgument struct {
	*QGenericArgument
}
type QGenericReturnArgument_ITF interface {
	QGenericArgument_ITF
	QGenericReturnArgument_PTR() *QGenericReturnArgument
}

func (ptr *QGenericReturnArgument) QGenericReturnArgument_PTR() *QGenericReturnArgument { return ptr }

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

// /usr/include/qt/QtCore/qobjectdefs.h:305
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QGenericReturnArgument(const char *, void *)

/*

 */
func (*QGenericReturnArgument) NewForInherit(aName string, aData unsafe.Pointer /*666*/) *QGenericReturnArgument {
	return NewQGenericReturnArgument(aName, aData)
}
func NewQGenericReturnArgument(aName string, aData unsafe.Pointer /*666*/) *QGenericReturnArgument {
	var convArg0 = qtrt.CString(aName)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGenericReturnArgumentC2EPKcPv", qtrt.FFI_TYPE_POINTER, convArg0, aData)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGenericReturnArgumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGenericReturnArgument)
	return gothis
}

// /usr/include/qt/QtCore/qobjectdefs.h:305
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QGenericReturnArgument(const char *, void *)

/*

 */
func (*QGenericReturnArgument) NewForInheritp() *QGenericReturnArgument {
	return NewQGenericReturnArgumentp()
}
func NewQGenericReturnArgumentp() *QGenericReturnArgument {
	// arg: 0, const char *=Pointer, =Invalid, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, void *=Pointer, =Invalid, , Invalid
	var aData unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGenericReturnArgumentC2EPKcPv", qtrt.FFI_TYPE_POINTER, convArg0, aData)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGenericReturnArgumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGenericReturnArgument)
	return gothis
}

// /usr/include/qt/QtCore/qobjectdefs.h:305
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QGenericReturnArgument(const char *, void *)

/*

 */
func (*QGenericReturnArgument) NewForInheritp1(aName string) *QGenericReturnArgument {
	return NewQGenericReturnArgumentp1(aName)
}
func NewQGenericReturnArgumentp1(aName string) *QGenericReturnArgument {
	var convArg0 = qtrt.CString(aName)
	defer qtrt.FreeMem(convArg0)
	// arg: 1, void *=Pointer, =Invalid, , Invalid
	var aData unsafe.Pointer
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

func init_unused_10157() {
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
