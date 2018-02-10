package qtcore

// /usr/include/qt/QtCore/qobject.h
// #include <qobject.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 48
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
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
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin

type QObjectUserData struct {
	*qtrt.CObject
}

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

// /usr/include/qt/QtCore/qobject.h:478
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QObjectUserData()
func DeleteQObjectUserData(this *QObjectUserData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QObjectUserDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
