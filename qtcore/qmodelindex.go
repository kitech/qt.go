package qtcore

// /usr/include/qt/QtCore/qabstractitemmodel.h
// #include <qabstractitemmodel.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QModelIndex struct {
	*qtrt.CObject
}
type QModelIndex_ITF interface {
	QModelIndex_PTR() *QModelIndex
}

func (ptr *QModelIndex) QModelIndex_PTR() *QModelIndex { return ptr }

func (this *QModelIndex) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QModelIndex) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQModelIndexFromPointer(cthis unsafe.Pointer) *QModelIndex {
	return &QModelIndex{&qtrt.CObject{cthis}}
}
func (*QModelIndex) NewFromPointer(cthis unsafe.Pointer) *QModelIndex {
	return NewQModelIndexFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:58
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QModelIndex()
func NewQModelIndex() *QModelIndex {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QModelIndexC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQModelIndex)
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int row()
func (this *QModelIndex) Row() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex3rowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:61
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int column()
func (this *QModelIndex) Column() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex6columnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:62
// index:0
// Public inline Visibility=Default Availability=Available
// [8] quintptr internalId()
func (this *QModelIndex) InternalId() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex10internalIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:63
// index:0
// Public inline Visibility=Default Availability=Available
// [8] void * internalPointer()
func (this *QModelIndex) InternalPointer() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex15internalPointerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QModelIndex parent()
func (this *QModelIndex) Parent() *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex6parentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:65
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int)
func (this *QModelIndex) Sibling(row int, column int) *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex7siblingEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:67
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QModelIndex child(int, int)
func (this *QModelIndex) Child(row int, column int) *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex5childEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QVariant data(int)
func (this *QModelIndex) Data(role int) *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex4dataEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags()
func (this *QModelIndex) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:71
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QAbstractItemModel * model()
func (this *QModelIndex) Model() *QAbstractItemModel /*777 const QAbstractItemModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex5modelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:72
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QModelIndex) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QModelIndex &)
func (this *QModelIndex) Operator_equal_equal(other QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QModelIndex_PTR() != nil {
		convArg0 = other.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndexeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:75
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QModelIndex &)
func (this *QModelIndex) Operator_not_equal(other QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QModelIndex_PTR() != nil {
		convArg0 = other.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndexneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const QModelIndex &)
func (this *QModelIndex) Operator_less_than(other QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QModelIndex_PTR() != nil {
		convArg0 = other.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndexltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQModelIndex(this *QModelIndex) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QModelIndexD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
