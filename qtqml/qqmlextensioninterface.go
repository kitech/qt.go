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
// extern C begin: 2
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
type QQmlExtensionInterface struct {
	*QQmlTypesExtensionInterface
}
type QQmlExtensionInterface_ITF interface {
	QQmlTypesExtensionInterface_ITF
	QQmlExtensionInterface_PTR() *QQmlExtensionInterface
}

func (ptr *QQmlExtensionInterface) QQmlExtensionInterface_PTR() *QQmlExtensionInterface { return ptr }

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
// Public inline virtual Visibility=Default Availability=Available
// [-2] void ~QQmlExtensionInterface()

/*

 */
func DeleteQQmlExtensionInterface(this *QQmlExtensionInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQmlExtensionInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlextensioninterface.h:62
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void initializeEngine(QQmlEngine *, const char *)

/*

 */
func (this *QQmlExtensionInterface) InitializeEngine(engine QQmlEngine_ITF /*777 QQmlEngine **/, uri string) {
	var convArg0 unsafe.Pointer
	if engine != nil && engine.QQmlEngine_PTR() != nil {
		convArg0 = engine.QQmlEngine_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(uri)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN22QQmlExtensionInterface16initializeEngineEP10QQmlEnginePKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11525() {
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
