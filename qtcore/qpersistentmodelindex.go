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
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QPersistentModelIndex struct {
	*qtrt.CObject
}
type QPersistentModelIndex_ITF interface {
	QPersistentModelIndex_PTR() *QPersistentModelIndex
}

func (ptr *QPersistentModelIndex) QPersistentModelIndex_PTR() *QPersistentModelIndex { return ptr }

func (this *QPersistentModelIndex) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPersistentModelIndex) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPersistentModelIndexFromPointer(cthis unsafe.Pointer) *QPersistentModelIndex {
	return &QPersistentModelIndex{&qtrt.CObject{cthis}}
}
func (*QPersistentModelIndex) NewFromPointer(cthis unsafe.Pointer) *QPersistentModelIndex {
	return NewQPersistentModelIndexFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPersistentModelIndex()
func NewQPersistentModelIndex() *QPersistentModelIndex {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QPersistentModelIndexC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPersistentModelIndexFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPersistentModelIndex)
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:108
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPersistentModelIndex(const QModelIndex &)
func NewQPersistentModelIndex_1(index QModelIndex_ITF) *QPersistentModelIndex {
	var convArg0 = index.QModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QPersistentModelIndexC2ERK11QModelIndex", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPersistentModelIndexFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPersistentModelIndex)
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QPersistentModelIndex()
func DeleteQPersistentModelIndex(this *QPersistentModelIndex) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QPersistentModelIndexD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:122
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QPersistentModelIndex &)
func (this *QPersistentModelIndex) Swap(other QPersistentModelIndex_ITF) {
	var convArg0 = other.QPersistentModelIndex_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QPersistentModelIndex4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] int row()
func (this *QPersistentModelIndex) Row() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex3rowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:128
// index:0
// Public Visibility=Default Availability=Available
// [4] int column()
func (this *QPersistentModelIndex) Column() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex6columnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] void * internalPointer()
func (this *QPersistentModelIndex) InternalPointer() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex15internalPointerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:130
// index:0
// Public Visibility=Default Availability=Available
// [8] quintptr internalId()
func (this *QPersistentModelIndex) InternalId() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex10internalIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:131
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex parent()
func (this *QPersistentModelIndex) Parent() *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex6parentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:132
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int)
func (this *QPersistentModelIndex) Sibling(row int, column int) *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex7siblingEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:134
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex child(int, int)
func (this *QPersistentModelIndex) Child(row int, column int) *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex5childEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:136
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant data(int)
func (this *QPersistentModelIndex) Data(role int) *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex4dataEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags()
func (this *QPersistentModelIndex) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] const QAbstractItemModel * model()
func (this *QPersistentModelIndex) Model() *QAbstractItemModel /*777 const QAbstractItemModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex5modelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:139
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QPersistentModelIndex) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
