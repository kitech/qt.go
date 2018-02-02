package qtcore

// /usr/include/qt/QtCore/qfactoryinterface.h
// #include <qfactoryinterface.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin

type QFactoryInterface struct {
	*qtrt.CObject
}

func (this *QFactoryInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QFactoryInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQFactoryInterfaceFromPointer(cthis unsafe.Pointer) *QFactoryInterface {
	return &QFactoryInterface{&qtrt.CObject{cthis}}
}
func (*QFactoryInterface) NewFromPointer(cthis unsafe.Pointer) *QFactoryInterface {
	return NewQFactoryInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtCore/qfactoryinterface.h:51
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFactoryInterface()
func DeleteQFactoryInterface(this *QFactoryInterface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QFactoryInterfaceD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
