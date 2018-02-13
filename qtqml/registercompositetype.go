package qtqml

// /usr/include/qt/QtQml/qqmlprivate.h
// #include <qqmlprivate.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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

type RegisterCompositeType struct {
	*qtrt.CObject
}
type RegisterCompositeType_ITF interface {
	RegisterCompositeType_PTR() *RegisterCompositeType
}

func (ptr *RegisterCompositeType) RegisterCompositeType_PTR() *RegisterCompositeType { return ptr }

func (this *RegisterCompositeType) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *RegisterCompositeType) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewRegisterCompositeTypeFromPointer(cthis unsafe.Pointer) *RegisterCompositeType {
	return &RegisterCompositeType{&qtrt.CObject{cthis}}
}
func (*RegisterCompositeType) NewFromPointer(cthis unsafe.Pointer) *RegisterCompositeType {
	return NewRegisterCompositeTypeFromPointer(cthis)
}

func DeleteRegisterCompositeType(this *RegisterCompositeType) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21RegisterCompositeTypeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
