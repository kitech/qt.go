package qtcore

// /usr/include/qt/QtCore/qvariant.h
// #include <qvariant.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 109
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
type QVariantComparisonHelper struct {
	*qtrt.CObject
}
type QVariantComparisonHelper_ITF interface {
	QVariantComparisonHelper_PTR() *QVariantComparisonHelper
}

func (ptr *QVariantComparisonHelper) QVariantComparisonHelper_PTR() *QVariantComparisonHelper {
	return ptr
}

func (this *QVariantComparisonHelper) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QVariantComparisonHelper) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQVariantComparisonHelperFromPointer(cthis unsafe.Pointer) *QVariantComparisonHelper {
	return &QVariantComparisonHelper{&qtrt.CObject{cthis}}
}
func (*QVariantComparisonHelper) NewFromPointer(cthis unsafe.Pointer) *QVariantComparisonHelper {
	return NewQVariantComparisonHelperFromPointer(cthis)
}

// /usr/include/qt/QtCore/qvariant.h:560
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QVariantComparisonHelper(const QVariant &)

/*

 */
func (*QVariantComparisonHelper) NewForInherit(var_ QVariant_ITF) *QVariantComparisonHelper {
	return NewQVariantComparisonHelper(var_)
}
func NewQVariantComparisonHelper(var_ QVariant_ITF) *QVariantComparisonHelper {
	var convArg0 unsafe.Pointer
	if var_ != nil && var_.QVariant_PTR() != nil {
		convArg0 = var_.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QVariantComparisonHelperC2ERK8QVariant", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantComparisonHelperFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariantComparisonHelper)
	return gothis
}

func DeleteQVariantComparisonHelper(this *QVariantComparisonHelper) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QVariantComparisonHelperD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
