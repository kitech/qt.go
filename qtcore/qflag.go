package qtcore

// /usr/include/qt/QtCore/qflags.h
// #include <qflags.h>
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
type QFlag struct {
	*qtrt.CObject
}
type QFlag_ITF interface {
	QFlag_PTR() *QFlag
}

func (ptr *QFlag) QFlag_PTR() *QFlag { return ptr }

func (this *QFlag) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QFlag) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQFlagFromPointer(cthis unsafe.Pointer) *QFlag {
	return &QFlag{&qtrt.CObject{cthis}}
}
func (*QFlag) NewFromPointer(cthis unsafe.Pointer) *QFlag {
	return NewQFlagFromPointer(cthis)
}

// /usr/include/qt/QtCore/qflags.h:57
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QFlag(int)

/*

 */
func (*QFlag) NewForInherit(value int) *QFlag {
	return NewQFlag(value)
}
func NewQFlag(value int) *QFlag {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFlagC2Ei", qtrt.FFI_TYPE_POINTER, value)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFlagFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFlag)
	return gothis
}

// /usr/include/qt/QtCore/qflags.h:68
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QFlag(uint)

/*

 */
func (*QFlag) NewForInherit1(value uint) *QFlag {
	return NewQFlag1(value)
}
func NewQFlag1(value uint) *QFlag {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFlagC2Ej", qtrt.FFI_TYPE_POINTER, value)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFlagFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFlag)
	return gothis
}

// /usr/include/qt/QtCore/qflags.h:69
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QFlag(short)

/*

 */
func (*QFlag) NewForInherit2(value int16) *QFlag {
	return NewQFlag2(value)
}
func NewQFlag2(value int16) *QFlag {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFlagC2Es", qtrt.FFI_TYPE_POINTER, value)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFlagFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFlag)
	return gothis
}

// /usr/include/qt/QtCore/qflags.h:70
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QFlag(ushort)

/*

 */
func (*QFlag) NewForInherit3(value uint16) *QFlag {
	return NewQFlag3(value)
}
func NewQFlag3(value uint16) *QFlag {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFlagC2Et", qtrt.FFI_TYPE_POINTER, value)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFlagFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFlag)
	return gothis
}

func DeleteQFlag(this *QFlag) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFlagD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
