package qtqml

// /usr/include/qt/QtQml/qqmlinfo.h
// #include <qqmlinfo.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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

type QQmlInfo struct {
	*qtrt.CObject
}

func (this *QQmlInfo) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlInfo) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQmlInfoFromPointer(cthis unsafe.Pointer) *QQmlInfo {
	return &QQmlInfo{&qtrt.CObject{cthis}}
}
func (*QQmlInfo) NewFromPointer(cthis unsafe.Pointer) *QQmlInfo {
	return NewQQmlInfoFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlinfo.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QQmlInfo()
func DeleteQQmlInfo(this *QQmlInfo) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfoD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
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
