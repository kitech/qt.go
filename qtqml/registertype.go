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
// extern C begin: 7
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
type RegisterType struct {
	*qtrt.CObject
}
type RegisterType_ITF interface {
	RegisterType_PTR() *RegisterType
}

func (ptr *RegisterType) RegisterType_PTR() *RegisterType { return ptr }

func (this *RegisterType) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *RegisterType) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewRegisterTypeFromPointer(cthis unsafe.Pointer) *RegisterType {
	return &RegisterType{&qtrt.CObject{cthis}}
}
func (*RegisterType) NewFromPointer(cthis unsafe.Pointer) *RegisterType {
	return NewRegisterTypeFromPointer(cthis)
}

func DeleteRegisterType(this *RegisterType) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12RegisterTypeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_11483() {
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
