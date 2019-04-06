package qtcore

// /usr/include/qt/QtCore/qfactoryinterface.h
// #include <qfactoryinterface.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QFactoryInterface struct {
	*qtrt.CObject
}
type QFactoryInterface_ITF interface {
	QFactoryInterface_PTR() *QFactoryInterface
}

func (ptr *QFactoryInterface) QFactoryInterface_PTR() *QFactoryInterface { return ptr }

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

/*

 */
func DeleteQFactoryInterface(this *QFactoryInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QFactoryInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qfactoryinterface.h:52
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QStringList keys() const

/*

 */
func (this *QFactoryInterface) Keys() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QFactoryInterface4keysEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

//  body block end

//  keep block begin

func init_unused_10405() {
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
