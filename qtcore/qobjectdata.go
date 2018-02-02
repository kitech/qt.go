package qtcore

// /usr/include/qt/QtCore/qobject.h
// #include <qobject.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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

type QObjectData struct {
	*qtrt.CObject
}

func (this *QObjectData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QObjectData) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQObjectDataFromPointer(cthis unsafe.Pointer) *QObjectData {
	return &QObjectData{&qtrt.CObject{cthis}}
}
func (*QObjectData) NewFromPointer(cthis unsafe.Pointer) *QObjectData {
	return NewQObjectDataFromPointer(cthis)
}

// /usr/include/qt/QtCore/qobject.h:97
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void ~QObjectData()
func DeleteQObjectData(this *QObjectData) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QObjectDataD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qobject.h:112
// index:0
// Public Visibility=Default Availability=Available
// [8] QMetaObject * dynamicMetaObject()
func (this *QObjectData) DynamicMetaObject() *QMetaObject /*777 QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QObjectData17dynamicMetaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
