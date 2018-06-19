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
// extern C begin: 1
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
type QGenericArgument struct {
	*qtrt.CObject
}
type QGenericArgument_ITF interface {
	QGenericArgument_PTR() *QGenericArgument
}

func (ptr *QGenericArgument) QGenericArgument_PTR() *QGenericArgument { return ptr }

func (this *QGenericArgument) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGenericArgument) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGenericArgumentFromPointer(cthis unsafe.Pointer) *QGenericArgument {
	return &QGenericArgument{&qtrt.CObject{cthis}}
}
func (*QGenericArgument) NewFromPointer(cthis unsafe.Pointer) *QGenericArgument {
	return NewQGenericArgumentFromPointer(cthis)
}

// /usr/include/qt/QtCore/qobjectdefs.h:297
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QGenericArgument(const char *, const void *)

/*

 */
func NewQGenericArgument(aName string, aData unsafe.Pointer /*666*/) *QGenericArgument {
	var convArg0 = qtrt.CString(aName)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGenericArgumentC2EPKcPKv", qtrt.FFI_TYPE_POINTER, convArg0, aData)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGenericArgumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGenericArgument)
	return gothis
}

// /usr/include/qt/QtCore/qobjectdefs.h:297
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QGenericArgument(const char *, const void *)

/*

 */
func NewQGenericArgument__() *QGenericArgument {
	// arg: 0, const char *=Pointer, =Invalid, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, const void *=Pointer, =Invalid, , Invalid
	var aData unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGenericArgumentC2EPKcPKv", qtrt.FFI_TYPE_POINTER, convArg0, aData)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGenericArgumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGenericArgument)
	return gothis
}

// /usr/include/qt/QtCore/qobjectdefs.h:297
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QGenericArgument(const char *, const void *)

/*

 */
func NewQGenericArgument__1(aName string) *QGenericArgument {
	var convArg0 = qtrt.CString(aName)
	defer qtrt.FreeMem(convArg0)
	// arg: 1, const void *=Pointer, =Invalid, , Invalid
	var aData unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGenericArgumentC2EPKcPKv", qtrt.FFI_TYPE_POINTER, convArg0, aData)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGenericArgumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGenericArgument)
	return gothis
}

// /usr/include/qt/QtCore/qobjectdefs.h:299
// index:0
// Public inline Visibility=Default Availability=Available
// [8] void * data() const

/*

 */
func (this *QGenericArgument) Data() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QGenericArgument4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qobjectdefs.h:300
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const char * name() const

/*

 */
func (this *QGenericArgument) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QGenericArgument4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

func DeleteQGenericArgument(this *QGenericArgument) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QGenericArgumentD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
