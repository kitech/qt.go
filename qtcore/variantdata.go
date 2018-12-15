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
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type VariantData struct {
	*qtrt.CObject
}
type VariantData_ITF interface {
	VariantData_PTR() *VariantData
}

func (ptr *VariantData) VariantData_PTR() *VariantData { return ptr }

func (this *VariantData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *VariantData) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewVariantDataFromPointer(cthis unsafe.Pointer) *VariantData {
	return &VariantData{&qtrt.CObject{cthis}}
}
func (*VariantData) NewFromPointer(cthis unsafe.Pointer) *VariantData {
	return NewVariantDataFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmetatype.h:836
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void VariantData(const int, const void *, const uint)

/*

 */
func (*VariantData) NewForInherit(metaTypeId_ int, data_ unsafe.Pointer /*666*/, flags_ uint) *VariantData {
	return NewVariantData(metaTypeId_, data_, flags_)
}
func NewVariantData(metaTypeId_ int, data_ unsafe.Pointer /*666*/, flags_ uint) *VariantData {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QtMetaTypePrivate11VariantDataC2EiPKvj", qtrt.FFI_TYPE_POINTER, metaTypeId_, data_, flags_)
	qtrt.ErrPrint(err, rv)
	gothis := NewVariantDataFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteVariantData)
	return gothis
}

func DeleteVariantData(this *VariantData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11VariantDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
