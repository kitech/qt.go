package qtcore

// /usr/include/qt/QtCore/qsharedpointer_impl.h
// #include <qsharedpointer_impl.h>
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
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type NormalDeleter struct {
	*qtrt.CObject
}

func (this *NormalDeleter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *NormalDeleter) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewNormalDeleterFromPointer(cthis unsafe.Pointer) *NormalDeleter {
	return &NormalDeleter{&qtrt.CObject{cthis}}
}
func (*NormalDeleter) NewFromPointer(cthis unsafe.Pointer) *NormalDeleter {
	return NewNormalDeleterFromPointer(cthis)
}

func DeleteNormalDeleter(this *NormalDeleter) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13NormalDeleterD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
