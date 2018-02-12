package qtcore

// /usr/include/qt/QtCore/qhash.h
// #include <qhash.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QHashDummyValue struct {
	*qtrt.CObject
}

func (this *QHashDummyValue) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QHashDummyValue) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQHashDummyValueFromPointer(cthis unsafe.Pointer) *QHashDummyValue {
	return &QHashDummyValue{&qtrt.CObject{cthis}}
}
func (*QHashDummyValue) NewFromPointer(cthis unsafe.Pointer) *QHashDummyValue {
	return NewQHashDummyValueFromPointer(cthis)
}

func DeleteQHashDummyValue(this *QHashDummyValue) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QHashDummyValueD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
}

//  keep block end
