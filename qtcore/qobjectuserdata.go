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
// extern C begin: 48
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QObjectUserData struct {
	*qtrt.CObject
}
type QObjectUserData_ITF interface {
	QObjectUserData_PTR() *QObjectUserData
}

func (ptr *QObjectUserData) QObjectUserData_PTR() *QObjectUserData { return ptr }

func (this *QObjectUserData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QObjectUserData) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQObjectUserDataFromPointer(cthis unsafe.Pointer) *QObjectUserData {
	return &QObjectUserData{&qtrt.CObject{cthis}}
}
func (*QObjectUserData) NewFromPointer(cthis unsafe.Pointer) *QObjectUserData {
	return NewQObjectUserDataFromPointer(cthis)
}

// /usr/include/qt/QtCore/qobject.h:479
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QObjectUserData()
func DeleteQObjectUserData(this *QObjectUserData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QObjectUserDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
