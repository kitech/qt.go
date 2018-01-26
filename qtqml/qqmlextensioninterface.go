package qtqml

// /usr/include/qt/QtQml/qqmlextensioninterface.h
// #include <qqmlextensioninterface.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
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
type QQmlExtensionInterface struct {
	*QQmlTypesExtensionInterface
}

func (this *QQmlExtensionInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QQmlTypesExtensionInterface.GetCthis()
	}
}
func (this *QQmlExtensionInterface) SetCthis(cthis unsafe.Pointer) {
	this.QQmlTypesExtensionInterface = NewQQmlTypesExtensionInterfaceFromPointer(cthis)
}
func NewQQmlExtensionInterfaceFromPointer(cthis unsafe.Pointer) *QQmlExtensionInterface {
	bcthis0 := NewQQmlTypesExtensionInterfaceFromPointer(cthis)
	return &QQmlExtensionInterface{bcthis0}
}
func (*QQmlExtensionInterface) NewFromPointer(cthis unsafe.Pointer) *QQmlExtensionInterface {
	return NewQQmlExtensionInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlextensioninterface.h:61
// index:0
// Public inline virtual
// void ~QQmlExtensionInterface()
func DeleteQQmlExtensionInterface(*QQmlExtensionInterface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QQmlExtensionInterfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlextensioninterface.h:62
// index:0
// Public pure virtual
// void initializeEngine(class QQmlEngine *, const char *)
func (this *QQmlExtensionInterface) InitializeEngine(engine *QQmlEngine /*777 QQmlEngine **/, uri string) {
	var convArg0 = engine.GetCthis()
	var convArg1 = qtrt.CString(uri)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QQmlExtensionInterface16initializeEngineEP10QQmlEnginePKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

//  body block end
