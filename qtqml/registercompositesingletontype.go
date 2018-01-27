package qtqml

// /usr/include/qt/QtQml/qqmlprivate.h
// #include <qqmlprivate.h>
// #include <QtQml>

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
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  ext block end

//  body block begin
type RegisterCompositeSingletonType struct {
	*qtrt.CObject
}

func (this *RegisterCompositeSingletonType) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *RegisterCompositeSingletonType) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewRegisterCompositeSingletonTypeFromPointer(cthis unsafe.Pointer) *RegisterCompositeSingletonType {
	return &RegisterCompositeSingletonType{&qtrt.CObject{cthis}}
}
func (*RegisterCompositeSingletonType) NewFromPointer(cthis unsafe.Pointer) *RegisterCompositeSingletonType {
	return NewRegisterCompositeSingletonTypeFromPointer(cthis)
}

//  body block end
