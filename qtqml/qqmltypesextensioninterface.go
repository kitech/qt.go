package qtqml

// /usr/include/qt/QtQml/qqmlextensioninterface.h
// #include <qqmlextensioninterface.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

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
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

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
	this.CObject = &qtrt.CObject{cthis}
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
func DeleteQQmlTypesExtensionInterface(*QQmlTypesExtensionInterface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QQmlTypesExtensionInterfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlextensioninterface.h:55
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void registerTypes(const char *)
func (this *QQmlTypesExtensionInterface) RegisterTypes(uri string) {
	var convArg0 = qtrt.CString(uri)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QQmlTypesExtensionInterface13registerTypesEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
