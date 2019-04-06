//  header block begin

// +build !minimal

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

// /usr/include/qt/QtCore/qvariant.h:254
// index:32
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QRegularExpression &)

/*
Constructs an invalid variant.
*/
func (*QVariant) NewForInherit32(re QRegularExpression_ITF) *QVariant {
	return NewQVariant32(re)
}
func NewQVariant32(re QRegularExpression_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if re != nil && re.QRegularExpression_PTR() != nil {
		convArg0 = re.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK18QRegularExpression", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:266
// index:40
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QModelIndex &)

/*
Constructs an invalid variant.
*/
func (*QVariant) NewForInherit40(modelIndex QModelIndex_ITF) *QVariant {
	return NewQVariant40(modelIndex)
}
func NewQVariant40(modelIndex QModelIndex_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if modelIndex != nil && modelIndex.QModelIndex_PTR() != nil {
		convArg0 = modelIndex.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK11QModelIndex", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:267
// index:41
// Public Visibility=Default Availability=Available
// [-2] void QVariant(const QPersistentModelIndex &)

/*
Constructs an invalid variant.
*/
func (*QVariant) NewForInherit41(modelIndex QPersistentModelIndex_ITF) *QVariant {
	return NewQVariant41(modelIndex)
}
func NewQVariant41(modelIndex QPersistentModelIndex_ITF) *QVariant {
	var convArg0 unsafe.Pointer
	if modelIndex != nil && modelIndex.QPersistentModelIndex_PTR() != nil {
		convArg0 = modelIndex.QPersistentModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QVariantC2ERK21QPersistentModelIndex", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVariantFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

// /usr/include/qt/QtCore/qvariant.h:330
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpression toRegularExpression() const

/*
Returns the variant as a QRegularExpression if the variant has userType() QRegularExpression; otherwise returns an empty QRegularExpression.

This function was introduced in  Qt 5.0.

See also canConvert(int targetTypeId) and convert().
*/
func (this *QVariant) ToRegularExpression() *QRegularExpression /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant19toRegularExpressionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpression)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:342
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex toModelIndex() const

/*
Returns the variant as a QModelIndex if the variant has userType() QModelIndex; otherwise returns a default constructed QModelIndex.

This function was introduced in  Qt 5.0.

See also canConvert(int targetTypeId), convert(), and toPersistentModelIndex().
*/
func (this *QVariant) ToModelIndex() *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant12toModelIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qvariant.h:343
// index:0
// Public Visibility=Default Availability=Available
// [8] QPersistentModelIndex toPersistentModelIndex() const

/*
Returns the variant as a QPersistentModelIndex if the variant has userType() QPersistentModelIndex; otherwise returns a default constructed QPersistentModelIndex.

This function was introduced in  Qt 5.5.

See also canConvert(int targetTypeId), convert(), and toModelIndex().
*/
func (this *QVariant) ToPersistentModelIndex() *QPersistentModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QVariant22toPersistentModelIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPersistentModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPersistentModelIndex)
	return rv2
}

//  body block end

//  keep block begin

func init_unused_10276() {
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
