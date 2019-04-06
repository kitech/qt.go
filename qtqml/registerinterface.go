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
type RegisterInterface struct {
	*qtrt.CObject
}
type RegisterInterface_ITF interface {
	RegisterInterface_PTR() *RegisterInterface
}

func (ptr *RegisterInterface) RegisterInterface_PTR() *RegisterInterface { return ptr }

func (this *RegisterInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *RegisterInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewRegisterInterfaceFromPointer(cthis unsafe.Pointer) *RegisterInterface {
	return &RegisterInterface{&qtrt.CObject{cthis}}
}
func (*RegisterInterface) NewFromPointer(cthis unsafe.Pointer) *RegisterInterface {
	return NewRegisterInterfaceFromPointer(cthis)
}

func DeleteRegisterInterface(this *RegisterInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17RegisterInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_11485() {
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
