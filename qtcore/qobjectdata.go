package qtcore

// /usr/include/qt/QtCore/qobject.h
// #include <qobject.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
type QObjectData struct {
	*qtrt.CObject
}
type QObjectData_ITF interface {
	QObjectData_PTR() *QObjectData
}

func (ptr *QObjectData) QObjectData_PTR() *QObjectData { return ptr }

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

/*

 */
func DeleteQObjectData(this *QObjectData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QObjectDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qobject.h:113
// index:0
// Public Visibility=Default Availability=Available
// [8] QMetaObject * dynamicMetaObject() const

/*

 */
func (this *QObjectData) DynamicMetaObject() *QMetaObject /*777 QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QObjectData17dynamicMetaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end

//  keep block begin

func init_unused_10249() {
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
