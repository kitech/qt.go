package qtcore

// /usr/include/qt/QtCore/qhashfunctions.h
// #include <qhashfunctions.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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

type QHashCombineCommutative struct {
	*qtrt.CObject
}

func (this *QHashCombineCommutative) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QHashCombineCommutative) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQHashCombineCommutativeFromPointer(cthis unsafe.Pointer) *QHashCombineCommutative {
	return &QHashCombineCommutative{&qtrt.CObject{cthis}}
}
func (*QHashCombineCommutative) NewFromPointer(cthis unsafe.Pointer) *QHashCombineCommutative {
	return NewQHashCombineCommutativeFromPointer(cthis)
}

func DeleteQHashCombineCommutative(this *QHashCombineCommutative) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QHashCombineCommutativeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
