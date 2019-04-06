package qtqml

// /usr/include/qt/QtQml/qqmlprivate.h
// #include <qqmlprivate.h>
// #include <QtQml>

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
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type RegisterSingletonType struct {
	*qtrt.CObject
}
type RegisterSingletonType_ITF interface {
	RegisterSingletonType_PTR() *RegisterSingletonType
}

func (ptr *RegisterSingletonType) RegisterSingletonType_PTR() *RegisterSingletonType { return ptr }

func (this *RegisterSingletonType) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *RegisterSingletonType) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewRegisterSingletonTypeFromPointer(cthis unsafe.Pointer) *RegisterSingletonType {
	return &RegisterSingletonType{&qtrt.CObject{cthis}}
}
func (*RegisterSingletonType) NewFromPointer(cthis unsafe.Pointer) *RegisterSingletonType {
	return NewRegisterSingletonTypeFromPointer(cthis)
}

func DeleteRegisterSingletonType(this *RegisterSingletonType) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21RegisterSingletonTypeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_11489() {
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
