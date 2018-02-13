package qtcore

// /usr/include/qt/QtCore/qmetatype.h
// #include <qmetatype.h>
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

type VectorBoolElements struct {
	*qtrt.CObject
}
type VectorBoolElements_ITF interface {
	VectorBoolElements_PTR() *VectorBoolElements
}

func (ptr *VectorBoolElements) VectorBoolElements_PTR() *VectorBoolElements { return ptr }

func (this *VectorBoolElements) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *VectorBoolElements) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewVectorBoolElementsFromPointer(cthis unsafe.Pointer) *VectorBoolElements {
	return &VectorBoolElements{&qtrt.CObject{cthis}}
}
func (*VectorBoolElements) NewFromPointer(cthis unsafe.Pointer) *VectorBoolElements {
	return NewVectorBoolElementsFromPointer(cthis)
}

func DeleteVectorBoolElements(this *VectorBoolElements) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18VectorBoolElementsD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
