package qtcore

// /usr/include/qt/QtCore/qabstractitemmodel.h
// #include <qabstractitemmodel.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QPersistentModelIndex struct {
	*qtrt.CObject
}

func (this *QPersistentModelIndex) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPersistentModelIndex) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQPersistentModelIndexFromPointer(cthis unsafe.Pointer) *QPersistentModelIndex {
	return &QPersistentModelIndex{&qtrt.CObject{cthis}}
}
func (*QPersistentModelIndex) NewFromPointer(cthis unsafe.Pointer) *QPersistentModelIndex {
	return NewQPersistentModelIndexFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:107
// index:0
// Public
// void QPersistentModelIndex()
func NewQPersistentModelIndex() *QPersistentModelIndex {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QPersistentModelIndexC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPersistentModelIndexFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:108
// index:1
// Public
// void QPersistentModelIndex(const class QModelIndex &)
func NewQPersistentModelIndex_1(index *QModelIndex) *QPersistentModelIndex {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QPersistentModelIndexC2ERK11QModelIndex", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPersistentModelIndexFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:110
// index:0
// Public
// void ~QPersistentModelIndex()
func DeleteQPersistentModelIndex(*QPersistentModelIndex) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QPersistentModelIndexD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:122
// index:0
// Public inline
// void swap(class QPersistentModelIndex &)
func (this *QPersistentModelIndex) Swap(other *QPersistentModelIndex) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QPersistentModelIndex4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:127
// index:0
// Public
// int row()
func (this *QPersistentModelIndex) Row() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex3rowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:128
// index:0
// Public
// int column()
func (this *QPersistentModelIndex) Column() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex6columnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:129
// index:0
// Public
// void * internalPointer()
func (this *QPersistentModelIndex) InternalPointer() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex15internalPointerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:130
// index:0
// Public
// quintptr internalId()
func (this *QPersistentModelIndex) InternalId() uint64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex10internalIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:131
// index:0
// Public
// QModelIndex parent()
func (this *QPersistentModelIndex) Parent() *QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex6parentEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:132
// index:0
// Public
// QModelIndex sibling(int, int)
func (this *QPersistentModelIndex) Sibling(row int, column int) *QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex7siblingEii", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:134
// index:0
// Public
// QModelIndex child(int, int)
func (this *QPersistentModelIndex) Child(row int, column int) *QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex5childEii", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:136
// index:0
// Public
// QVariant data(int)
func (this *QPersistentModelIndex) Data(role int) *QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex4dataEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:137
// index:0
// Public
// Qt::ItemFlags flags()
func (this *QPersistentModelIndex) Flags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:138
// index:0
// Public
// const QAbstractItemModel * model()
func (this *QPersistentModelIndex) Model() *QAbstractItemModel /*444 const QAbstractItemModel **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex5modelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:139
// index:0
// Public
// bool isValid()
func (this *QPersistentModelIndex) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QPersistentModelIndex7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
