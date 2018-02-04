package qtcore

// /usr/include/qt/QtCore/qflags.h
// #include <qflags.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"

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
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin

type QIncompatibleFlag struct {
	*qtrt.CObject
}

func (this *QIncompatibleFlag) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QIncompatibleFlag) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQIncompatibleFlagFromPointer(cthis unsafe.Pointer) *QIncompatibleFlag {
	return &QIncompatibleFlag{&qtrt.CObject{cthis}}
}
func (*QIncompatibleFlag) NewFromPointer(cthis unsafe.Pointer) *QIncompatibleFlag {
	return NewQIncompatibleFlagFromPointer(cthis)
}

// /usr/include/qt/QtCore/qflags.h:80
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QIncompatibleFlag(int)
func NewQIncompatibleFlag(i int) *QIncompatibleFlag {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QIncompatibleFlagC2Ei", qtrt.FFI_TYPE_POINTER, i)
	gopp.ErrPrint(err, rv)
	gothis := NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQIncompatibleFlag)
	return gothis
}

func DeleteQIncompatibleFlag(this *QIncompatibleFlag) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QIncompatibleFlagD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
