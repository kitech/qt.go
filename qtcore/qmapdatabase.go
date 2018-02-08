package qtcore

// /usr/include/qt/QtCore/qmap.h
// #include <qmap.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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

type QMapDataBase struct {
	*qtrt.CObject
}

func (this *QMapDataBase) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMapDataBase) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMapDataBaseFromPointer(cthis unsafe.Pointer) *QMapDataBase {
	return &QMapDataBase{&qtrt.CObject{cthis}}
}
func (*QMapDataBase) NewFromPointer(cthis unsafe.Pointer) *QMapDataBase {
	return NewQMapDataBaseFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmap.h:194
// index:0
// Public Visibility=Default Availability=Available
// [-2] void recalcMostLeftNode()
func (this *QMapDataBase) RecalcMostLeftNode() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMapDataBase18recalcMostLeftNodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmap.h:201
// index:0
// Public static Visibility=Default Availability=Available
// [8] QMapDataBase * createData()
func (this *QMapDataBase) CreateData() *QMapDataBase /*777 QMapDataBase **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMapDataBase10createDataEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQMapDataBaseFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QMapDataBase_CreateData() *QMapDataBase /*777 QMapDataBase **/ {
	var nilthis *QMapDataBase
	rv := nilthis.CreateData()
	return rv
}

func DeleteQMapDataBase(this *QMapDataBase) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMapDataBaseD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
