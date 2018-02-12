package qtqml

// /usr/include/qt/QtQml/qqmlextensioninterface.h
// #include <qqmlextensioninterface.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

type QQmlTypesExtensionInterface struct {
	*qtrt.CObject
}

func (this *QQmlTypesExtensionInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlTypesExtensionInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQmlTypesExtensionInterfaceFromPointer(cthis unsafe.Pointer) *QQmlTypesExtensionInterface {
	return &QQmlTypesExtensionInterface{&qtrt.CObject{cthis}}
}
func (*QQmlTypesExtensionInterface) NewFromPointer(cthis unsafe.Pointer) *QQmlTypesExtensionInterface {
	return NewQQmlTypesExtensionInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlextensioninterface.h:54
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [-2] void ~QQmlTypesExtensionInterface()
func DeleteQQmlTypesExtensionInterface(this *QQmlTypesExtensionInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN27QQmlTypesExtensionInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlextensioninterface.h:55
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void registerTypes(const char *)
func (this *QQmlTypesExtensionInterface) RegisterTypes(uri string) {
	var convArg0 = qtrt.CString(uri)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN27QQmlTypesExtensionInterface13registerTypesEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
